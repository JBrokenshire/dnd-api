package handlers

import (
	"dnd-api/api"
	"github.com/labstack/echo/v4"
)

type FileHandler struct {
	server *api.Server
}

func NewFileHandler(server *api.Server) *FileHandler {
	return &FileHandler{server: server}
}

// Get godoc
// @Summary Get file
// @Description Get file with authenticated routes
// @ID files-get
// @Tags Files Actions
// @Produce text/plain
// @Param filepath path string true "filepath"
// @Success 200 {file} runtime.File
// @Failure 404 {object} responses.Error
// @Security ApiKeyAuth
// @Router /files/{filepath} [get]
func (h *FileHandler) Get(c echo.Context) error {
	path := c.Param("filepath")
	return h.server.Dependencies.GetFileStore().Get(c, path)
}
