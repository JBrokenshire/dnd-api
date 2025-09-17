package models

const (
	WeaponAbilitySTR = "STR"
	WeaponAbilityDEX = "DEX"
)

const (
	WeaponTypeMelee  = "Melee Weapon"
	WeaponTypeRanged = "Ranged Weapon"
)

type Weapon struct {
	ItemId        uint `json:"item_id"`
	WeaponType    string
	Distance      int
	AltDistance   *int
	Ability       string
	Damage        string
	AltDamage     *string
	DamageType    string
	Bonus         int
	Proficiencies string
}
