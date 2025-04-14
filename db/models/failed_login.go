package models

import (
	"time"
)

type FailedLogin struct {
	ID        int       `gorm:"primary_key;auto_increment" json:"id"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
