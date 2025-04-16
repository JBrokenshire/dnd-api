package repositories

import (
	m "dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"time"
)

type UserRepository struct {
	*Repository
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		Repository: &Repository{
			Db: db,
		},
	}
}

func (r *UserRepository) GetById(id interface{}) *m.User {
	user := &m.User{}
	r.Db.Where("id = ?", id).Find(user)
	return user
}

func (r *UserRepository) GetByUsername(username string) *m.User {
	user := m.User{}
	r.Db.Where("username = ?", username).Find(&user)
	return &user
}

// BruteForceCount Count failed login attempts for user
func (r *UserRepository) BruteForceCount(userId string, since *time.Time) int64 {
	var count int64
	r.Db.Model(&m.FailedLogin{}).Where("user_id = ?", userId).Where("created_at > ?", since).Count(&count)
	return count
}
