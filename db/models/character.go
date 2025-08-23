package models

const (
	AdvancementTypeMilestone = "Milestone"
	AdvancementTypeXP        = "XP"
)

var ValidAdvancementTypes = []string{
	AdvancementTypeMilestone,
	AdvancementTypeXP,
}

const (
	HitPointTypeFixed  = "Fixed"
	HitPointTypeManual = "Manual"
)

var ValidHitPointTypes = []string{
	HitPointTypeFixed,
	HitPointTypeManual,
}

type Character struct {
	ID       uint   `gorm:"primary_key;auto_increment" json:"id"`
	UserId   uint   `json:"user_id"`
	Name     string `json:"name"`
	ClassId  uint   `json:"class_id"`
	RaceId   uint   `json:"race_id"`
	Level    int    `json:"level"`
	Pronouns string `json:"pronouns"`

	// Ability Scores
	Strength     uint `json:"strength"`
	Dexterity    uint `json:"dexterity"`
	Constitution uint `json:"constitution"`
	Intelligence uint `json:"intelligence"`
	Wisdom       uint `json:"wisdom"`
	Charisma     uint `json:"charisma"`

	// Saving Throw Adjustments
	StrengthSaveAdjustment     int `json:"strength_save_adjustment"`
	DexteritySaveAdjustment    int `json:"dexterity_save_adjustment"`
	ConstitutionSaveAdjustment int `json:"constitution_save_adjustment"`
	IntelligenceSaveAdjustment int `json:"intelligence_save_adjustment"`
	WisdomSaveAdjustment       int `json:"wisdom_save_adjustment"`
	CharismaSaveAdjustment     int `json:"charisma_save_adjustment"`

	Inspiration      bool `json:"inspiration"`
	CurrentHitPoints uint `json:"current_hit_points"`
	MaxHitPoints     uint `json:"max_hit_points"`
	TempHitPoints    uint `json:"temp_hit_points"`

	// Settings
	AdvancementType string `json:"advancement_type"`
	HitPointType    string `json:"hit_point_type"`

	Class          Class `json:"class"`
	Race           Race  `json:"race"`
	ProfilePicture File  `json:"profile_picture"`
}
