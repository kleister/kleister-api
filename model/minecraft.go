package model

import (
	"github.com/jinzhu/gorm"
)

type Minecraft struct {
	gorm.Model

	Slug   string   `json:"slug" sql:"unique_index"`
	Name   string   `json:"version" sql:"unique_index"`
	MD5    string   `json:"md5"`
	Builds []*Build `json:"builds"`
}
