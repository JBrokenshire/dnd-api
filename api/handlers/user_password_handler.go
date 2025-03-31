package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"purplevisits.com/mdm/config"
	"purplevisits.com/mdm/db/models"
	s "purplevisits.com/mdm/api"
	"purplevisits.com/mdm/api/requests"
	"purplevisits.com/mdm/api/responses"
	"purplevisits.com/mdm/services"
)

type UserPasswordHandler struct {
	server *s.Server
}

func NewUserPasswordHandler(server *s.Server) *UserPasswordHandler {
	return &UserPasswordHandler{server: server}
}

// Update godoc
// @Summary Update users passwords
// @Description Update User passwords
// @ID users-password-update
// @Tags User Actions
// @Accept json
// @Produce json
// @Param uid path string true "User UID"
// @Param params body requests.PasswordUpdateRequest true "Old password and new password"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Security ApiKeyAuth
// @Router /users/{uid}/password [put]
func (h *UserPasswordHandler) Update(c echo.Context) error {
	updatePasswordRequest := new(requests.PasswordUpdateRequest)
	uid := c.Param("uid")

	if err := c.Bind(updatePasswordRequest); err != nil {
		return err
	}

	if err := c.Validate(updatePasswordRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	// Get the enterprise
	enterpriseId := c.Get("enterpriseId").(string)

	user := models.User{}
	h.server.Repos.User.GetUserByUID(&user, uid, enterpriseId)

	if user.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "User not found")
	}

	if updatePasswordRequest.NewPasswordConfirm != updatePasswordRequest.NewPassword {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Confirmation password does not match")
	}

	enterprise := h.server.Repos.Enterprise.GetEnterprise(enterpriseId)
	if enterprise.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Cannot find the enterprise associated with this user")
	}
	isSecure, errMsg := h.server.Dependencies.PasswordSecurityService().IsPasswordSecure(updatePasswordRequest.NewPassword, enterprise.UserPasswordListCheck, enterprise.UserStrongPasswordRequired)
	if !isSecure {
		return responses.ErrorResponse(c, http.StatusBadRequest, errMsg)
	}

	userService := services.NewUserService(h.server.Db, config.Get().HashCost)
	if err := userService.UpdatePassword(&user, updatePasswordRequest.NewPassword); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong updating the user's password in the database.")
	}

	currentUser := c.Get("currentUser").(*models.User)
	models.CreatePublicActivity(h.server.Db, c.RealIP(), enterpriseId, currentUser.ID, user.ID, "User", "User Password Updated", currentUser.SuperAdmin)

	return responses.MessageResponse(c, http.StatusOK, "User password updated")
}
