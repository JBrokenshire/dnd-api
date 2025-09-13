package seeders

import (
	m "dnd-api/db/models"
	"log"
)

func (s *Seeder) SetRaces() {
	races := []m.Race{
		{
			ID:               1,
			Name:             "Earth Genasi",
			ShortDescription: "Genasi carry the power of the elemental planes of air, earth, fire, and water in their blood.\n\nAs an earth genasi, you are descended from the cruel and greedy dao, though you aren’t necessarily evil. You have inherited some measure of control over earth, reveling in superior strength and solid power. You tend to avoid rash decisions, pausing long enough to consider your options before taking action.\n\nElemental earth manifests differently from one individual to the next. Some earth genasi always have bits of dust falling from their bodies and mud clinging to their clothes, never getting clean no matter how often they bathe. Others are as shiny and polished as gemstones, with skin tones of deep brown or black, eyes sparkling like agates. Earth genasi can also have smooth metallic flesh, dull iron skin spotted with rust, a pebbled and rough hide, or even a coating of tiny embedded crystals. The most arresting have fissures in their flesh, from which faint light shines.",
			CreatureType:     "Humanoid",
			Size:             "Medium (about 5-6 feet tall)",
			BaseSpeed:        30,
		},
		{
			ID:               2,
			Name:             "Dragonborn",
			ShortDescription: "The ancestors of dragonborn hatched from the eggs of chromatic and metallic dragons.",
			CreatureType:     "Humanoid",
			Size:             "Medium (about 5-6 feet tall)",
			BaseSpeed:        30,
		},
		{
			ID:               3,
			Name:             "Stout Halfling",
			ShortDescription: "The diminutive halflings survive in a world full of larger creatures by avoiding notice or, barring that, avoiding offense.\n\nAs a stout halfling, you’re hardier than average and have some resistance to poison. Some say that stouts have dwarven blood. In the Forgotten Realms, these halflings are called stronghearts, and they’re most common in the south.",
			CreatureType:     "Humanoid",
			Size:             "Small (about 3 feet tall)",
			BaseSpeed:        25,
		},
	}

	for _, race := range races {
		err := s.DB.Where("id = ?", race.ID).FirstOrCreate(&race).Error
		if err != nil {
			log.Printf("Error creating race with id %v in seeder: %v", race.ID, err.Error())
		}
	}
}
