package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dchest/authcookie"
	"github.com/golang-jwt/jwt/v5"
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

// Claims defines all required custom claims.
type Claims struct {
	Ident string `json:"ident"`
	Login string `json:"login"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

// Authed generates a new authenticated token.
func Authed(
	secret string,
	exp time.Duration,
	ident string,
	login string,
	email string,
	name string,
	admin bool,
) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		Claims{
			Ident: ident,
			Login: login,
			Email: email,
			Name:  name,
			Admin: admin,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(exp)),
				IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
				Issuer:    "kleister",
			},
		},
	)

	signed, err := token.SignedString(
		[]byte(secret),
	)

	if err != nil {
		return "", err
	}

	return signed, nil
}

// Verify simply tries to verify a given token.
func Verify(secret, token string) (*Claims, error) {
	result, err := jwt.ParseWithClaims(
		token,
		&Claims{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}

			return []byte(secret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	if claims, ok := result.Claims.(*Claims); ok && result.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
