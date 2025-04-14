package models

import "time"

type User struct {
	ID        uint       `gorm:"primary_key;auto_increment" json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	Admin     bool       `json:"admin"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" sql:"index"`
	LastSeen  *time.Time `json:"last_seen,omitempty"`
}
