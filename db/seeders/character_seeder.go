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
		{
			ID:                2,
			UserId:            1,
			Name:              `Kaelrith "Kael" Drakeshield`,
			ClassId:           2,
			RaceId:            2,
			Level:             6,
			Pronouns:          "He/Him",
			BackgroundId:      2,
			Alignment:         "Lawful Good",
			Gender:            "Male",
			Eyes:              "Amber",
			Size:              "Medium",
			Height:            `6'6"`,
			Hair:              "A crown of curved horns",
			Skin:              "Brass/Gold-Brown",
			Age:               "26",
			Weight:            "270lb.",
			Strength:          20,
			Dexterity:         15,
			Constitution:      20,
			Intelligence:      6,
			Wisdom:            14,
			Charisma:          11,
			CurrentHitPoints:  60,
			MaxHitPoints:      72,
			AdvancementType:   m.AdvancementTypeMilestone,
			HitPointType:      m.HitPointTypeManual,
			Senses:            `["Blindsight 10ft."]`,
			Proficiencies:     `{"armour":["Heavy Armour","Light Armour","Medium Armour","Shields"],"weapons":["Martial Weapons","Simple Weapons"],"tools":["Lute","Thieves' Tools"],"languages":["Common","Draconic"]}`,
			PersonalityTraits: `["The best way to get me to do something is to tell me I can't do it","I am always calm, no matter what the situation. I never raise my voice or let my emotions control me."]`,
			Ideals:            `["Freedom. Chains are meant to be broken, as are those who would forge them (Chaotic)"]`,
			Bonds:             `["I'm trying to pay off an old debt I owe to a generous benefactor"]`,
			Flaws:             `["If there's a plan I'll forget it. If I don't forget it, I'll ignore it."]`,
			Organisations:     `["Inter-Isle Belonging of Fantastic Investigators - Marblethorn Unit"]`,
			Allies:            `["Eldrin","Idris","Vashta","Niko","Karma","Lycia","Theo","Edith","Eggy"]`,
			Enemies:           `[]`,
			Backstory: `Kaelrith Drakeshield was born in the harsh, arid lands of the Draken Wastes, a barren region where only the strongest survived. From a young age, he was taught to fend for himself, honing his skills as a warrior under the tutelage of his clan's elders. The Draken Wastes were a lawless land, and Kael quickly learned the ways of combat, developing into a skilled bounty hunter by his late teens. His physical prowess and brass dragonborn resilience made him a feared and respected figure among the bounty hunting community.

However, life as a bounty hunter was not without its complications. Kael's straightforward nature often got him into trouble, and he made more enemies than friends along the way. One such enemy was a notorious crime lord, Garvin the Ruthless, whom Kael had crossed during a high-stakes job. The confrontation left Kael with a considerable debt, as Garvin demanded retribution for the loss of his men and resources. Desperate and with few options, Kael found an unexpected benefactor in an eccentric wizard named Eldrin Valtoris.

Eldrin saw potential in the young dragonborn and offered to pay off his debt in exchange for Kael's service and loyalty. Kael accepted, grateful for the wizard's generosity. Under Eldrin's guidance, Kael was introduced to the arcane arts, specifically the Echo Knight discipline. This allowed Kael to summon spectral echoes of himself in battle, an ability that greatly enhanced his combat effectiveness.

After several years of service, Eldrin released Kael from his obligation, though Kael felt indebted to the wizard and vowed to repay him fully one day. Seeking a new purpose, Kael joined the Marblethorn Unit of the Inter-Isle Belonging of Fantastic Investigators (IIBFI) two years ago. The unit, known for tackling supernatural threats and complex investigations across the Isles, provided Kael with a new avenue to use his skills for a greater good.

Though his intellect often lags behind his peers, Kael's unwavering determination and combat expertise have earned him a place of respect within the Marblethorn Unit. He continues to work tirelessly, hoping to one day repay Eldrin and honor the legacy of his clan. Kael's journey is one of redemption and loyalty, as he seeks to balance the scales of his past while forging a future as a protector and defender.`,
		},
		{
			ID:                3,
			UserId:            1,
			Name:              "Faelan Haversham",
			ClassId:           3,
			RaceId:            3,
			Level:             3,
			Pronouns:          "He/Him",
			BackgroundId:      3,
			Alignment:         "Lawful Good",
			Gender:            "Male",
			Eyes:              "Green",
			Size:              "Small",
			Height:            `3' 2"'`,
			Hair:              "Tousled Brown",
			Skin:              "White",
			Age:               "23",
			Weight:            "45lb.",
			Strength:          12,
			Dexterity:         16,
			Constitution:      14,
			Intelligence:      10,
			Wisdom:            15,
			Charisma:          8,
			CurrentHitPoints:  28,
			MaxHitPoints:      28,
			AdvancementType:   m.AdvancementTypeMilestone,
			HitPointType:      m.HitPointTypeManual,
			Senses:            "[]",
			Proficiencies:     `{"armour":["Light Armour","Medium Armour","Shields"], "weapons":["Martial Weapons","Simple Weapons"], "tools":["Alchemist's Supplies","Viol"], "languages":["Common","Gnomish","Halfling"]}`,
			PersonalityTraits: `["I feel far more comfortable around animals than people.","I have a lesson for every situation, drawn from observing nature."]`,
			Ideals:            `["Nature. The natural world is more important than all the constructs of civilisation. (Neutral)"]`,
			Bonds:             `["An injury to the unspoiled wilderness of my home is an injury to me."]`,
			Flaws:             `["I am too enamored of ale, wine, and other intoxicants."]`,
			Backstory: `Faelan Haversham, a Stout Halfling of Acosis, was born into a family deeply rooted in the lumber and food trade that sustained their isolated community. The Havershams were known for their resilience and resourcefulness in the face of the harsh conditions surrounding Acosis, relying on a careful balance of agriculture and forestry to meet the needs of their people. Faelan's father, Thorian Haversham, was a skilled woodsman who oversaw the logging operations, while his mother, Elowen Haversham, managed the agricultural aspects, ensuring the community had enough sustenance.

Growing up in the shadow of the constant threat of Arreksis raids, Faelan developed a strong sense of duty towards the safety and prosperity of Acosis. He learned the art of survival in the depths of the forests where trade access was impossible. Faelan's childhood was marked by stories of the Mortonhelm disaster, a cautionary tale shared by elders to emphasise the dangers of unchecked magical forces.

One day, as Faelan explored the dense forests of Acosis, he stumbled upon a hidden glade where an ancient and mystical gem lay nestled among the roots of an ancient tree. The gem radiated a peculiar energy that captivated Faelan's senses. Intrigued and fueled by a desire to understand the gem's origins and purpose, he embarked on a journey beyond the borders of Acosis. 

Making his way to Myrinport, the bustling trade hub of Fortenua, Faelan sought out scholars and mages who could shed light on the mysterious gem. He was instructed to head to the famed Crenchai Isle, known for its strong magical presence and abundance of temples. He was assured that the priests and scholars of the Isle could answer his questions. 

He was set to head out from Myrinport dock and travel by boat to Crenchai Isle that evening. However, fate had other plans for him. As he navigated the crowded streets of Myrinport, the city's alarm bells suddenly pierced the air, and a frantic voice called out, "All to the square! All to the square!" With a mix of curiosity and concern, Faelan joined the gathering crowd, his hand resting on the gem he kept close. 

In the midst of the commotion, a mysterious figure touched Faelan's shoulder. Before he could react, a surge of energy coursed through him, and the world blurred as he lost consciousness…
`,
		},
	}

	for _, character := range characters {
		err := s.DB.Where("id = ?", character.ID).FirstOrCreate(&character).Error
		if err != nil {
			log.Printf("Error creating character with id %v in seeder: %v", character.ID, err.Error())
		}
	}
}
