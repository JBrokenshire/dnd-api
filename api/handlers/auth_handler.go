package handlers

import (
	s "dnd-api/api"
	"dnd-api/api/requests"
	"dnd-api/api/responses"
	"dnd-api/db/models"
	"dnd-api/pkg/jwt_service"
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
	server *s.Server
}

func NewAuthHandler(server *s.Server) *AuthHandler {
	return &AuthHandler{
		server: server,
	}
}

// Login godoc
// @Summary Authenticate a user
// @Description Perform user login
// @ID user-login
// @Tags Auth Actions
// @Accept json
// @Produce json
// @Param params body requests.LoginRequest true "User's credentials"
// @Success 200 {object} responses.LoginResponse
// @Failure 401 {object} responses.Error
// @Router /auth/login [post]
func (h *AuthHandler) Login(c echo.Context) error {
	request := new(requests.LoginRequest)
	if err := c.Bind(request); err != nil {
		return err
	}
	if err := c.Validate(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	user := models.User{}
	h.server.Repos.User.GetByEmail(request.Email)
	if user.ID == 0 || (bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil) {

		// In case a user was not found, run a bcrypt anyway to prevent timing attacks
		if user.ID == 0 {
			_ = bcrypt.CompareHashAndPassword([]byte("$2a$10$i2g71qLSxirZ/iMzX6snNusipVKj1Rc2ec8yfflyXfvL5GkblU5c6"), []byte("not-a-real-password"))
		}

		return responses.ErrorResponse(c, http.StatusUnauthorized, "Incorrect username or Password. Please try again.")
	}

	tokenService := jwt_service.NewTokenService()

	accessToken, exp, err := tokenService.CreateUserAccessToken(&user, true)
	if err != nil {
		log.Printf("There was an error creating an access token: %v", err)
		return responses.ErrorResponse(c, http.StatusUnauthorized, "There was a problem creating an access token")
	}
	refreshToken, refreshExpires, err := tokenService.CreateUserRefreshToken(&user, uint(10*time.Minute))
	if err != nil {
		log.Printf("There was an error creating a Refresh token: %v", err)
		return responses.ErrorResponse(c, http.StatusUnauthorized, "There was a problem creating a refresh token")
	}
	tokenService.AttachRefreshCookie(c, refreshToken, *refreshExpires)
	response := responses.NewLoginResponse(accessToken, exp, true, false)

	return responses.Response(c, http.StatusOK, response)
}

// RefreshToken Refresh godoc
// @Summary Refresh access token
// @Description Perform refresh access token
// @ID user-refresh
// @Tags Auth Actions
// @Accept json
// @Produce json
// @Success 200 {object} responses.RefreshResponse
// @Failure 401 {object} responses.Error
// @Router /auth/refresh [get]
func (h *AuthHandler) RefreshToken(c echo.Context) error {

	// Get the refresh Token if it exists from the "refresh" cookie
	cookie, err := c.Cookie("refresh")
	if err != nil {
		return responses.ErrorResponse(c, http.StatusUnauthorized, "No refresh cookie provided")
	}

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})

	if err != nil {
		return responses.ErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return responses.ErrorResponse(c, http.StatusUnauthorized, "Invalid token")
	}

	user := h.server.Repos.User.GetByUid(claims["id"].(string))
	if user.ID == 0 {
		return responses.ErrorResponse(c, http.StatusUnauthorized, "User not found")
	}

	tokenService := jwt_service.NewTokenService()
	accessToken, exp, err := tokenService.CreateUserAccessToken(user, true)
	if err != nil {
		return err
	}
	refreshToken, refreshExpiry, err := tokenService.CreateUserRefreshToken(user, uint(10*time.Minute))

	if err != nil {
		return err
	}

	res := responses.NewRefreshResponse(user, accessToken, exp)
	tokenService.AttachRefreshCookie(c, refreshToken, *refreshExpiry)

	return responses.Response(c, http.StatusOK, res)
}

// Logout godoc
// @Summary Log User Out
// @Description Set's the refresh cookie to an expired date and clears token.
// @ID user-logout
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
