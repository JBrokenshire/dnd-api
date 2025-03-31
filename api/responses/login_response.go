package responses

import m "dnd-api/db/models"

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
	Exp         int64  `json:"exp"`
	Authorised  bool   `json:"authorised"`
	Enabled     bool   `json:"enabled"` // Does the user have 2FA enabled?
}

type RefreshResponse struct {
	AccessToken string        `json:"accessToken"`
	Exp         int64         `json:"exp"`
	UserData    LoginUserData `json:"userData,omitempty"`
}

type LoginUserData struct {
	Uid         string               `json:"uid" example:"4758ad4a-73ea-4d91-b6bb-eca1fd12f015"`
	Email       string               `json:"email" example:"matt@example.com"`
	Name        string               `json:"name" example:"Matt Davidson"`
	Pronouns    string               `json:"pronouns" example:"he/him"`
	Permissions []PermissionResponse `json:"permissions"`
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
		UserData:    getLoginUserData(user),
	}
}

func NewProfileResponse(user *m.User) *LoginUserData {
	result := getLoginUserData(user)
	return &result
}

// GetLoginUserData returns the right response for a login. This includes a filtered user data set including enterprise
// permissions. These should have been preloaded. If they haven't been this function will just return them as blank.
func getLoginUserData(u *m.User) LoginUserData {

	permissions := u.GetPermissionList()
	permissionRes := make([]PermissionResponse, 0)
	for _, p := range permissions {
		permissionRes = append(permissionRes, PermissionResponse{
			Subject: string(p.Subject),
			Action:  string(p.Action),
		})
	}

	res := LoginUserData{
		Uid:         u.Uid,
		Email:       u.Email,
		Name:        u.Name,
		Pronouns:    u.Pronouns,
		Permissions: permissionRes,
	}
	
	return res
}
