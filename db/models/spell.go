package models

const (
	MagicSchoolAbjuration    = "Abjuration"
	MagicSchoolConjuration   = "Conjuration"
	MagicSchoolDivination    = "Divination"
	MagicSchoolEnchantment   = "Enchantment"
	MagicSchoolEvocation     = "Evocation"
	MagicSchoolIllusion      = "Illusion"
	MagicSchoolNecromancy    = "Necromancy"
	MagicSchoolTransmutation = "Transmutation"
)

var ValidMagicSchools = []string{
	MagicSchoolAbjuration,
	MagicSchoolConjuration,
	MagicSchoolDivination,
	MagicSchoolEnchantment,
	MagicSchoolEvocation,
	MagicSchoolIllusion,
	MagicSchoolNecromancy,
	MagicSchoolTransmutation,
}

type Spell struct {
	ID          uint    `gorm:"primary_key;auto_increment" json:"id"`
	Name        string  `json:"name"`
	School      string  `json:"school"`
	Level       int     `json:"level"`
	CastingTime string  `json:"casting_time"`
	Distance    string  `json:"range"`
	IsAttack    bool    `json:"is_attack"`
	IsSave      bool    `json:"is_save"`
	SaveAbility *string `json:"save_ability"`
	Effect      string  `json:"effect" gorm:"default:'Damage'"`
	Damage      *string `json:"damage"`
	DamageType  *string `json:"damage_type"`
	Notes       string  `json:"notes"`
	CanUpcast   bool    `json:"can_upcast"`
}
