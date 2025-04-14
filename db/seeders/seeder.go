package seeders

import (
	"github.com/jinzhu/gorm"
)
	
type Seeder struct {
	DB *gorm.DB
}

func NewSeeder(db *gorm.DB) *Seeder {
	return &Seeder{DB: db}
}
