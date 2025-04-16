package responses

import m "dnd-api/db/models"

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
	Exp         int64  `json:"exp"`
	Authorised  bool   `json:"authorised"`
	Enabled     bool   `json:"enabled"` // Does the user have 2FA enabled?
}

type RefreshResponse struct {
	AccessToken string            `json:"access_token"`
	Exp         int64             `json:"exp"`
	UserData    LoginUserResponse `json:"user_data,omitempty"`
}

type LoginUserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

func NewLoginResponse(token string, exp int64, authorised, enabled bool) *LoginResponse {
	return &LoginResponse{
		AccessToken: token,
		Exp:         exp,
		Authorised:  authorised,
		Enabled:     enabled,
	}
}

func NewRefreshResponse(user *m.User, token string, exp int64) *RefreshResponse {
	return &RefreshResponse{
		AccessToken: token,
		Exp:         exp,
		UserData:    *NewLoginUserResponse(user),
	}
}

func NewLoginUserResponse(user *m.User) *LoginUserResponse {
	return &LoginUserResponse{
		ID:       user.ID,
		Username: user.Username,
	}
}
