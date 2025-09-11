package seeders

import (
	m "dnd-api/db/models"
	"log"
)

func (s *Seeder) SetCharacters() {
	characters := []m.Character{
		{
			ID:                1,
			UserId:            1,
			Name:              "Axel Claystride",
			ClassId:           1,
			RaceId:            1,
			Level:             6,
			Pronouns:          "He/Him",
			BackgroundId:      1,
			Alignment:         "Chaotic Good",
			Gender:            "Male",
			Eyes:              "Dark Brown",
			Size:              "Medium",
			Height:            "6' 0\"",
			Hair:              "Black, Short and Spiky",
			Skin:              "Rocky, Gray and Cracked",
			Age:               "23",
			Weight:            "200lb.",
			Strength:          9,
			Dexterity:         16,
			Constitution:      14,
			Intelligence:      11,
			Wisdom:            12,
			Charisma:          16,
			CurrentHitPoints:  38,
			MaxHitPoints:      38,
			AdvancementType:   m.AdvancementTypeMilestone,
			HitPointType:      m.HitPointTypeManual,
			Proficiencies:     `{"armour":["Light Armour"], "weapons":["Simple Weapons"], "tools":["Lute","Zulkoon","Painters' Supplies"], "languages":["Common","Primordial","Celestial"]}`,
			PersonalityTraits: `["I feel restless if I stay in one place too long."]`,
			Ideals:            `["No stone can hold me, no tradition can bind me."]`,
			Bonds:             `["I carry a keepsake from my family, a reminder that I always have a home somewhere."]`,
			Flaws:             `["I hide my doubts behind bravado and humour, even when it hurts me."]`,
			Organisations:     "[]",
			Allies:            "[]",
			Enemies:           "[]",
			Backstory: `Axel Claystride was born to the wandering Claystride tribe, earth genasi who roamed the land as nomads, forever in step with the shifting world beneath their feet. His childhood was filled with stone and dust, laughter and hardship. He was close to his parents, who cherished him, and surrounded by cousins who - while often teasing him for fumbling through traditional duties like forging or quarry work - were secretly impressed by the strange gift he carried. Axel wasn’t made for hammering steel or shaping stone. His hands found rhythm, not craft.

It was during a midsummer celebration that Axel’s gift first revealed itself. The tribe drank, danced, and told stories around great fires while Axel sat quietly to the side, humming to himself. Without realizing it, his humming grew into song, and his song drew ears. First a few heads turned. Then a few voices joined in. Before long, the whole tribe was listening, clapping, and laughing as Axel spun his melody into a rowdy chorus. When the song ended, the eruption of applause and cheers was like an avalanche, both exhilarating and terrifying. Axel had never known such joy or such belonging as in that moment.

But when the years rolled on, the weight of tradition pressed down. His cousins grew strong at the forge, his kin proved themselves in quarry and craft, while Axel still carried only music in his chest. It wasn’t enough to stay. With pride for his tribe but no place among their path, Axel left trading security for song, roots for rhythm. He promised to keep in touch, but both he and the tribe wander constantly, and those bonds now survive only in memories and keepsakes he carries in his pack.

On the road, Axel learned to make an honest living in taverns and keeps. He’d trade songs for food, tales for ale, and laughter for a bed. When times grew lean, he relied on charm: winning over travelers and kind strangers with an easy grin and the promise of a tune. He never stooped to crime, but he walked close to the line, always leaning on the belief that music should be freely given, and that joy was worth any gamble.

On stage, Axel is not just a bard, he is a force. When the rhythm takes him, the cracks of his stone-hardened skin glow faint gold, light pulsing with the beat of his performance. Patrons swear they can feel the music more than hear it, their spirits pulled upward by a strange, radiant energy. He has no need of stage names or false masks... Axel Claystride is enough. He is proud of his name, his tribe, and his earth-born heritage.

Yet, in the quiet after the crowds fade, Axel’s doubts creep in. He wonders if he was a fool to abandon the certainty of the forge, to trade stone walls for shifting roads. There are nights when loneliness weighs heavier than applause, when he longs for the steady warmth of his family. In those moments, he whispers regrets to himself, imagining the alternate life he could have lived. But when the music flows again, when laughter erupts from weary faces, when cheers rise from strangers who forget their troubles even for a song Axel remembers his purpose.

He may not know happiness for himself yet. But if his music can carve it out of stone for others, then he’ll keep walking, keep singing, until the rhythm of the world brings that joy home to him as well.`,
		},
	}

	for _, character := range characters {
		err := s.DB.Where("id = ?", character.ID).FirstOrCreate(&character).Error
		if err != nil {
			log.Printf("Error creating character with id %v in seeder: %v", character.ID, err.Error())
		}
	}
}
