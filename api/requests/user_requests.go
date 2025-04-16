package requests

type CreateUserRequest struct {
	Username        string `json:"username" validate:"required,max=200"`
	Password        string `json:"password" validate:"required,max=72"`
	ConfirmPassword string `json:"confirm_password" validate:"required,max=72"`
}
