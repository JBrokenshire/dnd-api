package handlers

import (
	"dnd-api/api"
	"dnd-api/api/requests"
	"dnd-api/api/responses"
	"dnd-api/db/models"
	"dnd-api/pkg/closer"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/random"
	"golang.org/x/exp/slices"
	"log"
	"math"
	"net/http"
	"path/filepath"
)

type CharacterHandler struct {
	server *api.Server
}

func NewCharacterHandler(server *api.Server) *CharacterHandler {
	return &CharacterHandler{
		server: server,
	}
}

// List godoc
// @Summary List characters
// @Description List characters (paginated)
// @ID characters-list
// @Tags Character Actions
// @Accept json
// @Produce json
// @Param search query string false "Search characters by name"
// @Param page query int false "The page number"
// @Param page_size query int false "The numbers of items to return. Max 100"
// @Success 200 {object} responses.CharacterPaginatedResponse
// @Router /characters [get]
func (h *CharacterHandler) List(c echo.Context) error {
	currentUser := c.Get("currentUser").(*models.User)

	var scopes []func(db *gorm.DB) *gorm.DB

	search := c.QueryParam("search")
	if search != "" {
		searchTerm := fmt.Sprintf("%%%s%%", search)
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("name LIKE ?", searchTerm)
		})
	}

	characters, page, pageSize := h.server.Repos.Character.GetCharacters(c, currentUser.ID, scopes)
	count := h.server.Repos.Character.CountCharacters(currentUser.ID, scopes)

	res := responses.NewCharacterPaginatedResponse(characters, count, page, pageSize)
	return responses.Response(c, http.StatusOK, res)
}

// Get godoc
// @Summary Get character by ID
// @Description Get character by ID
// @ID characters-get
// @Tags Character Actions
// @Accept json
// @Produce json
// @Param id path int true "Character ID"
// @Success 200 {object} responses.CharacterResponse
// @Router /characters/{id} [get]
func (h *CharacterHandler) Get(c echo.Context) error {
	currentUser := c.Get("currentUser").(*models.User)
	id := c.Param("id")

	character := h.server.Repos.Character.GetById(id, currentUser.ID)
	if character.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Character not found")
	}

	res := responses.NewCharacterResponse(character)
	return responses.Response(c, http.StatusOK, res)
}

// Create godoc
// @Summary Create character
// @Description Create character
// @ID characters-create
// @Tags Character Actions
// @Accept json
// @Produce json
// @Param params body requests.CreateCharacterRequest true "Character information"
// @Success 201 {object} responses.CharacterResponse
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /characters [post]
func (h *CharacterHandler) Create(c echo.Context) error {
	currentUser := c.Get("currentUser").(*models.User)

	request := new(requests.CreateCharacterRequest)
	if err := c.Bind(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid: "+err.Error())
	}

	class := h.server.Repos.Class.GetById(request.ClassId)
	if class.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Class not found")
	}

	race := h.server.Repos.Race.GetById(request.RaceId)
	if race.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Race not found")
	}

	if !slices.Contains(models.ValidAdvancementTypes, request.AdvancementType) {
		return responses.ErrorResponse(c, http.StatusBadRequest, "AdvancementType is not valid")
	}

	if !slices.Contains(models.ValidHitPointTypes, request.HitPointType) {
		return responses.ErrorResponse(c, http.StatusBadRequest, "HitPointType is not valid")
	}

	character := &models.Character{
		UserId:   currentUser.ID,
		Name:     request.Name,
		ClassId:  class.ID,
		RaceId:   race.ID,
		Pronouns: request.Pronouns,
		Level:    request.Level,

		Strength:     request.Strength,
		Dexterity:    request.Dexterity,
		Constitution: request.Constitution,
		Intelligence: request.Intelligence,
		Wisdom:       request.Wisdom,
		Charisma:     request.Charisma,

		AdvancementType: request.AdvancementType,
		HitPointType:    request.HitPointType,
	}

	err := h.server.Repos.Character.Create(character)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong creating the character")
	}

	character.Class = *class
	character.Race = *race

	res := responses.NewCharacterResponse(character)
	return responses.Response(c, http.StatusCreated, res)
}

