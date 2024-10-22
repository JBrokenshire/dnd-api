package seeders

import (
	"dnd-api/db/models"
	"log"
)

func (s *Seeder) SetFeatures() {
	actionBloodMaledict := "Blood Maledict"
	actionRage := "Rage"
	actionHybridTransformation := "Hybrid Transformation"
	actionSecondWind := "Second Wind"
	actionActionSurge := "Action Surge"
	actionUnleashIncarnation := "Unleash Incarnation"
	actionInnateSorcery := "Innate Sorcery"
	actionFontOfMagic := "Font Of Magic"
	actionSorcerousRestoration := "Sorcerous Restoration"

	actionTypeBonus := "Bonus Action"

	actionUsesOne := 1
	actionUsesTwo := 2
	actionUsesThree := 3
	actionUsesFour := 4
	actionUsesFive := 5

	actionResetShortRest := "Short Rest"
	actionResetLongRest := "Long Rest"

	features := []models.Feature{
		{
			ID:   1,
			Name: "Ability Score Improvement",
		},
		{
			ID:   2,
			Name: "Proficiencies",
		},
		{
			ID:          3,
			Name:        "Hunter's Bane",
			Description: "You have advantage on Wisdom (Survival) checks to track fey, fiends, or undead, and you have advantage on Intelligence ability checks to recall information about them.",
		},
		{
			ID:          4,
			Name:        "Blood Maledict",
			Description: "You can invoke a blood curse a number of times per short rest based on your level. Before it affects the target, you can choose to amplify the curse. An amplified curse gains an additional effect noted in the curse's description. Amplifying a curse causes you to take 1d4 points of necrotic damage that cannot be reduced in any way.\n\nCreatures that don't have blood in their bodies are immune to blood curses, unless the curse has been amplified.",
			Action:      &actionBloodMaledict,
			ActionUses:  &actionUsesOne,
			ActionReset: &actionResetShortRest,
		},
		{
			ID:          5,
			Name:        "Blood Curses",
			Description: "Your chosen Blood Curses for use with the Blood Maledict feature.",
		},
		{
			ID:          6,
			Name:        "Fighting Style",
			Description: "You adopt a fighting style specialty.",
		},
		{
			ID:          7,
			Name:        "Crimson Rite",
			Description: "As a bonus action, you can activate any rite you know on a weapon you're holding. A weapon can only hold one active rite at a time. When activated, you take 1d4 necrotic damage which can't be reduced in any way.\n\nWhile your rite is in effect, attacks made with this weapon are magical and you deal an additional 1d4 damage of the type determined by the chosen rite. The rite’s effect lasts until you finish a short/long rest and other creatures cannot gain the benefit of your rite.",
		},
		{
			ID:   8,
			Name: "Blood Hunter Order",
		},
		{
			ID:          9,
			Name:        "Mutagencraft",
			Description: "You can create 1 mutagen(s) from the formulas you know during a short/long rest. Mutagens are designed for the specific biology of the character that concocted them, so your mutagens have no effect on other creatures. They're unstable and become inert if not used before your next short/long rest.\n\nYou can consume a mutagen with a bonus action, whose effects and side effects last until the end of your next rest. If you have at least one mutagen affecting you, you can use an action to flush all mutagens from your system, ending their effects and side effects.",
		},
		{
			ID:          10,
			Name:        "Formulas",
			Description: "You begin to uncover forbidden alchemical formulas that temporarily alter your mental and physical abilities.",
		},
		{
			ID:          11,
			Name:        "Rage",
			Description: "As a bonus action enter a rage for up to 1 minute (10 rounds).\n\nYou gain advantage on STR checks and saving throws (not attacks), +2 melee damage with STR weapons, resistance to bludgeoning, piercing, slashing damage. You can't cast or concentrate on spells while raging.\n\nYour rage ends early if you are knocked unconscious or if your turn ends and you haven’t attacked a hostile creature since your last turn or taken damage since then. You can also end your rage as a bonus action.",
			Action:      &actionRage,
			ActionType:  &actionTypeBonus,
			ActionUses:  &actionUsesThree,
			ActionReset: &actionResetLongRest,
		},
		{
			ID:          12,
			Name:        "Unarmoured Defense",
			Description: "While not wearing armor, your AC equals 10 + DEX modifier + CON modifier + any shield bonus.",
		},
		{
			ID:          13,
			Name:        "Reckless Attack",
			Description: "When you make your first attack on your turn, you can decide to attack recklessly, giving you advantage on melee weapon attack rolls using STR during this turn, but attack rolls against you have advantage until your next turn.",
		},
		{
			ID:          14,
			Name:        "Danger Sense",
			Description: "You have advantage on DEX saving throws against effects that you can see while not blinded, deafened, or incapacitated.",
		},
		{
			ID:   15,
			Name: "Primal Path",
		},
		{
			ID:          16,
			Name:        "Spirit Seeker",
			Description: "Yours is a path that seeks attunement with the natural world, giving you a kinship with beasts. At 3rd level when you adopt this path, you gain the ability to cast the Beast Sense and Speak with Animals spells, but only as rituals.",
		},
		{
			ID:   17,
			Name: "Totem Spirit",
		},
		{
			ID:          18,
			Name:        "Extra Attack",
			Description: "You can attack twice, instead of once, whenever you take the Attack action on your turn.",
		},
		{
			ID:          19,
			Name:        "Fast Movement",
			Description: "Your speed increases by 10 ft. while you aren't wearing heavy armor.",
		},
		{
			ID:          20,
			Name:        "Heightened Senses",
			Description: "You gain advantage on Wisdom (Perception) checks that rely on hearing or smell.",
		},
		{
			ID:          21,
			Name:        "Hybrid Transformation",
			Description: "As a bonus action, you transform into a special hybrid form for up to 1 hour. You can speak, use equipment, and wear armor while in this form, and can revert to your normal form as a bonus action. You automatically revert to your normal form if you fall unconscious or die.",
			Action:      &actionHybridTransformation,
			ActionType:  &actionTypeBonus,
			ActionUses:  &actionUsesOne,
			ActionReset: &actionResetShortRest,
		},
		{
			ID:          22,
			Name:        "Second Wind",
			Description: "Once per short rest, you can use a bonus action to regain 1d10 + %%level%% HP.",
			Action:      &actionSecondWind,
			ActionType:  &actionTypeBonus,
			ActionUses:  &actionUsesOne,
			ActionReset: &actionResetShortRest,
		},
		{
			ID:          23,
			Name:        "Action Surge",
			Description: "You can take one additional action on your turn. This can be used 1 times per short rest.",
			Action:      &actionActionSurge,
			ActionUses:  &actionUsesOne,
			ActionReset: &actionResetShortRest,
		},
		{
			ID:   24,
			Name: "Martial Archetype",
		},
		{
			ID:          25,
			Name:        "Manifest Echo",
			Description: "You can use a bonus action to magically manifest an echo of yourself in an unoccupied space you can see within 15 ft. of you. The echo lasts until it is destroyed, you dismiss it as a bonus action, you manifest another echo, or until you’re incapacitated.\n\nYour echo has AC 17, 1 hit point, and immunity to all conditions. If it has to make a saving throw, it uses your bonus for the roll. It is your size, and occupies its space. If your echo is ever more than 30 ft. from you at the end of your turn, it is destroyed.\n\nYou can use the echo in the following ways:\n\n- You can mentally command the echo to move up to 30 ft. in any direction (no action required)\n\n- As a bonus action you can teleport, magically swapping places with your echo at a cost of 15 ft. of your movement, regardless of the distance between the two of you.\n\n- When you take the Attack action, any attack you make can originate from your or the echo's space. You make this choice with each attack.\n\n- When a creature you can see within 5 ft. of your echo moves at least 5 ft. away from it, you can use your reaction to make an opportunity attack as if you were in the echo's space.",
		},
		{
			ID:          26,
			Name:        "Unleash Incarnation",
			Description: "You can heighten your echo’s fury. Whenever you take the Attack action, you can make one additional melee attack from the echo’s position.\n\nYou can use this feature 4 time(s). You regain all expended uses when you finish a long rest.",
			Action:      &actionUnleashIncarnation,
			ActionUses:  &actionUsesFour,
			ActionReset: &actionResetLongRest,
		},
		{
			ID:          27,
			Name:        "Spellcasting",
			Description: "Drawing from your innate magic, you can cast spells.",
		},
		{
			ID:          28,
			Name:        "Innate Sorcery",
			Description: "Twice per Long Rest, you can take a Bonus Action to unleash the simmering magic within you for 1 minute.",
			Action:      &actionInnateSorcery,
			ActionType:  &actionTypeBonus,
			ActionUses:  &actionUsesTwo,
			ActionReset: &actionResetLongRest,
		},
		{
			ID:          29,
			Name:        "Font of Magic",
			Description: "You can tap into the wellspring of magic within yourself. This wellspring is represented by Sorcery Points, which allow you to create a variety of magical effects.",
			Action:      &actionFontOfMagic,
			ActionUses:  &actionUsesFive,
			ActionReset: &actionResetLongRest,
		},
		{
			ID:          30,
			Name:        "Metamagic",
			Description: "You can alter spells to suit your needs; you know 2 Metamagic options which can be used to temporarily modify spells you cast.",
		},
		{
			ID:   31,
			Name: "Metamagic Options",
		},
		{
			ID:   32,
			Name: "Sorcerer Subclass",
		},
		{
			ID:          33,
			Name:        "Sorcerous Restoration",
			Description: "When you finish a Short Rest, you can regain up to 2 Sorcery Points. Once used, you can’t use this feature again until you finish a Long Rest.",
			Action:      &actionSorcerousRestoration,
			ActionUses:  &actionUsesOne,
			ActionReset: &actionResetLongRest,
		},
	}

	for _, feature := range features {
		err := s.DB.Where("id = ?", feature.ID).FirstOrCreate(&feature).Error
		if err != nil {
			log.Printf("error creating feature with id %v - %v", feature.ID, err)
		}
	}
}
