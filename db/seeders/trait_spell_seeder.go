package seeders

import (
	m "dnd-api/db/models"
	"dnd-api/pkg/utils"
	"log"
)

func (s *Seeder) SetTraitSpells() {
	traitSpells := []*m.TraitSpell{
		{
			ID:      1,
			TraitId: "earth-genasi-merge-with-stone",
			SpellId: 11,
			Uses:    utils.IntPointer(1),
			Reset:   utils.StrPointer(m.RestTypeLong),
		},
	}

	for _, traitSpell := range traitSpells {
		err := s.DB.Where("id = ?", traitSpell.ID).FirstOrCreate(&traitSpell).Error
		if err != nil {
			log.Printf("Error creating trait spell with id %v in seeder: %v", traitSpell.ID, err.Error())
		}
	}
}
