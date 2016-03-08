package model

import (
	"time"
)

// Permissions is simply a collection of permission structs.
type Permissions []*Permission

// Permission represents a permission model definition.
type Permission struct {
	ID             uint      `json:"-" gorm:"primary_key"`
	User           *User     `json:"-"`
	UserID         uint      `json:"-" sql:"index"`
	DisplayUsers   bool      `json:"display_users" sql:"default:false"`
	ChangeUsers    bool      `json:"change_users" sql:"default:false"`
	DeleteUsers    bool      `json:"delete_users" sql:"default:false"`
	DisplayKeys    bool      `json:"display_keys" sql:"default:false"`
	ChangeKeys     bool      `json:"change_keys" sql:"default:false"`
	DeleteKeys     bool      `json:"delete_keys" sql:"default:false"`
	DisplayClients bool      `json:"display_clients" sql:"default:false"`
	ChangeClients  bool      `json:"change_clients" sql:"default:false"`
	DeleteClients  bool      `json:"delete_clients" sql:"default:false"`
	DisplayPacks   bool      `json:"display_packs" sql:"default:false"`
	ChangePacks    bool      `json:"change_packs" sql:"default:false"`
	DeletePacks    bool      `json:"delete_packs" sql:"default:false"`
	DisplayMods    bool      `json:"display_mods" sql:"default:false"`
	ChangeMods     bool      `json:"change_mods" sql:"default:false"`
	DeleteMods     bool      `json:"delete_mods" sql:"default:false"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
	Packs          *Packs    `json:"packs" gorm:"many2many:permission_packs;"`
}

// Defaults prefills the struct with some default values.
func (u *Permission) Defaults() {
	// Currently no default values required.
}
