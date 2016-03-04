package model

import (
	"time"
)

type Forge struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	Slug      string    `json:"slug" sql:"unique_index"`
	Name      string    `json:"version" sql:"unique_index"`
	Minecraft string    `json:"minecraft"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Builds    []*Build  `json:"builds"`
}

// http://files.minecraftforge.net/maven/net/minecraftforge/forge/json
