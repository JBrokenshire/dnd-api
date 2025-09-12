package models

const (
	RarityCommon    = "Common"
	RarityUncommon  = "Uncommon"
	RarityRare      = "Rare"
	RarityVeryRare  = "Very Rare"
	RarityLegendary = "Legendary"
)

const (
	ItemTypeArmour = "armour"
	ItemTypeWeapon = "weapon"
	ItemTypeItem   = "item"
)

const (
	SavingThrowBonusTypeAdvantage = "advantage"
	SavingThrowBonusTypeAdd       = "add"
)

type Item struct {
	ID         uint `gorm:"primary_key;auto_increment"`
	Name       string
	Rarity     string
	Origin     string
	Weight     float32
	Cost       float32
	Equippable bool
	Notes      string
	Type       string

	// Bonuses
	StrengthBonus         int
	DexterityBonus        int
	ConstitutionBonus     int
	IntelligenceBonus     int
	WisdomBonus           int
	CharismaBonus         int
	StrengthSaveBonus     int
	DexteritySaveBonus    int
	ConstitutionSaveBonus int
	IntelligenceSaveBonus int
	WisdomSaveBonus       int
	CharismaSaveBonus     int
	SavingThrowBonusType  *string
	SavingThrowBonusText  string
	InitiativeBonus       int
	InitiativeAdvantage   bool
	ArmourClassBonus      int

	Armour Armour
}
