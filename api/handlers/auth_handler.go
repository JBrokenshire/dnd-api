package handlers

import (
	"dnd-api/api"
	"dnd-api/api/requests"
	"dnd-api/api/responses"
	"dnd-api/config"
	"dnd-api/services/jwt_service"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

type AuthHandler struct {
	server *api.Server
}

func NewAuthHandler(server *api.Server) *AuthHandler {
	return &AuthHandler{server: server}
}

// Login godoc
// @Summary User Login
// @Description Log a user in
// @ID auth-login
// @Tags Auth Actions
// @Accept json
// @Produce json
// @Param params body requests.LoginRequest true "User's credentials"
// @Success 200 {object} responses.LoginResponse
// @Failure 400 {object} responses.Error
// @Failure 401 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Router /auth/login [post]
func (h *AuthHandler) Login(c echo.Context) error {

	request := new(requests.LoginRequest)
	if err := c.Bind(request); err != nil {
		return err
	}
	if err := c.Validate(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid: "+err.Error())
	}

	// check the last month first
	since := time.Now().Add(-1 * 30 * 24 * time.Hour)
	failedAttempts := h.server.Repos.User.BruteForceCount(request.Username, &since)
	if failedAttempts >= config.Get().BruteForceMonthlyLimit {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Too many attempts in the past month, your account has been blocked temporarily.")
	}

	user := h.server.Repos.User.GetByUsername(request.Username)
	if user.ID == 0 {
		// In case a user was not found, run a bcrypt anyway to prevent timing attacks
		_ = bcrypt.CompareHashAndPassword([]byte("$2a$10$i2g71qLSxirZ/iMzX6snNusipVKj1Rc2ec8yfflyXfvL5GkblU5c6"), []byte("not-a-real-password"))

		return responses.ErrorResponse(c, http.StatusNotFound, "User not found")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		return responses.ErrorResponse(c, http.StatusUnauthorized, "Incorrect password. Please try again.")
	}

	tokenService := jwt_service.NewTokenService()

	accessToken, exp, err := tokenService.CreateUserAccessToken(user, true)
	if err != nil {
		log.Printf("There was an error creating an access token: %v", err)
		return responses.ErrorResponse(c, http.StatusUnauthorized, "There was a problem creating an access token")
	}
	refreshToken, refreshExpires, err := tokenService.CreateUserRefreshToken(user)
	if err != nil {
		log.Printf("There was an error creating a Refresh token: %v", err)
		return responses.ErrorResponse(c, http.StatusUnauthorized, "There was a problem creating a refresh token")
	}
	tokenService.AttachRefreshCookie(c, refreshToken, *refreshExpires)

	res := responses.NewLoginResponse(accessToken, exp, true, false)
	return responses.Response(c, http.StatusOK, res)
}
