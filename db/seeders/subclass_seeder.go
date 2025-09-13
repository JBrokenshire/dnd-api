package seeders

import (
	m "dnd-api/db/models"
	"log"
)

func (s *Seeder) SetSubclasses() {
	subclasses := []*m.Subclass{
		{
			ID:               1,
			ClassId:          1,
			Name:             "College of Glamour",
			ShortDescription: "The College of Glamour is the home of bards who mastered their craft in the vibrant realm of the Feywild or under the tutelage of someone who dwelled there. Tutored by satyrs, eladrin, and other fey, these bards learn to use their magic to delight and captivate others.",
		},
		{
			ID:               2,
			ClassId:          2,
			Name:             "Echo Knight",
			ShortDescription: "A mysterious and feared frontline warrior of the Kryn Dynasty, the Echo Knight has mastered the art of using dunamis to summon the fading shades of unrealized timelines to aid them in battle. Surrounded by echoes of their own might, they charge into the fray as a cycling swarm of shadows and strikes.",
		},
		{
			ID:               3,
			ClassId:          3,
			Name:             "Order of the Mutant",
			ShortDescription: "Over generations of experimentation, a splinter order of blood hunters honed the way in which hemocraft alters the body, using corrupted alchemy and toxic elixirs to alter their blood even further. Over time, they have modified their capabilities in battle, becoming something beyond what they once were. Calling themselves the Order of the Mutant, these blood hunters now specialize in assessing the strengths and weaknesses of their foes, altering their biology to be best prepared for any conflict.",
		},
	}

	for _, subclass := range subclasses {
		err := s.DB.Where("id = ?", subclass.ID).FirstOrCreate(&subclass).Error
		if err != nil {
			log.Printf("Error creating subclass with id %v in seeder: %v", subclass.ID, err.Error())
		}
	}
}
