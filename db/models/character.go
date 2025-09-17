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
	ID           uint   `gorm:"primary_key;auto_increment" json:"id"`
	UserId       uint   `json:"user_id"`
	Name         string `json:"name"`
	ClassId      uint   `json:"class_id"`
	RaceId       uint   `json:"race_id"`
	Level        int    `json:"level"`
	Pronouns     string `json:"pronouns"`
	BackgroundId uint   `json:"background_id"`

	Alignment string `json:"alignment"`
	Gender    string `json:"gender"`
	Eyes      string `json:"eyes"`
	Size      string `json:"size"`
	Height    string `json:"height"`
	Faith     string `json:"faith"`
	Hair      string `json:"hair"`
	Skin      string `json:"skin"`
	Age       string `json:"age"`
	Weight    string `json:"weight"`

	// Ability Scores
	Strength     uint `json:"strength"`
	Dexterity    uint `json:"dexterity"`
	Constitution uint `json:"constitution"`
	Intelligence uint `json:"intelligence"`
	Wisdom       uint `json:"wisdom"`
	Charisma     uint `json:"charisma"`

	Inspiration      bool `json:"inspiration"`
	CurrentHitPoints uint `json:"current_hit_points"`
	MaxHitPoints     uint `json:"max_hit_points"`
	TempHitPoints    uint `json:"temp_hit_points"`
	AttacksPerAction int  `json:"attacks_per_action"`

	// Settings
	AdvancementType string `json:"advancement_type"`
	HitPointType    string `json:"hit_point_type"`

	Senses        string `json:"senses"`
	Proficiencies string `json:"proficiencies"`

	PersonalityTraits string `json:"personality_traits"`
	Ideals            string `json:"ideals"`
	Bonds             string `json:"bonds"`
	Flaws             string `json:"flaws"`
	Organisations     string `json:"organisations"`
	Allies            string `json:"allies"`
	Enemies           string `json:"enemies"`
	Backstory         string `json:"backstory"`

	Class            Class                       `json:"class"`
	Race             Race                        `json:"race"`
	ProfilePicture   File                        `json:"profile_picture"`
	ProficientSkills []*CharacterProficientSkill `json:"proficient_skills"`
	Defenses         []*CharacterDefense         `json:"defenses"`
	Background       Background                  `json:"background"`
	Spells           []*Spell                    `json:"spells"`
	UsedSpellSlots   []*CharacterUsedSpellSlot   `json:"used_spell_slots"`
	Inventory        []*CharacterInventoryItem   `json:"inventory"`
}
