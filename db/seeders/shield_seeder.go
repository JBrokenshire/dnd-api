package seeders

import (
	"dnd-api/db/models"
	"log"
)

func (s *Seeder) SetShields() {
	shields := []models.Shield{
		{
			ItemID:  37,
			BonusAC: 2,
		},
	}

	for _, shield := range shields {
		err := s.DB.Where("item_id = ?", shield.ItemID).FirstOrCreate(&shield).Error
		if err != nil {
			log.Printf("error creating shield with item_id '%v': %v", shield.ItemID, err)
		}
	}
}
