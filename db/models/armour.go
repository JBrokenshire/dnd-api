package models

const (
	ArmourTypeLight  = "Light Armour"
	ArmourTypeMedium = "Medium Armour"
	ArmourTypeHeavy  = "Heavy Armour"
)

type Armour struct {
	ItemId              uint `gorm:"primary_key;foreignKey:ItemID"`
	BaseAC              int
	StrRequirement      *int
	Type                string
	MaxDexModifier      *int
	StealthDisadvantage bool
}
