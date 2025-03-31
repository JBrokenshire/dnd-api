package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// EapiPolicy creates a policy which can be attached to an API key.
type EapiPolicy struct {
	ID          uint       `json:"id" gorm:"primary_key"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	OsconfigId  uint       `json:"osconfig_id"`

	Policy *MposPolicy `json:"policy"`
}

type MposPolicy struct {
	S3link   string `json:"s3_link"`
	Filename string `json:"filename"`
	Onboard  bool   `json:"onboard"`
}

func (mp *MposPolicy) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), &mp)
}

func (mp MposPolicy) Value() (driver.Value, error) {
	val, err := json.Marshal(mp)
	return string(val), err
}
