package models

type CharacterInventoryItem struct {
	ID          uint `gorm:"primary_key;auto_increment"`
	CharacterId uint
	ItemId      uint
	Location    string
	Quantity    int
	Equipped    *bool

	Item Item
}
