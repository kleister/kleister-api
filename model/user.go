package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	Permission *Permission `json:"permission"`
	Slug       string      `json:"slug" sql:"unique_index"`
	Username   string      `json:"username" sql:"unique_index"`
	Email      string      `json:"email"`
	Password   string      `json:"-"`
}
