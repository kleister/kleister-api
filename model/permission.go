package model

import (
	"time"
)

type Permission struct {
	ID            int64     `json:"id" gorm:"primary_key"`
	User          *User     `json:"-"`
	UserID        int       `json:"user_id" sql:"index"`
	Admin         bool      `json:"admin" sql:"default:false"`
	CreateUsers   bool      `json:"create_users" sql:"default:false"`
	DeleteUsers   bool      `json:"delete_users" sql:"default:false"`
	ManageUsers   bool      `json:"manage_users" sql:"default:false"`
	CreateKeys    bool      `json:"create_keys" sql:"default:false"`
	DeleteKeys    bool      `json:"delete_keys" sql:"default:false"`
	ManageKeys    bool      `json:"manage_keys" sql:"default:false"`
	CreateClients bool      `json:"create_clients" sql:"default:false"`
	DeleteClients bool      `json:"delete_clients" sql:"default:false"`
	ManageClients bool      `json:"manage_clients" sql:"default:false"`
	CreatePacks   bool      `json:"create_packs" sql:"default:false"`
	DeletePacks   bool      `json:"delete_packs" sql:"default:false"`
	ManagePacks   bool      `json:"manage_packs" sql:"default:false"`
	CreateMods    bool      `json:"create_mods" sql:"default:false"`
	DeleteMods    bool      `json:"delete_mods" sql:"default:false"`
	ManageMods    bool      `json:"manage_mods" sql:"default:false"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Packs         []*Pack   `json:"packs" gorm:"many2many:permission_packs;"`
}
