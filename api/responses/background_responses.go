package responses

import m "dnd-api/db/models"

type BackgroundResponse struct {
	Name        string `json:"name"`
	Feature     string `json:"feature"`
	Description string `json:"description"`
}

func NewBackgroundResponse(bg *m.Background) *BackgroundResponse {
	return &BackgroundResponse{
		Name:        bg.Name,
		Feature:     bg.Feature,
		Description: bg.Description,
	}
}
