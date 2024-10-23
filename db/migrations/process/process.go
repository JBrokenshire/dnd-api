package process

import (
	"dnd-api/db/migrations/list"
	gm "github.com/ShkrutDenis/go-migrations"
	gmStore "github.com/ShkrutDenis/go-migrations/store"
)

func Run() {
	gm.Run(getMigrationsList())
}

func getMigrationsList() []gmStore.Migratable {
	return []gmStore.Migratable{
		&list.CreateClassesTable{},
		&list.CreateRacesTable{},
		&list.CreateCharactersTable{},
		&list.CharacterAbilityScoreProficiencies{},
		&list.CreateCharacterProficientSkillsTable{},
		&list.CreateCharacterSensesTable{},
		&list.CreateCharacterProficientArmourTypes{},
		&list.CreateCharacterProficientWeapons{},
		&list.CreateCharacterProficientTools{},
		&list.CreateCharacterLanguages{},
		&list.UpdateCharactersInitiativeModifier{},
		&list.CreateCharacterDefenses{},
		&list.CreateCharacterConditions{},
		&list.UpdateCharacterAttacksPerAction{},
		&list.CreateItemsTable{},
		&list.CreateWeaponsTable{},
		&list.CreateCharacterInventory{},
		&list.UpdateWeaponsAddBonus{},
		&list.CreateCharacterMoney{},
		&list.UpdateItemsAddEquippable{},
		&list.CreateArmourTable{},
		&list.CreateShieldsTable{},
		&list.UpdateInventoryItemAddType{},
		&list.UpdateCharacterAddNotes{},
		&list.CreateCharacterSkillsAdvantagesTable{},
		&list.CreateBackgroundsTable{},
		&list.UpdateCharactersAddBackground{},
		&list.UpdateCharactersAddCharacteristics{},
		&list.UpdateClassesAddSpellcastingAbility{},
		&list.CreateSpellsTable{},
		&list.CreateCharacterSpellsTable{},
		&list.CreateTraitsTable{},
		&list.CreateRaceTraitsTable{},
		&list.CreateFeaturesTable{},
		&list.CreateClassFeaturesTable{},
		&list.CreateSubclassesTable{},
		&list.CreateSubclassFeaturesTable{},
		&list.UpdateCharactersAddSubclassID{},
		&list.UpdateClassFeaturesAddLevel{},
		&list.UpdateSubclassFeaturesAddLevel{},
	}
}
