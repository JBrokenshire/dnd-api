package repositories

import (
	"github.com/jinzhu/gorm"
)

type KeyRepository struct {
	*Repository
}

func NewKeyRepository(db *gorm.DB) *KeyRepository {
	return &KeyRepository{&Repository{Db: db}}
}

func (r *KeyRepository) GetKeys(enterpriseUid string) []m.ApiKey {
	var keys []m.ApiKey
	r.Db.Preload("Policy").Where("enterprise_uid = ?", enterpriseUid).Find(&keys)
	return keys
}

func (r *KeyRepository) GetKey(enterpriseUid string, uid string) *m.ApiKey {
	key := &m.ApiKey{}
	r.Db.Where("enterprise_uid = ?", enterpriseUid).Where("uid = ?", uid).Find(key)
	return key
}

func (r *KeyRepository) FindKey(apiKey string) *m.ApiKey {
	key := &m.ApiKey{}
	r.Db.Where("enabled = 1 and uid = ?", apiKey).Take(key)
	return key
}
