package model

import (
	"github.com/jinzhu/gorm"
)

type Attachment struct {
	gorm.Model

	OwnerId   int    `json:"-"`
	OwnerType string `json:"-"`
	URL       string `json:"url"`
	MD5       string `json:"md5"`
}
