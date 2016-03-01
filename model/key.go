package model

import (
	"github.com/jinzhu/gorm"
)

type Key struct {
	gorm.Model

	Slug  string `json:"slug" sql:"unique_index"`
	Name  string `json:"name" sql:"unique_index"`
	Value string `json:"key" sql:"unique_index"`
}
