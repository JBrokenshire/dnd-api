package seeders

import (
	m "dnd-api/db/models"
	"log"
)

func (s *Seeder) SetRaces() {
	races := []*m.Race{
		{
			ID:               1,
			Name:             "Dragonborn",
			ShortDescription: "The ancestors of dragonborn hatched from the eggs of chromatic and metallic dragons.",
			CreatureType:     "Humanoid",
			Size:             "Medium (about 5-7 feet tall)",
			BaseSpeed:        30,
		},
		{
			ID:               2,
			Name:             "Dwarf",
			ShortDescription: "Dwarves were raised from the earth in the elder days by a deity of the forge.",
			CreatureType:     "Humanoid",
			Size:             "Medium (about 4-5 feet tall)",
			BaseSpeed:        30,
		},
		{
			ID:               3,
			Name:             "Elf",
			ShortDescription: "The elves’ curiosity led many of them to explore other planes of existence.",
			CreatureType:     "Humanoid",
			Size:             "Medium (about 5-6 feet tall)",
			BaseSpeed:        30,
		},
		{
			ID:               4,
			Name:             "Gnome",
			ShortDescription: "Gnomes are magical folk created by gods of invention, illusions, and life underground.",
			CreatureType:     "Humanoid",
			Size:             "Small (about 3-4 feet tall)",
			BaseSpeed:        30,
		},
		{
			ID:               5,
			Name:             "Goliath",
			ShortDescription: "Goliaths are distant descendants of giants and seek heights above those reached by their ancestors.",
			CreatureType:     "Humanoid",
			Size:             "Medium (about 6-8 feet tall)",
			BaseSpeed:        35,
		},
		{
			ID:               6,
			Name:             "Halfling",
			ShortDescription: "Halflings possess a brave and adventurous spirit that leads them on journeys of discovery.",
			CreatureType:     "Humanoid",
			Size:             "Small (about 2-3 feet tall)",
			BaseSpeed:        30,
		},
		{
			ID:               7,
			Name:             "Human",
			ShortDescription: "Found throughout the multiverse, humans are as varied as they are numerous.",
			CreatureType:     "Humanoid",
			Size:             "Medium (about 4-7 feet tall) or Small (about 2-4 feet tall)",
			BaseSpeed:        30,
		},
		{
			ID:               8,
			Name:             "Orc",
			ShortDescription: "Orcs are equipped with gifts to help them wander great plains, vast caverns, and churning seas.",
			CreatureType:     "Humanoid",
			Size:             "Medium (about 6-7 feet tall)",
			BaseSpeed:        30,
		},
		{
			ID:               9,
			Name:             "Tiefling",
			ShortDescription: "Tieflings are either born in the Lower Planes or have fiendish ancestors who originated there.",
			CreatureType:     "Humanoid",
			Size:             "Medium (about 4-7 feet tall) or Small (about 3-4 feet tall)",
			BaseSpeed:        30,
		},
		{
			ID:               10,
			Name:             "Aarakocra",
			ShortDescription: "Sequestered in high mountains atop tall trees, the aarakocra, sometimes called birdfolk, evoke fear and wonder.",
			CreatureType:     "Humanoid",
			Size:             "Medium (about 5 feet tall)",
			BaseSpeed:        25,
		},
		{
			ID:               11,
			Name:             "Deep Gnome",
			ShortDescription: "A gnome’s energy and enthusiasm for living shines through every inch of his or her tiny body.",
			CreatureType:     "Humanoid",
			Size:             "Small (about 3-4 feet tall)",
			BaseSpeed:        30,
		},
		{
			ID:               12,
			Name:             "Earth Genasi",
			ShortDescription: "Genasi carry the power of the elemental planes of air, earth, fire, and water in their blood.\n\nAs an earth genasi, you are descended from the cruel and greedy dao, though you aren’t necessarily evil. You have inherited some measure of control over earth, reveling in superior strength and solid power. You tend to avoid rash decisions, pausing long enough to consider your options before taking action.\n\nElemental earth manifests differently from one individual to the next. Some earth genasi always have bits of dust falling from their bodies and mud clinging to their clothes, never getting clean no matter how often they bathe. Others are as shiny and polished as gemstones, with skin tones of deep brown or black, eyes sparkling like agates. Earth genasi can also have smooth metallic flesh, dull iron skin spotted with rust, a pebbled and rough hide, or even a coating of tiny embedded crystals. The most arresting have fissures in their flesh, from which faint light shines.",
			CreatureType:     "Humanoid",
			Size:             "Medium (about 4-7 feet tall)",
			BaseSpeed:        30,
		},
		{
			ID:               13,
			Name:             "Fire Genasi",
			ShortDescription: "Genasi carry the power of the elemental planes of air, earth, fire, and water in their blood.\n\nAs a fire genasi, you have inherited the volatile mood and keen mind of the efreet. You tend toward impatience and making snap judgments. Rather than hide your distinctive appearance, you exult in it.\n\nNearly all fire genasi are feverishly hot as if burning inside, an impression reinforced by flaming red, coal- black, or ash-gray skin tones. The more human-looking have fiery red hair that writhes under extreme emotion, while more exotic specimens sport actual flames dancing on their heads. Fire genasi voices might sound like crackling flames, and their eyes flare when angered. Some are accompanied by the faint scent of brimstone.",
			CreatureType:     "Humanoid",
			Size:             "Medium (about 4-7 feet tall)",
			BaseSpeed:        30,
		},
		{
			ID:               14,
			Name:             "Water Genasi",
			ShortDescription: "Genasi carry the power of the elemental planes of air, earth, fire, and water in their blood.\n\nThe lapping of waves, the spray of sea foam on the wind, the ocean depths—all of these things call to your heart. You wander freely and take pride in your independence, though others might consider you selfish.\n\nMost water genasi look as if they just finished bathing, with beads of moisture collecting on their skin and hair. They smell of fresh rain and clean water. Blue or green skin is common, and most have somewhat overlarge eyes, blue-black in color. A water genasi’s hair might float freely, swaying and waving as if underwater. Some have voices with undertones reminiscent of whale song or trickling streams.",
			CreatureType:     "Humanoid",
			Size:             "Medium (about 4-7 feet tall)",
			BaseSpeed:        30,
		},
		{
			ID:               15,
			Name:             "Air Genasi",
			ShortDescription: "Genasi carry the power of the elemental planes of air, earth, fire, and water in their blood.\n\nAs an air genasi, you are descended from the djinn. As changeable as the weather, your moods shift from calm to wild and violent with little warning, but these storms rarely last long.\n\nAir genasi typically have light blue skin, hair, and eyes. A faint but constant breeze accompanies them, tousling the hair and stirring the clothing. Some air genasi speak with breathy voices, marked by a faint echo. A few display odd patterns in their flesh or grow crystals from their scalps.",
			CreatureType:     "Humanoid",
			Size:             "Medium (about 4-7 feet tall)",
			BaseSpeed:        30,
		},
	}

	for _, race := range races {
		err := s.DB.Where("id = ?", race.ID).FirstOrCreate(&race).Error
		if err != nil {
			log.Printf("Error creating race with id %v in seeder: %v", race.ID, err.Error())
		}
	}
}
