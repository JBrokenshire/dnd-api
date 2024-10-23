package stores

import "github.com/jinzhu/gorm"

type Stores struct {
	Db                        *gorm.DB
	Character                 *GormCharacterStore
	Class                     *GormClassStore
	Subclass                  *GormSubclassStore
	Race                      *GormRaceStore
	Background                *GormBackgroundStore
	CharacterSkills           *GormCharacterSkillsStore
	CharacterSenses           *GormCharacterSensesStore
	CharacterProficiencies    *GormCharacterProficienciesStore
	CharacterDefenses         *GormCharacterDefensesStore
	CharacterConditions       *GormCharacterConditionsStore
	Item                      *GormItemStore
	Weapon                    *GormWeaponStore
	CharacterInventory        *GormCharacterInventoryStore
	CharacterMoney            *GormCharacterMoneyStore
	Armour                    *GormArmourStore
	CharacterSkillsAdvantages *GormCharacterSkillsAdvantagesStore
	CharacterSpells           *GormCharacterSpellsStore
	CharacterFeatures         *GormCharacterFeaturesStore
}

func NewStores(db *gorm.DB) *Stores {
	return &Stores{
		Db:                        db,
		Character:                 NewGormCharacterStore(db),
		Class:                     NewGormClassStore(db),
		Subclass:                  NewGormSubclassStore(db),
		Race:                      NewGormRaceStore(db),
		Background:                NewGormBackgroundStore(db),
		CharacterSkills:           NewGormCharacterSkillsStore(db),
		CharacterSenses:           NewGormCharacterSensesStore(db),
		CharacterProficiencies:    NewGormCharacterProficienciesStore(db),
		CharacterDefenses:         NewGormCharacterDefensesStore(db),
		CharacterConditions:       NewGormCharacterConditionsStore(db),
		Item:                      NewGormItemStore(db),
		Weapon:                    NewGormWeaponsStore(db),
		CharacterInventory:        NewGormCharacterInventoryStore(db),
		CharacterMoney:            NewGormCharacterMoneyStore(db),
		Armour:                    NewGormArmourStore(db),
		CharacterSkillsAdvantages: NewGormCharacterSkillsAdvantagesStore(db),
		CharacterSpells:           NewGormCharacterSpellsStore(db),
		CharacterFeatures:         NewGormCharacterFeaturesStore(db),
	}
}
