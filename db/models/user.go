package models

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	ID         uint       `json:"-" gorm:"primary_key"`
	Uid        string     `json:"uid"`
	CreatedAt  time.Time  `json:"-"`
	UpdatedAt  time.Time  `json:"-"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" sql:"index"`
	LastSeen   *time.Time `json:"last_seen,omitempty"`
	Email      string     `json:"email" gorm:"type:varchar(200);"`
	Name       string     `json:"name" gorm:"type:varchar(200);"`
	Pronouns   string     `json:"pronouns"`
	Password   string     `json:"password,omitempty" gorm:"type:varchar(200);"`
	SuperAdmin bool       `json:"-"`

	Roles []Role `json:"roles" gorm:"many2many:user_roles;"`
}

// BeforeCreate will set Action UUID rather than numeric ID.
func (u *User) BeforeCreate(scope *gorm.Scope) error {
	// If the UID is already set then just return.
	if u.Uid != "" {
		return nil
	}
	// Generate a new uid and set
	newUuid, _ := uuid.NewV4()
	return scope.SetColumn("Uid", newUuid.String())
}

func (u User) MarshalJSON() ([]byte, error) {
	type user User // prevent recursion
	x := user(u)
	x.Password = ""
	return json.Marshal(x)
}

// HasPermission checks if the current user has permission
func (u *User) HasPermission(subject Subject, action Action) bool {
	// Loop through the users roles and permissions to see if they have the permission specified
	for _, p := range u.GetPermissionList() {
		if p.Action == action && p.Subject == subject {
			return true
		}
	}
	return false
}

// GetPermissionList returns a unique list of permissions form the user, based form the roles they have
func (u *User) GetPermissionList() []Permission {
	uniquePermissionsMap := make(map[uint]Permission)
	var uniquePermissions []Permission
	for _, r := range u.Roles {
		for _, p := range r.Permissions {
			uniquePermissionsMap[p.ID] = p
		}
	}
	for _, value := range uniquePermissionsMap {
		uniquePermissions = append(uniquePermissions, value)
	}

	// Everyone should have the default permissions
	uniquePermissions = append(uniquePermissions,
		Permission{
			Subject: "Default",
			Action:  "Read",
		}, Permission{
			Subject: "Auth",
			Action:  "Read",
		},
	)

	// If they are a super admin, pass in a manage all field
	if u.SuperAdmin == true {
		uniquePermissions = append(uniquePermissions, Permission{
			Subject: "All",
			Action:  "Manage",
		})
	}

	return uniquePermissions
}
