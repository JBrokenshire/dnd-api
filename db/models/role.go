package models

import "time"

type Role struct {
	ID        uint       `gorm:"primary_key" diff:"slice"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" sql:"index"`
	Name      string     `json:"name" diff:"slice"`
	CreatedBy uint       `json:"created_by,omitempty"`
	DeletedBy uint       `json:"deleted_by,omitempty"`

	Permissions []Permission `json:"Permissions,omitempty" gorm:"many2many:role_permissions;"`
}
