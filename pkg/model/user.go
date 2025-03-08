package model

import (
	"context"
	"strings"
	"time"

	"github.com/dchest/uniuri"
	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

var (
	_ bun.BeforeAppendModelHook = (*User)(nil)
)

// User defines the model for users table.
type User struct {
	bun.BaseModel `bun:"table:users"`

	ID        string       `bun:",pk,type:varchar(20)"`
	Scim      string       `bun:"type:varchar(255)"`
	Username  string       `bun:",unique,type:varchar(255)"`
	Password  string       `bun:"-"`
	Hashword  string       `bun:"type:varchar(255)"`
	Email     string       `bun:"type:varchar(255)"`
	Fullname  string       `bun:"type:varchar(255)"`
	Profile   string       `bun:"-"`
	Active    bool         `bun:"default:false"`
	Admin     bool         `bun:"default:false"`
	CreatedAt time.Time    `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time    `bun:",nullzero,notnull,default:current_timestamp"`
	Auths     []*UserAuth  `bun:"rel:has-many,join:id=user_id"`
	Groups    []*UserGroup `bun:"rel:has-many,join:id=user_id"`
	Mods      []*UserMod   `bun:"rel:has-many,join:id=user_id"`
	Packs     []*UserPack  `bun:"rel:has-many,join:id=user_id"`
}

// BeforeAppendModel implements the bun hook interface.
func (m *User) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		if m.ID == "" {
			m.ID = strings.ToLower(uniuri.NewLen(uniuri.UUIDLen))
		}

		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		if m.ID == "" {
			m.ID = strings.ToLower(uniuri.NewLen(uniuri.UUIDLen))
		}

		m.UpdatedAt = time.Now()
	}

	if m.Password != "" {
		bytes, err := bcrypt.GenerateFromPassword(
			[]byte(m.Password),
			10,
		)

		if err != nil {
			return err
		}

		m.Hashword = string(bytes)
	}

	return nil
}
