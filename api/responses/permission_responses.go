package responses

import (
	m "dnd-api/db/models"
)

type PermissionResponse struct {
	ID      uint   `json:"id" example:"1"`
	Subject string `json:"subject" example:"Role"`
	Action  string `json:"action" example:"Read"`
}

func NewPermissionResponse(permissions []m.Permission) []PermissionResponse {
	permissionResponses := make([]PermissionResponse, 0)

	for i := range permissions {
		permissionResponses = append(permissionResponses, PermissionResponse{
			ID:      permissions[i].ID,
			Subject: string(permissions[i].Subject),
			Action:  string(permissions[i].Action),
		})
	}

	return permissionResponses
}
