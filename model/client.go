package model

import (
	"github.com/jinzhu/gorm"
)

type Client struct {
	gorm.Model

	Slug  string `json:"slug" sql:"unique_index"`
	Name  string `json:"name" sql:"unique_index"`
	Value string `json:"uuid" sql:"unique_index"`
}
