package seeders

import (
	m "dnd-api/db/models"
	"log"
)

func (s *Seeder) SetSubclasses() {
	subclasses := []*m.Subclass{
		{
			ID:               1,
			ClassId:          2,
			Name:             "College of Glamour",
			ShortDescription: "The College of Glamour is the home of bards who mastered their craft in the vibrant realm of the Feywild or under the tutelage of someone who dwelled there. Tutored by satyrs, eladrin, and other fey, these bards learn to use their magic to delight and captivate others.",
		},
	}

	for _, subclass := range subclasses {
		err := s.DB.Where("id = ?", subclass.ID).FirstOrCreate(&subclass).Error
		if err != nil {
			log.Printf("Error creating subclass with id %v in seeder: %v", subclass.ID, err.Error())
		}
	}
}
