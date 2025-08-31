package models

const (
	SpellcastingAbilityIntelligence = "Intelligence"
	SpellcastingAbilityWisdom       = "Wisdom"
	SpellcastingAbilityCharisma     = "Charisma"
)

type Class struct {
	ID                  uint    `gorm:"primary_key;auto_increment" json:"id"`
	Name                string  `json:"name"`
	ShortDescription    string  `json:"short_description"`
	PrimaryAbility      string  `json:"primary_ability"`
	HitPointDieValue    int     `json:"hit_point_die_value"`
	Saves               string  `json:"saves"`
	SpellcastingAbility *string `json:"spellcasting_ability"`

	Logo            File               `json:"logo"`
	BackgroundImage File               `json:"background_image"`
	SpellSlots      []*ClassSpellSlots `json:"spell_slots"`
}
