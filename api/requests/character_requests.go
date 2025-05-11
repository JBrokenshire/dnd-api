package requests

type CreateCharacterRequest struct {
	Name     string `json:"name" validate:"required,max=200"`
	ClassId  uint   `json:"class_id" validate:"required"`
	RaceId   uint   `json:"race_id" validate:"required"`
	Pronouns string `json:"pronouns" validate:"max=32"`
	Level    int    `json:"level" validate:"required,min=1,max=20"`

	Strength     uint `json:"strength" validate:"required,min=3,max=30"`
	Dexterity    uint `json:"dexterity" validate:"required,min=3,max=30"`
	Constitution uint `json:"constitution" validate:"required,min=3,max=30"`
	Intelligence uint `json:"intelligence" validate:"required,min=3,max=30"`
	Wisdom       uint `json:"wisdom" validate:"required,min=3,max=30"`
	Charisma     uint `json:"charisma" validate:"required,min=3,max=30"`
}

type UpdateCharacterRequest struct {
	Name     string `json:"name" validate:"required,max=200"`
	ClassId  uint   `json:"class_id" validate:"required"`
	RaceId   uint   `json:"race_id" validate:"required"`
	Pronouns string `json:"pronouns" validate:"max=32"`
	Level    int    `json:"level" validate:"required,min=1,max=20"`

	Strength     uint `json:"strength" validate:"required,min=3,max=30"`
	Dexterity    uint `json:"dexterity" validate:"required,min=3,max=30"`
	Constitution uint `json:"constitution" validate:"required,min=3,max=30"`
	Intelligence uint `json:"intelligence" validate:"required,min=3,max=30"`
	Wisdom       uint `json:"wisdom" validate:"required,min=3,max=30"`
	Charisma     uint `json:"charisma" validate:"required,min=3,max=30"`
}
