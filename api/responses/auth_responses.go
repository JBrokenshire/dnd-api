package responses

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
	Exp         int64  `json:"exp"`
	Authorised  bool   `json:"authorised"`
	Enabled     bool   `json:"enabled"` // Does the user have 2FA enabled?
}

func NewLoginResponse(token string, exp int64, authorised, enabled bool) *LoginResponse {
	return &LoginResponse{
		AccessToken: token,
		Exp:         exp,
		Authorised:  authorised,
		Enabled:     enabled,
	}
}
