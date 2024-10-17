package factories

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
)

func NewArmour(db *gorm.DB, weapon *models.Armour) {
	fillArmourDetails(weapon)
	db.Table("armour").Create(weapon)
}

func fillArmourDetails(a *models.Armour) {
	if a.Type == "" {
		a.Type = "Light Armour"
	}
}
