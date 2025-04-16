package handlers

import (
	"dnd-api/api"
	"dnd-api/api/requests"
	"dnd-api/api/responses"
	"dnd-api/db/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RaceHandler struct {
	server *api.Server
}

func NewRaceHandler(server *api.Server) *RaceHandler {
	return &RaceHandler{
		server: server,
	}
}

// List godoc
// @Summary List races
// @Description List races (paginated)
// @ID races-list
// @Tags Race Actions
// @Accept json
// @Produce json
// @Param search query string false "Search races by name"
// @Param page query int false "The page number"
// @Param page_size query int false "The numbers of items to return. Max 100"
// @Success 200 {object} responses.RacePaginatedResponse
// @Router /races [get]
func (h *RaceHandler) List(c echo.Context) error {
	var scopes []func(db *gorm.DB) *gorm.DB

	search := c.QueryParam("search")
	if search != "" {
		searchTerm := fmt.Sprintf("%%%s%%", search)
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("name LIKE ?", searchTerm)
		})
	}

	races, page, pageSize := h.server.Repos.Race.GetRaces(c, scopes)
	count := h.server.Repos.Race.CountRaces(scopes)

	res := responses.NewRacePaginatedResponse(races, count, page, pageSize)
	return responses.Response(c, http.StatusOK, res)
}

// Get godoc
// @Summary Get race by ID
// @Description Get race by ID
// @ID races-get
// @Tags Race Actions
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} responses.RaceResponse
// @Failure 404 {object} responses.Error
// @Router /races/{id} [get]
func (h *RaceHandler) Get(c echo.Context) error {
	id := c.Param("id")

	race := h.server.Repos.Race.GetById(id)
	if race.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Race not found")
	}

	res := responses.NewRaceResponse(race)
	return responses.Response(c, http.StatusOK, res)
}

// Create godoc
// @Summary Create race
// @Description Create race
// @ID races-create
// @Tags Race Actions
// @Accept json
// @Produce json
// @Param params body requests.CreateRaceRequest true "Race information"
// @Success 201 {object} responses.RaceResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /races [post]
func (h *RaceHandler) Create(c echo.Context) error {
	request := new(requests.CreateRaceRequest)
	if err := c.Bind(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid: "+err.Error())
	}

	race := &models.Race{
		Name: request.Name,
	}
	err := h.server.Repos.Race.Create(race)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong creating the race")
	}

	res := responses.NewRaceResponse(race)
	return responses.Response(c, http.StatusCreated, res)
}

// Update godoc
// @Summary Update race
// @Description Update race
// @ID races-update
// @Tags Race Actions
// @Accept json
// @Produce json
// @Param id path string true "Race ID"
// @Param params body requests.UpdateRaceRequest true "Race information"
// @Success 200 {object} responses.RaceResponse
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /races/{id} [put]
func (h *RaceHandler) Update(c echo.Context) error {
	id := c.Param("id")

	request := new(requests.UpdateRaceRequest)
	if err := c.Bind(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid: "+err.Error())
	}

	race := h.server.Repos.Race.GetById(id)
	if race.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Race not found")
	}

	race.Name = request.Name

	err := h.server.Repos.Race.Update(race)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong updating the race")
	}

	res := responses.NewRaceResponse(race)
	return responses.Response(c, http.StatusOK, res)
}

// Delete godoc
// @Summary Delete race
// @Description Delete race
// @ID races-delete
// @Tags Race Actions
// @Accept json
// @Produce json
// @Param id path string true "Race ID"
// @Success 200 {object} responses.Data
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /races/{id} [delete]
func (h *RaceHandler) Delete(c echo.Context) error {
	id := c.Param("id")

	race := h.server.Repos.Race.GetById(id)
	if race.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Race not found")
	}

	err := h.server.Repos.Race.Delete(race)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong deleting the race")
	}

	return responses.MessageResponse(c, http.StatusOK, "Race deleted successfully")
}
