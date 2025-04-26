package requests

type CreateSubclassRequest struct {
	ClassId          uint   `json:"class_id" validate:"required"`
	Name             string `json:"name" validate:"required,max=200"`
	ShortDescription string `json:"short_description" validate:"required"`
}

type UpdateSubclassRequest struct {
	Name             string `json:"name" validate:"required,max=200"`
	ShortDescription string `json:"short_description" validate:"required"`
}
