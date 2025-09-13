package models

type Trait struct {
	ID          string `gorm:"primary_key"`
	Name        string
	Description string

	TraitSpells []*TraitSpell
	Options     []*TraitOption
}
