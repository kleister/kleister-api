package token

import (
	"errors"
	"time"

	"github.com/dchest/authcookie"
)

var (
	// ErrTokenExpired declares a token as expired and invalid.
	ErrTokenExpired = errors.New("token already expired")
)

// Result represents to token to the outer world for HTTP responses.
type Result struct {
	Token     string
	ExpiresAt time.Time
}

// Token is internally used to differ between the kinds of tokens.
type Token struct {
	Text      string
	ExpiresAt time.Time
}

// Unlimited signs a token the never expires.
func (t *Token) Unlimited(secret string) (*Result, error) {
	return t.Expiring(secret, 0)
}

// Expiring signs a token that maybe expires.
func (t *Token) Expiring(secret string, exp time.Duration) (*Result, error) {
	if exp > 0 {
		expire := time.Now().Add(exp)

		return &Result{
			Token:     authcookie.New(t.Text, expire, []byte(secret)),
			ExpiresAt: expire,
		}, nil
	}

	return &Result{
		Token: authcookie.New(t.Text, time.Time{}, []byte(secret)),
	}, nil
}

// New initializes a new simple token of a specified kind.
func New(text string) *Token {
	return &Token{
		Text: text,
	}
}

// Parse can parse the token directly without a request.
func Parse(cookie, secret string) (*Token, error) {
	login, expires, err := authcookie.Parse(cookie, []byte(secret))

	if err != nil {
		return nil, err
	}

	if !expires.IsZero() && expires.Before(time.Now()) {
		return nil, ErrTokenExpired
	}

	return &Token{
		Text:      login,
		ExpiresAt: expires,
	}, nil
}
