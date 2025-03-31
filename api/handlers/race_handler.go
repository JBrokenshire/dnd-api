package handlers

import (
	"dnd-api/api"
	"dnd-api/api/requests"
	"dnd-api/api/responses"
	m "dnd-api/db/models"
	"fmt"
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
// @Description Get a paginated list of races
// @ID races-list
// @Tags Race
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Number of items to return"
// @Success 200 {object} responses.RacePaginatedResponse
// @Router /races [get]
func (h *RaceHandler) List(c echo.Context) error {
	races, page, pageSize := h.server.Repos.Race.GetRaces(c)
	count := h.server.Repos.Race.Count()

	res := responses.NewRacePaginatedResponse(races, count, page, pageSize)
	return responses.Response(c, http.StatusOK, res)
}

// Get godoc
// @Summary Get race
// @Description Get a race by its ID
// @ID races-get
// @Tags Race
// @Accept json
// @Produce json
// @Param id path int true "Race ID"
// @Success 200 {object} responses.RaceResponse
// @Failure 404 {object} responses.Error
// @Router /races/{id} [get]
func (h *RaceHandler) Get(c echo.Context) error {
	id := c.Param("id")

	race := h.server.Repos.Race.GetRace(id)
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
// @Tags Race
// @Accept json
// @Produce json
// @Param params body requests.CreateRaceRequest true "Race information"
// @Success 200 {object} responses.RaceResponse
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /races [post]
func (h *RaceHandler) Create(c echo.Context) error {
	request := new(requests.CreateRaceRequest)
	if err := c.Bind(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Required fields are empty or not valid: %v", err.Error()))
	}

	race := &m.Race{
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
// @Tags Race
// @Accept json
// @Produce json
// @Param id path int true "Race ID"
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
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Required fields are empty or not valid: %v", err.Error()))
	}

	race := h.server.Repos.Race.GetRace(id)
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
// @Tags Race
// @Accept json
// @Produce json
// @Param id path int true "Race ID"
// @Success 200 {object} responses.Data
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /races/{id} [delete]
func (h *RaceHandler) Delete(c echo.Context) error {
	id := c.Param("id")

	race := h.server.Repos.Race.GetRace(id)
	if race.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Race not found")
	}

	err := h.server.Repos.Race.Delete(race)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong deleting the race")
	}

	return responses.MessageResponse(c, http.StatusOK, "Race deleted successfully")
}
