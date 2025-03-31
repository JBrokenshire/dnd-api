package requests

type LoginRequest struct {
	Email    string `json:"email" validate:"required,max=200" example:"example@email.com"`
	Password string `json:"password" validate:"required,max=472" example:"Example1$"`
}

type AuthoriseRequest struct {
	OTP string `json:"otp" validate:"required" example:"895741"`
}

type RefreshRequest struct {
	Token string `json:"token" validate:"required" example:"refresh_token"`
}

type CreateUserRequest struct {
	Email string `json:"email" validate:"required,email,max=200" example:"example@email.com"`
	Name  string `json:"name" validate:"required,max=500" example:"Matt Nelson"`
	Roles []uint `json:"roles" validate:"required" example:"1,2"`
}

type UpdateUserRequest struct {
	Email    string `json:"email" validate:"required,email,max=200" example:"example@email.com"`
	Name     string `json:"name" validate:"required,max=500" example:"Matt Nelson"`
	Roles    []uint `json:"roles" validate:"required" example:"1,2"`
	JobTitle string `json:"job_title" validate:"max=500" example:"Junior Developer"`
	Pronouns string `json:"pronouns" validate:"max=500" example:"he/him"`
}

type UpdateSelfRequest struct {
	Email    string `json:"email" validate:"required,email,max=200" example:"example@email.com"`
	Name     string `json:"name" validate:"required,max=500" example:"Matt Nelson"`
	JobTitle string `json:"job_title" validate:"max=500" example:"Junior Developer"`
	Pronouns string `json:"pronouns" validate:"max=500" example:"he/him"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email,max=200" example:"example@email.com"`
}

type PasswordResetRequest struct {
	Token              string `json:"token" validate:"required,max=100" example:"1234abcd"`
	NewPassword        string `json:"new_password" validate:"required,passLen,passComplexity,max=472" example:"1234abcd"`
	NewPasswordConfirm string `json:"new_password_confirm" validate:"required,max=472" example:"1234abcd"`
}

type PasswordUpdateRequest struct {
	NewPassword        string `json:"new_password" validate:"required,passLen,passComplexity,max=472" example:"1234abcd"`
	NewPasswordConfirm string `json:"new_password_confirm" validate:"required,max=472" example:"1234abcd"`
}

type OwnPasswordResetRequest struct {
	OldPassword        string `json:"old_password" validate:"required,max=472" example:"Example1$"`
	NewPassword        string `json:"new_password" validate:"required,passLen,passComplexity,max=472" example:"1234abcd"`
	NewPasswordConfirm string `json:"new_password_confirm" validate:"required,max=472" example:"1234abcd"`
}
