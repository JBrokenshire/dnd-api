package requests

type CreateCharacterRequest struct {
	Name    string `json:"name" validate:"required,max=200"`
	ClassId uint   `json:"class_id" validate:"required"`
	RaceId  uint   `json:"race_id" validate:"required"`
}

type UpdateCharacterRequest struct {
	Name    string `json:"name" validate:"required,max=200"`
	ClassId uint   `json:"class_id" validate:"required"`
	RaceId  uint   `json:"race_id" validate:"required"`
}
