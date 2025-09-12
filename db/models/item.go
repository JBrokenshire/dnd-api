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

	Armour Armour
}
