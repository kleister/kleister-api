package model

import (
	"github.com/jinzhu/gorm"
)

type Pack struct {
	gorm.Model

	Icon          *Attachment `json:"icon" gorm:"polymorphic:Owner"`
	Logo          *Attachment `json:"logo" gorm:"polymorphic:Owner"`
	Background    *Attachment `json:"background" gorm:"polymorphic:Owner"`
	Recommended   *Build      `json:"recommended"`
	RecommendedID int         `json:"recommended_id" sql:"index"`
	Latest        *Build      `json:"latest"`
	LatestID      int         `json:"latest_id" sql:"index"`
	Slug          string      `json:"slug" sql:"unique_index"`
	Name          string      `json:"name" sql:"unique_index"`
	Website       string      `json:"website"`
	Hidden        bool        `json:"hidden" sql:"default:true"`
	Private       bool        `json:"private" sql:"default:false"`
	Builds        []*Build    `json:"builds"`
}
