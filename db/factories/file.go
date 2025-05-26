package factories

import (
	"dnd-api/db/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
	"log"
)

func NewFile(db *gorm.DB, file *models.File) {
	fillFileDefaults(file)
	if err := db.Create(&file).Error; err != nil {
		log.Println("Error creating file in factory: ", err.Error())
	}
}

func fillFileDefaults(file *models.File) {
	if file.Model == "" {
		file.Model = models.FileModelClassLogo
	}
	if file.ModelId == 0 {
		file.ModelId = 1
	}
	if file.Filename == "" {
		file.Filename = random.String(16) + ".jpg"
	}
	if file.FileLocation == "" {
		file.FileLocation = fmt.Sprintf("/%v", random.String(16))
	}
}