// Update godoc
// @Summary Update character
// @Description Update character
// @ID characters-update
// @Tags Character Actions
// @Accept json
// @Produce json
// @Param id path int true "Character ID"
// @Param params body requests.UpdateCharacterRequest true "Character information"
// @Success 200 {object} responses.CharacterResponse
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /characters/{id} [put]
func (h *CharacterHandler) Update(c echo.Context) error {
	currentUser := c.Get("currentUser").(*models.User)
	id := c.Param("id")

	request := new(requests.UpdateCharacterRequest)
	if err := c.Bind(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid: "+err.Error())
	}

	character := h.server.Repos.Character.GetById(id, currentUser.ID)
	if character.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Character not found")
	}

	class := h.server.Repos.Class.GetById(request.ClassId)
	if class.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Class not found")
	}

	race := h.server.Repos.Race.GetById(request.RaceId)
	if race.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Race not found")
	}

	if !slices.Contains(models.ValidAdvancementTypes, request.AdvancementType) {
		return responses.ErrorResponse(c, http.StatusBadRequest, "AdvancementType is not valid")
	}

	if !slices.Contains(models.ValidHitPointTypes, request.HitPointType) {
		return responses.ErrorResponse(c, http.StatusBadRequest, "HitPointType is not valid")
	}

	character.Name = request.Name
	character.ClassId = request.ClassId
	character.RaceId = request.RaceId
	character.Pronouns = request.Pronouns
	character.Level = request.Level

	character.Strength = request.Strength
	character.Dexterity = request.Dexterity
	character.Constitution = request.Constitution
	character.Intelligence = request.Intelligence
	character.Wisdom = request.Wisdom
	character.Charisma = request.Charisma

	character.AdvancementType = request.AdvancementType
	character.HitPointType = request.HitPointType

	character.Class = *class
	character.Race = *race

	err := h.server.Repos.Character.Update(character)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong updating the character")
	}

	res := responses.NewCharacterResponse(character)
	return responses.Response(c, http.StatusOK, res)
}

// Delete godoc
// @Summary Delete character
// @Description Delete character
// @ID characters-delete
// @Tags Character Actions
// @Accept json
// @Produce json
// @Param id path int true "Character ID"
// @Success 200 {object} responses.Data
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /characters/{id} [delete]
func (h *CharacterHandler) Delete(c echo.Context) error {
	currentUser := c.Get("currentUser").(*models.User)
	id := c.Param("id")

	character := h.server.Repos.Character.GetById(id, currentUser.ID)
	if character.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Character not found")
	}

	// Delete profile picture
	if character.ProfilePicture.ID != 0 {
		// Delete from file store
		err := h.server.Dependencies.GetFileStore().Delete(fmt.Sprintf("%v/%v", character.ProfilePicture.FileLocation, character.ProfilePicture.Filename))
		if err != nil {
			log.Println("Error deleting profile picture from the file store: ", err.Error())
			return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong deleting the profile picture from the file store")
		}

		// Delete from DB
		err = h.server.Repos.File.Delete(character.ProfilePicture)
		if err != nil {
			log.Println("Error deleting profile picture from the database: ", err.Error())
			return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong deleting the profile picture from the database")
		}
	}

	err := h.server.Repos.Character.Delete(character)
	if err != nil {
		log.Println("Error deleting character from the database: ", err.Error())
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong deleting the character")
	}

	return responses.MessageResponse(c, http.StatusOK, "Character deleted successfully")
}

