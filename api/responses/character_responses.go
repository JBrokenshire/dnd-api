package responses

import (
	m "dnd-api/db/models"
)

type CharacterResponse struct {
	ID       uint   `json:"id"`
	UserId   uint   `json:"user_id"`
	Name     string `json:"name"`
	ClassId  uint   `json:"class_id"`
	RaceId   uint   `json:"race_id"`
	Pronouns string `json:"pronouns"`
	Level    int    `json:"level"`

	// Details
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

	Class            ClassResponse                      `json:"class"`
	Race             RaceResponse                       `json:"race"`
	ProfilePicture   FileResponse                       `json:"profile_picture"`
	ProficientSkills []CharacterProficientSkillResponse `json:"proficient_skills"`
	Defenses         []CharacterDefenseResponse         `json:"defenses"`
	Background       BackgroundResponse                 `json:"background"`
	Spells           []SpellResponse                    `json:"spells"`
	UsedSpellSlots   []CharacterUsedSpellSlotResponse   `json:"used_spell_slots"`
	Inventory        []CharacterInventoryItemResponse   `json:"inventory"`
}

type SlimCharacterResponse struct {
	ID       uint   `json:"id"`
	UserId   uint   `json:"user_id"`
	Name     string `json:"name"`
	ClassId  uint   `json:"class_id"`
	RaceId   uint   `json:"race_id"`
	Pronouns string `json:"pronouns"`
	Level    int    `json:"level"`

	Class          ClassResponse `json:"class"`
	Race           RaceResponse  `json:"race"`
	ProfilePicture FileResponse  `json:"profile_picture"`
}

type CharacterPaginatedResponse struct {
	Data []SlimCharacterResponse `json:"data"`
	Meta ResponseMeta            `json:"meta"`
}

func NewCharacterResponse(character *m.Character) *CharacterResponse {
	res := &CharacterResponse{
		ID:       character.ID,
		UserId:   character.UserId,
		Name:     character.Name,
		ClassId:  character.ClassId,
		RaceId:   character.RaceId,
		Pronouns: character.Pronouns,
		Level:    character.Level,

		Alignment: character.Alignment,
		Gender:    character.Gender,
		Eyes:      character.Eyes,
		Size:      character.Size,
		Height:    character.Height,
		Faith:     character.Faith,
		Hair:      character.Hair,
		Skin:      character.Skin,
		Age:       character.Age,
		Weight:    character.Weight,

		Strength:     character.Strength,
		Dexterity:    character.Dexterity,
		Constitution: character.Constitution,
		Intelligence: character.Intelligence,
		Wisdom:       character.Wisdom,
		Charisma:     character.Charisma,

		Inspiration:      character.Inspiration,
		CurrentHitPoints: character.CurrentHitPoints,
		MaxHitPoints:     character.MaxHitPoints,
		TempHitPoints:    character.TempHitPoints,
		AttacksPerAction: character.AttacksPerAction,

		AdvancementType: character.AdvancementType,
		HitPointType:    character.HitPointType,

		Senses:        character.Senses,
		Proficiencies: character.Proficiencies,

		PersonalityTraits: character.PersonalityTraits,
		Ideals:            character.Ideals,
		Bonds:             character.Bonds,
		Flaws:             character.Flaws,
		Organisations:     character.Organisations,
		Allies:            character.Allies,
		Enemies:           character.Enemies,
		Backstory:         character.Backstory,
	}

	if character.Class.ID != 0 {
		res.Class = *NewClassResponse(&character.Class)
	}
	if character.Race.ID != 0 {
		res.Race = *NewRaceResponse(&character.Race)
	}
	if character.ProfilePicture.ID != 0 {
		res.ProfilePicture = *NewFileResponse(&character.ProfilePicture)
	}
	if len(character.ProficientSkills) > 0 {
		res.ProficientSkills = NewCharacterProficientSkillResponses(character.ProficientSkills)
	}
	if len(character.Defenses) > 0 {
		res.Defenses = NewCharacterDefenseResponses(character.Defenses)
	}
	if character.Background.ID != 0 {
		res.Background = *NewBackgroundResponse(&character.Background)
	}
	if len(character.Spells) > 0 {
		res.Spells = NewSpellResponses(character.Spells)
	}
	if len(character.UsedSpellSlots) > 0 {
		res.UsedSpellSlots = NewCharacterUsedSpellSlotResponses(character.UsedSpellSlots)
	}
	if len(character.Inventory) > 0 {
		res.Inventory = NewCharacterInventoryItemResponses(character.Inventory)
	}

	return res
}

func NewSlimCharacterResponse(character *m.Character) *SlimCharacterResponse {
	res := &SlimCharacterResponse{
		ID:       character.ID,
		UserId:   character.UserId,
		Name:     character.Name,
		ClassId:  character.ClassId,
		RaceId:   character.RaceId,
		Pronouns: character.Pronouns,
		Level:    character.Level,
	}

	if character.Class.ID != 0 {
		res.Class = *NewClassResponse(&character.Class)
	}
	if character.Race.ID != 0 {
		res.Race = *NewRaceResponse(&character.Race)
	}
	if character.ProfilePicture.ID != 0 {
		res.ProfilePicture = *NewFileResponse(&character.ProfilePicture)
	}

	return res
}

func NewCharacterResponses(characters []*m.Character) []SlimCharacterResponse {
	var res []SlimCharacterResponse
	for _, character := range characters {
		res = append(res, *NewSlimCharacterResponse(character))
	}
	return res
}

func NewCharacterPaginatedResponse(characters []*m.Character, count, page, pageSize int) *CharacterPaginatedResponse {
	return &CharacterPaginatedResponse{
		Data: NewCharacterResponses(characters),
		Meta: ResponseMeta{
			TotalCount: count,
			Page:       page,
			PageSize:   pageSize,
		},
	}
}
