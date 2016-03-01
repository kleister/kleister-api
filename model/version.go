package model

import (
	"github.com/jinzhu/gorm"
)

type Version struct {
	gorm.Model

	Mod   *Mod   `json:"mod"`
	ModID int    `json:"mod_id" sql:"index"`
	Slug  string `json:"slug" sql:"unique_index"`
	Name  string `json:"name" sql:"unique_index"`
	MD5   string `json:"md5"`
}
