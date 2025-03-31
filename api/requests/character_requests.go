package requests

type CreateCharacterRequest struct {
	Name    string `json:"name" validate:"required,max=200"`
	RaceId  int    `json:"race_id" validate:"required"`
	ClassId int    `json:"class_id" validate:"required"`
}

type UpdateCharacterRequest struct {
	Name    string `json:"name" validate:"required,max=200"`
	RaceId  int    `json:"race_id" validate:"required"`
	ClassId int    `json:"class_id" validate:"required"`
}
