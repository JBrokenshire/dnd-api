package models

const (
	SkillAcrobatics     = "Acrobatics"
	SkillAnimalHandling = "Animal Handling"
	SkillArcana         = "Arcana"
	SkillAthletics      = "Athletics"
	SkillDeception      = "Deception"
	SkillHistory        = "History"
	SkillInsight        = "Insight"
	SkillIntimidation   = "Intimidation"
	SkillInvestigation  = "Investigation"
	SkillMedicine       = "Medicine"
	SkillNature         = "Nature"
	SkillPerception     = "Perception"
	SkillPerformance    = "Performance"
	SkillPersuasion     = "Persuasion"
	SkillReligion       = "Religion"
	SkillSleightOfHand  = "Sleight of Hand"
	SkillStealth        = "Stealth"
	SkillSurvival       = "Survival"
)

var ValidSkills = []string{
	SkillAcrobatics,
	SkillAnimalHandling,
	SkillArcana,
	SkillAthletics,
	SkillDeception,
	SkillHistory,
	SkillInsight,
	SkillIntimidation,
	SkillInvestigation,
	SkillMedicine,
	SkillNature,
	SkillPerception,
	SkillPerformance,
	SkillPersuasion,
	SkillReligion,
	SkillSleightOfHand,
	SkillStealth,
	SkillSurvival,
}

const (
	ProficiencyTypeHalf        = "Half Proficiency"
	ProficiencyTypeProficiency = "Proficiency"
	ProficiencyTypeExpertise   = "Expertise"
)

var ValidProficiencyTypes = []string{
	ProficiencyTypeHalf,
	ProficiencyTypeProficiency,
	ProficiencyTypeExpertise,
}

type CharacterProficientSkill struct {
	ID              uint   `gorm:"primary_key;auto_increment" json:"id"`
	CharacterID     uint   `json:"character_id"`
	Skill           string `json:"skill"`
	ProficiencyType string `json:"proficiency_type"`
}
