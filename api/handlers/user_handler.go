package handlers

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"net/http"
	s "purplevisits.com/mdm/api"
	"purplevisits.com/mdm/api/requests"
	"purplevisits.com/mdm/api/responses"
	"purplevisits.com/mdm/config"
	"purplevisits.com/mdm/db/models"
	"purplevisits.com/mdm/pkg/rand"
	"purplevisits.com/mdm/services"
	"time"
)

type UserHandler struct {
	server  *s.Server
	service *services.UserService
}

func NewUserHandler(server *s.Server) *UserHandler {
	uh := &UserHandler{server: server}
	uh.service = services.NewUserService(server.Db, config.Get().HashCost)
	return uh
}

// Create godoc
// @Summary Create
// @Description Create a new User
// @ID user-register
// @Tags User Actions
// @Accept json
// @Produce json
// @Param params body requests.CreateUserRequest true "User's email, user's password"
// @Success 201 {object} responses.UserResponse
// @Failure 400 {object} responses.Error
// @Security ApiKeyAuth
// @Router /users [post]
func (h *UserHandler) Create(c echo.Context) error {
	createRequest := new(requests.CreateUserRequest)
	if err := c.Bind(createRequest); err != nil {
		return err
	}
	if err := c.Validate(createRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	existUser := models.User{}
	h.server.Repos.User.GetUserByEmailWithDeleted(&existUser, createRequest.Email)
	if existUser.ID != 0 {
		message := "Email already in use"
		if existUser.DeletedAt != nil {
			message = "User already exists, but has been deleted"
		}
		return responses.ErrorResponse(c, http.StatusBadRequest, message)
	}

	// Get the enterprise
	enterpriseId := c.Get("enterpriseId").(string)

	var roles []models.Role
	h.server.Repos.Role.GetRolesByIds(&roles, enterpriseId, createRequest.Roles)
	if len(createRequest.Roles) != len(roles) {
		return responses.ErrorResponse(c, http.StatusBadRequest, "You do not have access to a role selected")
	}

	user := &models.User{}
	if err := h.service.Create(enterpriseId, createRequest, user); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong storing the user in the database.")
	}

	currentUser := c.Get("currentUser").(*models.User)
	models.CreatePublicActivity(h.server.Db, c.RealIP(), enterpriseId, currentUser.ID, user.ID, "User", "User Created", currentUser.SuperAdmin)

	h.server.Db.Model(user).Association("Roles").Append(roles)

	// Generate a reset password token
	reset := &models.PasswordReset{
		Token:      rand.StringBase64Fixed(32),
		UserUid:    user.Uid,
		ValidUntil: time.Now().Add(time.Hour * 720),
	}
	if err := h.server.Repos.PasswordReset.Create(reset); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong inviting the user.")
	}

	// Send welcome email
	h.server.Dependencies.GetMailer().SendWelcomeEmail(user, reset.Token)

	response := responses.NewUserResponse(user)
	return responses.Response(c, http.StatusCreated, response)
}

