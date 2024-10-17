package requests

type CharacterRequest struct {
	Name              string `json:"name"`
	Level             int    `json:"level"`
	ProfilePictureURL string `json:"profile_picture_url"`
	ClassID           int    `json:"class_id"`
	RaceID            int    `json:"race_id"`

	Strength               int  `json:"strength"`
	Dexterity              int  `json:"dexterity"`
	Constitution           int  `json:"constitution"`
	Intelligence           int  `json:"intelligence"`
	Wisdom                 int  `json:"wisdom"`
	Charisma               int  `json:"charisma"`
	ProficientStrength     bool `json:"proficient_strength"`
	ProficientDexterity    bool `json:"proficient_dexterity"`
	ProficientConstitution bool `json:"proficient_constitution"`
	ProficientIntelligence bool `json:"proficient_intelligence"`
	ProficientWisdom       bool `json:"proficient_wisdom"`
	ProficientCharisma     bool `json:"proficient_charisma"`

	WalkingSpeedModifier int  `json:"walking_speed_modifier"`
	Inspiration          bool `json:"inspiration"`
	CurrentHitPoints     int  `json:"current_hit_points"`
	MaxHitPoints         int  `json:"max_hit_points"`
	TempHitPoints        int  `json:"temp_hit_points"`

	InitiativeModifier      int  `json:"initiative_modifier"`
	BaseArmourClass         int  `json:"base_armour_class"`
	ArmourClassAddDexterity bool `json:"armour_class_add_dexterity"`

	BackgroundName string `json:"background_name"`
	Alignment      string `json:"alignment"`
	Gender         string `json:"gender"`
	Eyes           string `json:"eyes"`
	Size           string `json:"size"`
	Height         string `json:"height"`
	Faith          string `json:"faith"`
	Hair           string `json:"hair"`
	Skin           string `json:"skin"`
	Age            int    `json:"age"`
	Weight         int    `json:"weight"`

	Organisations string `json:"organisations"`
	Allies        string `json:"allies"`
	Enemies       string `json:"enemies"`
	Backstory     string `json:"backstory"`

	AttacksPerAction int `json:"attacks_per_action"`
}

func NewCharacterRequest(cr *CharacterRequest) CharacterRequest {
	if cr.Name == "" {
		cr.Name = "DEFAULT NAME"
	}
	if cr.Level == 0 {
		cr.Level = 1
	}
	if cr.ClassID == 0 {
		cr.ClassID = 1
	}
	if cr.RaceID == 0 {
		cr.RaceID = 1
	}
	if cr.Strength == 0 {
		cr.Strength = 10
	}
	if cr.Dexterity == 0 {
		cr.Dexterity = 10
	}
	if cr.Constitution == 0 {
		cr.Constitution = 10
	}
	if cr.Intelligence == 0 {
		cr.Intelligence = 10
	}
	if cr.Wisdom == 0 {
		cr.Wisdom = 10
	}
	if cr.Charisma == 0 {
		cr.Charisma = 10
	}
	if cr.CurrentHitPoints == 0 {
		cr.CurrentHitPoints = 1
	}
	if cr.MaxHitPoints == 0 {
		cr.MaxHitPoints = 1
	}
	if cr.BaseArmourClass == 0 {
		cr.BaseArmourClass = 10
	}
	if cr.AttacksPerAction == 0 {
		cr.AttacksPerAction = 1
	}
	if cr.BackgroundName == "" {
		cr.BackgroundName = "Outlander"
	}
	if cr.Size == "" {
		cr.Size = "Medium"
	}
	if cr.Alignment == "" {
		cr.Alignment = "Lawful Good"
	}

	return *cr
}

func (cr CharacterRequest) IsEmpty() bool {
	if cr.Name == "" && cr.Level == 0 && cr.ClassID == 0 && cr.RaceID == 0 && cr.Strength == 0 && cr.Dexterity == 0 && cr.Constitution == 0 && cr.Intelligence == 0 && cr.Wisdom == 0 && cr.Charisma == 0 {
		return true
	}
	return false
}
