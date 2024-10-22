package seeders

import (
	"dnd-api/db/models"
	"log"
)

func (s *Seeder) SetTraits() {
	actionStonesEndurance := "Stone's Endurance"
	actionBreathWeapon := "Breath Weapon"

	actionUsesOne := 1

	resetShortRest := "Short Rest"

	actionTypeAction := "Action"
	actionTypeReaction := "Reaction"

	traits := []models.Trait{
		{
			ID:          1,
			Name:        "Lucky",
			Description: "When you roll a 1 on the d20 for an attack roll, ability check, or saving throw, you can reroll the die and must use the new roll.",
		},
		{
			ID:          2,
			Name:        "Brave",
			Description: "You have advantage on saving throws against being frightened.",
		},
		{
			ID:          3,
			Name:        "Halfling Nimbleness",
			Description: "You can move through the space of any creature that is of a size larger than yours.",
		},
		{
			ID:          4,
			Name:        "Stout Resilience",
			Description: "You have advantage on saving throws against poison, and you have resistance against poison damage.",
		},
		{
			ID:          5,
			Name:        "Darkvision",
			Description: "You can see in darkness (shades of gray) up to 60 ft.",
		},
		{
			ID:          6,
			Name:        "Dwarven Resilience",
			Description: "You have advantage on saves against poison and resistance against poison damage.",
		},
		{
			ID:          7,
			Name:        "Tool Proficiency",
			Description: "You gain proficiency with your choice of smith’s tools, brewer’s supplies, or mason’s tools.",
		},
		{
			ID:          8,
			Name:        "Stonecunning",
			Description: "Whenever you make an Intelligence (History) check related to the origin of stonework, you are considered proficient in the History skill and add double your proficiency bonus to the check.",
		},
		{
			ID:          9,
			Name:        "Dwarven Toughness",
			Description: "Your hit point maximum increases by 1, and it increases by 1 every time you gain a level.",
		},
		{
			ID:          10,
			Name:        "Natural Athlete",
			Description: "You have proficiency in the Athletics skill.",
		},
		{
			ID:          11,
			Name:        "Stone's Endurance",
			Description: "As a reaction, reduce damage dealt to you by 1d12 %%modifier:con%% once per short rest.",
			Action:      &actionStonesEndurance,
			ActionType:  &actionTypeReaction,
			ActionUses:  &actionUsesOne,
			ActionReset: &resetShortRest,
		},
		{
			ID:          12,
			Name:        "Powerful Build",
			Description: "You count as one size larger when determining your carrying capacity and the weight you can push, drag, or lift.",
		},
		{
			ID:          13,
			Name:        "Mountain Born",
			Description: "You don't suffer the penalties for being in high altitudes, and have resistance to cold damage.",
		},
		{
			ID:          14,
			Name:        "Ability Score Increase",
			Description: "Your Strength score increases by 2, and your Charisma score increases by 1.",
		},
		{
			ID:          15,
			Name:        "Draconic Ancestry",
			Description: "You gain a breath weapon and damage resistance with your chosen dragon type.",
		},
		{
			ID:          16,
			Name:        "Breath Weapon",
			Description: "Once per short rest as an action, exhale destructive energy based on your Draconic Ancestry. Each creature in the area must make a DC %%savedc:con%% saving throw (type determined by your ancestry), taking 2d6 ([6th] 3d6, [11th] 4d6, [16th] 5d6) on a failed save, and half damage on a successful one.",
			Action:      &actionBreathWeapon,
			ActionType:  &actionTypeAction,
			ActionUses:  &actionUsesOne,
			ActionReset: &resetShortRest,
		},
		{
			ID:          17,
			Name:        "Damage Resistance",
			Description: "You have resistance to the damage type associated with your draconic ancestry.",
		},
		{
			ID:          18,
			Name:        "Ability Score Increase",
			Description: "Your Dexterity score increases by 2.",
		},
		{
			ID:          19,
			Name:        "Ability Score Increase",
			Description: "Your Charisma score increases by 1.",
		},
		{
			ID:          20,
			Name:        "Naturally Stealthy",
			Description: "You can attempt to hide even when you are obscured only by a creature that is at least one size larger than you.",
		},
	}

	for _, trait := range traits {
		err := s.DB.Where("id = ?", trait.ID).FirstOrCreate(&trait).Error
		if err != nil {
			log.Printf("error create trait with id %v - %v", trait.ID, err)
		}
	}
}
