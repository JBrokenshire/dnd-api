package models

import (
	"dnd-api/pkg/rand"
	"github.com/jinzhu/gorm"
	"time"
)

type ApiKey struct {
	ID           uint        `json:"-" gorm:"primary_key"`
	Uid          string      `json:"uid"`
	IpRanges     string      `json:"ip_ranges"`
	Label        string      `json:"label"`
	Enabled      bool        `json:"enabled"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	DeletedAt    *time.Time  `json:"deleted_at"`
	PolicyId     uint        `json:"policy_id"`
	Policy       *EapiPolicy `json:"-" gorm:"foreignKey:PolicyId;references:PolicyId"`
	AppBlacklist string      `json:"app_blacklist"` // csv list of app packages to block
}

// BeforeCreate will set UUID rather than numeric ID.
func (g *ApiKey) BeforeCreate(scope *gorm.Scope) error {
	// If the UID is already set then just return.
	if g.Uid != "" {
		return nil
	}
	// Generate a new uid and set
	uuidShort := rand.ApiKey()
	return scope.SetColumn("Uid", uuidShort)
}
