package requests

type CreateRaceRequest struct {
	Name string `json:"name" validate:"required,max=200"`
}

type UpdateRaceRequest struct {
	Name string `json:"name" validate:"required,max=200"`
}
