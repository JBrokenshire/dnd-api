package repositories

import (
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type FailedKeyValidationRepository struct {
	Db *gorm.DB
}

func NewFailedKeyValidateRepository(db *gorm.DB) *FailedKeyValidationRepository {
	return &FailedKeyValidationRepository{Db: db}
}

func (r *FailedKeyValidationRepository) CreateFailedValidation(ipAddress string) {
	failedValidation := &models.FailedKeyValidation{
		IpAddress: ipAddress,
	}
	err := r.Db.Create(failedValidation).Error
	if err != nil {
		log.Println(err)
	}
}

func (r *FailedKeyValidationRepository) BruteForceCount(ip string, since time.Time) int64 {
	var count int64
	err := r.Db.Model(&models.FailedKeyValidation{}).Where("ip_address = ?", ip).Where("created_at > ?", since).Count(&count).Error
	if err != nil {
		log.Println(err)
	}
	return count
}
