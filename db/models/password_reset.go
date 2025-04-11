package models

import "time"

type PasswordReset struct {
	ID         uint       `json:"-"`
	CreatedAt  time.Time  `json:"-"`
	UpdatedAt  time.Time  `json:"-"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
	Token      string     `json:"token"`
	Used       bool       `json:"used"`
	UserUid    string     `json:"user_uid"`
	ValidUntil time.Time  `json:"valid_until"`
	
	User User `json:"omitempty"`
}
