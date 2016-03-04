package model

import (
	"time"

	"github.com/satori/go.uuid"
)

type Forges []*Forge

type Forge struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	Slug      string    `json:"slug" sql:"unique_index"`
	Name      string    `json:"version" sql:"unique_index"`
	Minecraft string    `json:"minecraft"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Builds    []*Build  `json:"-"`
}

func (u *Forge) BeforeSave() (err error) {
	if u.Slug == "" {
		u.Slug = uuid.NewV4().String()
	}

	return nil
}
