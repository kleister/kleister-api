package model

import (
	"context"
	"strings"
	"time"

	"github.com/dchest/uniuri"
	"github.com/uptrace/bun"
)

var (
	_ bun.BeforeAppendModelHook = (*UserToken)(nil)
)

// UserTokenKind is the custom type for kind of tokens.
type UserTokenKind string

const (
	// UserTokenKindRedirect defines the kind used for redirects.
	UserTokenKindRedirect UserTokenKind = "redirect"
)

// UserToken defines the model for user_tokens table.
type UserToken struct {
	bun.BaseModel `bun:"table:user_tokens"`

	ID        string        `bun:",pk,type:varchar(20)"`
	UserID    string        `bun:"type:varchar(20)"`
	User      *User         `bun:"rel:belongs-to,join:user_id=id"`
	Kind      UserTokenKind `bun:"type:varchar(128)"`
	Token     string        `bun:",unique,type:varchar(128)"`
	CreatedAt time.Time     `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time     `bun:",nullzero,notnull,default:current_timestamp"`
}

// BeforeAppendModel implements the bun hook interface.
func (m *UserToken) BeforeAppendModel(_ context.Context, query bun.Query) error {
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

	return nil
}
