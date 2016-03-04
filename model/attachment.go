package model

import (
	"time"
)

type Attachment struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	OwnerId   int       `json:"-"`
	OwnerType string    `json:"-"`
	URL       string    `json:"url"`
	MD5       string    `json:"md5"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
