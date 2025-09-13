package models

type CharacterUsedSpellSlot struct {
	ID             uint `gorm:"primary_key;auto_increment"`
	CharacterId    uint
	SpellLevel     int
	SpellSlotsUsed int
}