// Update godoc
// @Summary Update users
// @Description Update User
// @ID users-update
// @Tags User Actions
// @Accept json
// @Produce json
// @Param uid path string true "User UID"
// @Param params body requests.UpdateUserRequest true "User email and name"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Security ApiKeyAuth
// @Router /users/{uid} [put]
func (h *UserHandler) Update(c echo.Context) error {
	updateUserRequest := new(requests.UpdateUserRequest)
	uid := c.Param("uid")

	if err := c.Bind(updateUserRequest); err != nil {
		return err
	}
	if err := c.Validate(updateUserRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Required fields are empty: %v", err))
	}

	// Get the enterprise
	enterpriseId := c.Get("enterpriseId").(string)

	user := models.User{}
	h.server.Repos.User.GetUserByUID(&user, uid, enterpriseId)
	if user.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "User not found")
	}

	var roles []models.Role
	h.server.Repos.Role.GetRolesByIds(&roles, enterpriseId, updateUserRequest.Roles)
	if len(updateUserRequest.Roles) != len(roles) {
		return responses.ErrorResponse(c, http.StatusBadRequest, "You do not have access to a role selected")
	}

	existUser := models.User{}
	h.server.Repos.User.GetUserByEmail(&existUser, updateUserRequest.Email)
	if existUser.ID != 0 && existUser.ID != user.ID {
		return responses.ErrorResponse(c, http.StatusBadRequest, "User already exists")
	}

	user.Name = updateUserRequest.Name
	user.Email = updateUserRequest.Email
	user.JobTitle = updateUserRequest.JobTitle
	user.Pronouns = updateUserRequest.Pronouns
	if err := h.service.Update(&user); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong updating the user in the database.")
	}

	currentUser := c.Get("currentUser").(*models.User)
	models.CreatePublicActivity(h.server.Db, c.RealIP(), enterpriseId, currentUser.ID, user.ID, "User", "User Updated", currentUser.SuperAdmin)

	h.server.Db.Model(user).Association("Roles").Clear().Append(roles)

	return responses.MessageResponse(c, http.StatusOK, "User successfully updated")
}

// List godoc
// @Summary Get users
// @Description Get the paginated list of users
// @ID users-index
// @Tags User Actions
// @Produce json
// @Param page query int false "The page number"
// @Param page_size query int false "The numbers of items to return. Max 100" minimum(1) maximum(100)"
// @Param search query string false "Search users name or email"
// @Success 200 {array} responses.UsersResponse
// @Failure 400 {object} responses.Error
// @Security ApiKeyAuth
// @Router /users [get]
func (h *UserHandler) List(c echo.Context) error {
	enterpriseId := c.Get("enterpriseId").(string)
	var scopes []func(db *gorm.DB) *gorm.DB

	// Handle search query
	search := c.QueryParam("search")
	if search != "" {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			searchString := fmt.Sprintf("%%%v%%", search)
			return db.Where("(name LIKE ? OR email LIKE ?)", searchString, searchString)
		})
	}

	var users []models.User
	page, pageSize := h.server.Repos.User.GetUsers(&users, enterpriseId, c, scopes)
	totalCount := h.server.Repos.User.Count(enterpriseId)

	response := responses.NewUsersResponse(users, totalCount, page, pageSize)
	return responses.Response(c, http.StatusOK, response)
}

// Get godoc
// @Summary Get user
// @Description Get specified users
// @ID users-get
// @Tags User Actions
// @Produce json
// @Param uid path string true "User UID"
// @Success 200 {array} responses.UserResponse
// @Failure 404 {object} responses.Error
// @Security ApiKeyAuth
// @Router /users/{uid} [get]
func (h *UserHandler) Get(c echo.Context) error {
	enterpriseId := c.Get("enterpriseId").(string)
	uid := c.Param("uid")

	user := &models.User{}
	h.server.Repos.User.GetUserByUID(user, uid, enterpriseId)
	if user.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "User not found")
	}

	response := responses.NewUserResponse(user)
	return responses.Response(c, http.StatusOK, response)
}

// Delete godoc
// @Summary Delete users
// @Description Delete User
// @ID users-delete
// @Tags User Actions
// @Accept json
// @Produce json
// @Param uid path string true "User UID"
// @Success 204 {object} responses.Data
// @Failure 404 {object} responses.Error
// @Security ApiKeyAuth
// @Router /users/{uid} [delete]
func (h *UserHandler) Delete(c echo.Context) error {
	enterpriseId := c.Get("enterpriseId").(string)
	uid := c.Param("uid")

	user := models.User{}
	h.server.Repos.User.GetUserByUID(&user, uid, enterpriseId)
	if user.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "User not found")
	}

	if err := h.service.Delete(&user); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong deleting the user from the database.")
	}

	currentUser := c.Get("currentUser").(*models.User)
	models.CreatePublicActivity(h.server.Db, c.RealIP(), enterpriseId, currentUser.ID, user.ID, "User", "User Deleted", currentUser.SuperAdmin)

	return responses.MessageResponse(c, http.StatusOK, "User successfully deleted")
}
