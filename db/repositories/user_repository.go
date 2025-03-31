package repositories

import (
	m "dnd-api/db/models"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (r *UserRepository) GetByUid(uid string) *m.User {
	var user m.User
	r.Db.
		Preload("Roles.Permissions").
		Where("uid = ?", uid).
		Find(&user)

	return &user
}

func (r *UserRepository) GetByEmail(email string) *m.User {
	var user m.User
	r.Db.
		Preload("Roles.Permissions").
		Where("email = ?", email).
		Find(&user)

	return &user
}
