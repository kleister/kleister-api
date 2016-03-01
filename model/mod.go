package model

import (
	"github.com/jinzhu/gorm"
)

type Mod struct {
	gorm.Model

	Slug        string     `json:"slug" sql:"unique_index"`
	Name        string     `json:"name" sql:"unique_index"`
	Description string     `json:"description" sql:"type:text"`
	Author      string     `json:"author"`
	Website     string     `json:"website"`
	Donate      string     `json:"donate"`
	Versions    []*Version `json:"versions"`
}
