package requests

type ClassCreateRequest struct {
	Name             string `json:"name" validate:"required"`
	ShortDescription string `json:"short_description" validate:"required"`
	LongDescription  string `json:"long_description" validate:"required"`
}
