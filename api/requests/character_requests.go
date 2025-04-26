package requests

type CreateCharacterRequest struct {
	Name     string `json:"name" validate:"required,max=200"`
	ClassId  uint   `json:"class_id" validate:"required"`
	RaceId   uint   `json:"race_id" validate:"required"`
	Pronouns string `json:"pronouns" validate:"max=32"`
	Level    int    `json:"level" validate:"required,min=1,max=20"`
}

type UpdateCharacterRequest struct {
	Name     string `json:"name" validate:"required,max=200"`
	ClassId  uint   `json:"class_id" validate:"required"`
	RaceId   uint   `json:"race_id" validate:"required"`
	Pronouns string `json:"pronouns" validate:"max=32"`
	Level    int    `json:"level" validate:"required,min=1,max=20"`
}
