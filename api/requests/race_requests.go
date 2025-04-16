package requests

type CreateRaceRequest struct {
	Name string `json:"name" validate:"required,max=200" example:"Barbarian"`
}

type UpdateRaceRequest struct {
	Name string `json:"name" validate:"required,max=200" example:"Barbarian"`
}
