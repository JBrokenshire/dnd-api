package requests

type CreateClassRequest struct {
	Name string `json:"name" validate:"required,max=200"`
}

type UpdateClassRequest struct {
	Name string `json:"name" validate:"required,max=200"`
}
