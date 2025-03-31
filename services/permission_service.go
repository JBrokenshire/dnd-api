package services

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"log"
)

// PermissionService helper for managing `permission`
type PermissionService struct {
	Db          *gorm.DB
	Permissions []models.Permission
}

func NewPermissionService(Db *gorm.DB) *PermissionService {
	p := &PermissionService{
		Db: Db,
	}
	// Pull a list of permissions form the database
	result := Db.Find(&p.Permissions)
	if result.Error != nil {
		log.Printf("There was an error fetching the permissions: %v", result.Error)
	}
	return p
}

// SyncPermissions Will see what permissions are in the database, and will add any that are missing from the hardcoded
// list above.
func (p *PermissionService) SyncPermissions() {
	for _, sp := range models.SystemPermission {
		// See if this permission exists. If not create it
		if p.permissionExists(sp.Subject, sp.Action) == false {
			p.createPermission(sp.Subject, sp.Action)
		}
	}
}

// permissionExists checks if the permission can be found in the local list
func (p *PermissionService) permissionExists(s models.Subject, a models.Action) bool {
	for _, dbPerm := range p.Permissions {
		if dbPerm.Subject == s && dbPerm.Action == a {
			// We have found the permission so continue
			return true
		}
	}
	return false
}

// createPermission adds the permission to the database and also adds it to the main "Admin" role.
func (p *PermissionService) createPermission(s models.Subject, a models.Action) {
	// We need to add the permission to the Admin role by default, so we can use the auto create associates to add the
	// permission and generate the associate at the same time.
	var adminRole models.Role
	p.Db.Where("id = 1").First(&adminRole)
	p.Db.Model(&adminRole).Association("Permissions").Append([]models.Permission{{
		Subject: s,
		Action:  a,
	}})
}
