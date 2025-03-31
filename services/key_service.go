package services

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type KeyService struct {
	Db         *gorm.DB
	FailedRepo *repositories.FailedKeyValidationRepository
	KeyRepo    *repositories.KeyRepository
}

func NewKeyService(Db *gorm.DB) *KeyService {
	return &KeyService{
		Db:         Db,
		FailedRepo: repositories.NewFailedKeyValidateRepository(Db),
		KeyRepo:    repositories.NewKeyRepository(Db),
	}
}

// CheckBlocked returns true if the user is blocked because of a bruteforce attempt
func (f *KeyService) CheckBlocked(userIp string) error {
	conf := config.Get()
	since := time.Now().Add(-1 * conf.KeyBruteForceDuration)
	failedAttempts := f.FailedRepo.BruteForceCount(userIp, since)
	if failedAttempts > conf.KeyBruteForceLimit {
		return errors.New("too many attempts, you have been blocked temporarily")
	}
	return nil
}

// ExternalKeyOk check if the external key is valid
func (f *KeyService) ExternalKeyOk(key, userIp string) error {

	if err := f.CheckBlocked(userIp); err != nil {
		return err
	}

	conf := config.Get()
	if key != conf.ExternalApiKey {
		// Create failed key validation
		f.FailedRepo.CreateFailedValidation(userIp)
		return errors.New("missing or invalid API Key")
	}
	return nil
}

// EnterpriseKeyValidate checks isf the api key presented is in teh DB and matches the IP range
func (f *KeyService) EnterpriseKeyValidate(key, userIp string) (*models.ApiKey, error) {

	if err := f.CheckBlocked(userIp); err != nil {
		return nil, err
	}

	apiKey := f.KeyRepo.FindKey(key)
	if apiKey.ID == 0 {
		f.FailedRepo.CreateFailedValidation(userIp)
		return nil, errors.New("invalid API key")
	}

	// Check it's in the allowed ranges of the Key
	if inRange := ip.InRanges(userIp, apiKey.IpRanges); inRange == false {
		return nil, errors.New("invalid IP address for KEY used")
	}

	return apiKey, nil
}
