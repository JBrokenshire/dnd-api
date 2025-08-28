package models

const (
	DamageTypeAcid        = "Acid"
	DamageTypeBludgeoning = "Bludgeoning"
	DamageTypeCold        = "Cold"
	DamageTypeFire        = "Fire"
	DamageTypeForce       = "Force"
	DamageTypeLightning   = "Lightning"
	DamageTypeNecrotic    = "Necrotic"
	DamageTypePiercing    = "Piercing"
	DamageTypePoison      = "Poison"
	DamageTypePsychic     = "Psychic"
	DamageTypeRadiant     = "Radiant"
	DamageTypeSlashing    = "Slashing"
	DamageTypeThunder     = "Thunder"
)

var ValidDamageTypes = []string{
	DamageTypeAcid,
	DamageTypeBludgeoning,
	DamageTypeCold,
	DamageTypeFire,
	DamageTypeForce,
	DamageTypeLightning,
	DamageTypeNecrotic,
	DamageTypePiercing,
	DamageTypePoison,
	DamageTypePsychic,
	DamageTypeRadiant,
	DamageTypeSlashing,
	DamageTypeThunder,
}

const (
	DefenseTypeResistance    = "Resistance"
	DefenseTypeImmunity      = "Immunity"
	DefenseTypeVulnerability = "Vulnerability"
)

var ValidDefenseTypes = []string{
	DefenseTypeResistance,
	DefenseTypeImmunity,
	DefenseTypeVulnerability,
}

type CharacterDefense struct {
	ID          uint   `gorm:"primary_key;auto_increment" json:"id"`
	CharacterID uint   `json:"character_id"`
	DamageType  string `json:"damage_type"`
	DefenseType string `json:"defense_type"`
}
