package model

import (
	"time"
)

// Users is simply a collection of user structs.
type Users []*User

// User represents a user model definition.
type User struct {
	ID         int64       `json:"id" gorm:"primary_key"`
	Permission *Permission `json:"permission"`
	Slug       string      `json:"slug" sql:"unique_index"`
	Username   string      `json:"username" sql:"unique_index"`
	Email      string      `json:"email"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	Password   string      `json:"-"`
}