// UploadProfilePicture godoc
// @Summary Upload character profile picture
// @Description Upload character profile picture
// @ID characters-upload-profile-picture
// @Tags Character File Actions
// @Accept json
// @Produce json
// @Param id path string true "Character ID"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /characters/{id}/upload/profile-picture [post]
func (h *CharacterHandler) UploadProfilePicture(c echo.Context) error {
	id := c.Param("id")
	currentUser := c.Get("currentUser").(*models.User)

	character := h.server.Repos.Character.GetById(id, currentUser.ID)
	if character.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Character not found")
	}

	// Check the file mimetype - We only want to accept images
	fileName, fileMimeType, err := h.server.Dependencies.GetFileService().ReadFileInfo(c)
	if err != nil {
		log.Printf("Error reading file info: %v", err)
		return responses.ErrorResponse(c, http.StatusBadRequest, "Unable to read file")
	}
	fileExtension := filepath.Ext(fileName)
	allowedFileExtensions := []string{"jpg", "jpeg", "png", "webp"}
	if !h.server.Dependencies.GetFileService().FileExtensionAllowed(fileExtension, allowedFileExtensions) {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Invalid file type")
	}

	allowedMIMEs := []string{"image/jpeg", "image/png", "image/webp"}
	if !h.server.Dependencies.GetFileService().MIMETypeAllowed(fileMimeType, allowedMIMEs) {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Invalid file content type")
	}

	// Get file from request
	file, err := c.FormFile("file")
	if err != nil {
		log.Printf("Error getting file from form: %v", err)
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Error getting file from request")
	}

	// Change the file name
	newFilename := random.String(32)
	if fileExtension != "" {
		newFilename += fileExtension
	}
	file.Filename = newFilename

	// Upload file
	src, err := file.Open()
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Error opening file")
	}
	defer closer.Close(src)

	path := fmt.Sprintf("characters/%v", character.ID)
	err = h.server.Dependencies.GetFileStore().Save(path, src, file.Filename)
	if err != nil {
		log.Printf("Error saving file: %v", err)
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Error saving file")
	}

	// Create DB record
	profilePictureFile := &models.File{
		Model:        models.FileModelCharacterProfilePicture,
		ModelId:      character.ID,
		Filename:     newFilename,
		FileLocation: path,
	}
	if err := h.server.Repos.File.Create(profilePictureFile); err != nil {
		log.Printf("Error creating file record: %v", err)
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Error creating file record")
	}

	return responses.MessageResponse(c, http.StatusOK, "File uploaded")
}

// ToggleInspiration godoc
// @Summary Toggle character inspiration
// @Description Toggle character inspiration
// @ID characters-toggle-inspiration
// @Tags Character Actions
// @Param id path string true "Character ID"
// @Success 200 {object} responses.CharacterResponse
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /characters/{id}/inspiration [post]
func (h *CharacterHandler) ToggleInspiration(c echo.Context) error {
	id := c.Param("id")
	currentUser := c.Get("currentUser").(*models.User)

	character := h.server.Repos.Character.GetById(id, currentUser.ID)
	if character.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Character not found")
	}

	character.Inspiration = !character.Inspiration
	err := h.server.Repos.Character.Update(character)
	if err != nil {
		log.Printf("Error toggling character inspiration: %v", err)
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Error toggling character inspiration")
	}

	res := responses.NewCharacterResponse(character)
	return responses.Response(c, http.StatusOK, res)
}

// UpdateHealth godoc
// @Summary Update character health
// @Description Update character health
// @ID characters-update-health
// @Tags Character Actions
// @Accept json
// @Produce json
// @Param id path int true "Character ID"
// @Param params body requests.UpdateCharacterHealthRequest true "Health information"
// @Success 200 {object} responses.CharacterResponse
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /characters/{id}/health [put]
func (h *CharacterHandler) UpdateHealth(c echo.Context) error {
	id := c.Param("id")
	currentUser := c.Get("currentUser").(*models.User)

	request := new(requests.UpdateCharacterHealthRequest)
	if err := c.Bind(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid: "+err.Error())
	}

	character := h.server.Repos.Character.GetById(id, currentUser.ID)
	if character.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Character not found")
	}

	character.CurrentHitPoints = uint(math.Max(math.Min(float64(request.CurrentHitPoints), float64(character.MaxHitPoints)), 0))
	err := h.server.Repos.Character.Update(character)
	if err != nil {
		log.Printf("Error updating character health: %v", err)
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Error updating character health")
	}

	res := responses.NewCharacterResponse(character)
	return responses.Response(c, http.StatusOK, res)
}
