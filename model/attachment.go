package model

import (
	"time"
)

// Attachment represents any uploadable asset.
type Attachment struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	OwnerID   uint      `json:"-"`
	OwnerType string    `json:"-"`
	URL       string    `json:"url"`
	MD5       string    `json:"md5"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
