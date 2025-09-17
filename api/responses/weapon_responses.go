package responses

import m "dnd-api/db/models"

type WeaponResponse struct {
	WeaponType    string  `json:"weapon_type"`
	Distance      int     `json:"distance"`
	AltDistance   *int    `json:"alt_distance"`
	Ability       string  `json:"ability"`
	Damage        string  `json:"damage"`
	AltDamage     *string `json:"alt_damage"`
	DamageType    string  `json:"damage_type"`
	Bonus         int     `json:"bonus"`
	Proficiencies string  `json:"proficiencies"`
}

func NewWeaponResponse(weapon *m.Weapon) *WeaponResponse {
	return &WeaponResponse{
		WeaponType:    weapon.WeaponType,
		Distance:      weapon.Distance,
		AltDistance:   weapon.AltDistance,
		Ability:       weapon.Ability,
		Damage:        weapon.Damage,
		AltDamage:     weapon.AltDamage,
		DamageType:    weapon.DamageType,
		Bonus:         weapon.Bonus,
		Proficiencies: weapon.Proficiencies,
	}
}
