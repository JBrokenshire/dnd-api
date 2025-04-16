package requests

type CreateClassRequest struct {
	Name string `json:"name" validate:"required,max=200" example:"Barbarian"`
}

type UpdateClassRequest struct {
	Name string `json:"name" validate:"required,max=200" example:"Barbarian"`
}
