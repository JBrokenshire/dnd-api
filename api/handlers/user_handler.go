package handlers

import (
	"dnd-api/api"
	"dnd-api/api/requests"
	"dnd-api/api/responses"
	m "dnd-api/db/models"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserHandler struct {
	server *api.Server
}

func NewUserHandler(server *api.Server) *UserHandler {
	return &UserHandler{server: server}
}

// Create godoc
// @Summary Create user
// @Description Create user
// @ID users-create
// @Tags User Auth Actions
// @Accept json
// @Produce json
// @Param params body requests.CreateUserRequest true "User information"
// @Success 201 {object} responses.LoginUserResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /users [post]
func (h *UserHandler) Create(c echo.Context) error {
	request := new(requests.CreateUserRequest)
	if err := c.Bind(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid: "+err.Error())
	}
	if request.Password != request.ConfirmPassword {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Passwords do not match")
	}

	existingUser := h.server.Repos.User.GetByUsername(request.Username)
	if existingUser.ID != 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "That username is already in use")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong encrypting the password")
	}

	user := &m.User{
		Username: request.Username,
		Password: string(hashedPassword),
	}

	err = h.server.Repos.User.Create(user)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong creating the user")
	}

	res := responses.NewLoginUserResponse(user)
	return responses.Response(c, http.StatusCreated, res)
}
