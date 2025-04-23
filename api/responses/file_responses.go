package responses

import (
	"dnd-api/db/models"
	"time"
)

type FileResponse struct {
	ID            uint      `json:"id" example:"1"`
	EnterpriseUid string    `json:"enterprise_uid"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Model         string    `json:"model"`
	ModelId       uint      `json:"model_id"`
	Filename      string    `json:"filename"`
	FileLocation  string    `json:"file_location"`
}

func NewFileResponse(file *models.File) *FileResponse {
	return &FileResponse{
		ID:           file.ID,
		CreatedAt:    file.CreatedAt,
		UpdatedAt:    file.UpdatedAt,
		Model:        file.Model,
		ModelId:      file.ModelId,
		Filename:     file.Filename,
		FileLocation: file.FileLocation,
	}
}
