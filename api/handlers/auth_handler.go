package handlers

import (
	"dnd-api/api"
	"dnd-api/api/requests"
	"dnd-api/api/responses"
	"dnd-api/config"
	"dnd-api/services/jwt_service"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
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

// RefreshToken Refresh godoc
// @Summary Refresh access token
// @Description Perform refresh access token
// @ID auth-refresh
// @Tags Auth Actions
// @Accept json
// @Produce json
// @Success 200 {object} responses.RefreshResponse
// @Failure 400 {object} responses.Error
// @Failure 401 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /auth/refresh [get]
func (h *AuthHandler) RefreshToken(c echo.Context) error {

	// Get the refresh Token if it exists from the "refresh" cookie
	cookie, err := c.Cookie("refresh")
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "No refresh cookie provided")
	}

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	if err != nil {
		return responses.ErrorResponse(c, http.StatusUnauthorized, "Invalid token: "+err.Error())
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return responses.ErrorResponse(c, http.StatusUnauthorized, "Invalid token")
	}

	// Check user exists
	user := h.server.Repos.User.GetById(claims["id"])
	if user.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "User not found")
	}

	// Create access token
	tokenService := jwt_service.NewTokenService()
	accessToken, exp, err := tokenService.CreateUserAccessToken(user, true)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "There was a problem creating the access token")
	}

	// Create refresh token
	refreshToken, refreshExpiry, err := tokenService.CreateUserRefreshToken(user)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "There was a problem creating the refresh token")
	}

	// Attach refresh cookie
	tokenService.AttachRefreshCookie(c, refreshToken, *refreshExpiry)

	res := responses.NewRefreshResponse(user, accessToken, exp)
	return responses.Response(c, http.StatusOK, res)
}

// Logout godoc
// @Summary Log User Out
// @Description Set's the refresh cookie to an expired date and clears token.
// @ID auth-logout
// @Tags Auth Actions
// @Accept json
// @Produce json
// @Success 200 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /auth/logout [get]
func (h *AuthHandler) Logout(c echo.Context) error {

	tokenService := jwt_service.NewTokenService()
	tokenService.AttachRefreshCookie(c, "LOGGED-OUT", time.Unix(0, 0))

	return responses.Response(c, http.StatusOK, "Logged Out")
}
