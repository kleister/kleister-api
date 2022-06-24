package model

import (
	"time"
)

// ClientPack represents a client pack model definition.
type ClientPack struct {
	ClientID  string `storm:"id,index" gorm:"index:idx_id,unique;length:36"`
	Client    *Client
	PackID    string `storm:"id,index" gorm:"index:idx_id,unique;length:36"`
	Pack      *Pack
	CreatedAt time.Time
	UpdatedAt time.Time
}
