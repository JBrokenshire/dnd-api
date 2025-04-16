package factories

import (
	m "dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
	"log"
)

func NewUser(db *gorm.DB, user *m.User) {
	fillUserDefaults(user)
	err := db.Create(user).Error
	if err != nil {
		log.Println("Error creating user in factory: ", err.Error())
	}
}

func fillUserDefaults(user *m.User) {
	if user.Username == "" {
		user.Username = random.String(16)
	}
	if user.Password == "" {
		user.Password = random.String(16)
	}
}
