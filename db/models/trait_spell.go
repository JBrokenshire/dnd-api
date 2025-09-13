package models

const (
	RestTypeLong  = "Long Rest"
	RestTypeShort = "Short Rest"
)

type TraitSpell struct {
	ID      uint `gorm:"primary_key;auto_increment"`
	TraitId string
	SpellId uint
	Uses    *int
	Reset   *string

	Spell Spell
}
