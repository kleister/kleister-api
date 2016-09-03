package token

import (
	"encoding/base32"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

const (
	// UserToken is the kind of token to represent a user token.
	UserToken = "user"

	// SessToken is the kind of token to represent a session token.
	SessToken = "sess"

	// SignerAlgo is the default algorithm used to sign JWT tokens.
	SignerAlgo = "HS256"
)

// SecretFunc is a helper function to retrieve the used JWT secret.
type SecretFunc func(*Token) ([]byte, error)

// Result represents to token to the outer world for HTTP responses.
type Result struct {
	Token  string `json:"token,omitempty"`
	Expire string `json:"expire,omitempty"`
}

// Token is internally used to differ between the kinds of tokens.
type Token struct {
	Kind string
	Text string
}

// SignUnlimited signs a token the never expires.
func (t *Token) SignUnlimited(secret string) (*Result, error) {
	return t.SignExpiring(secret, 0)
}

// SignExpiring signs a token that maybe expires.
func (t *Token) SignExpiring(secret string, exp time.Duration) (*Result, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["type"] = t.Kind
	claims["text"] = t.Text

	signingKey, _ := base32.StdEncoding.DecodeString(secret)
	tokenString, err := token.SignedString(signingKey)

	if exp > 0 {
		expire := time.Now().Add(exp)
		claims["exp"] = expire.Unix()

		return &Result{
			Token:  tokenString,
			Expire: expire.Format(time.RFC3339),
		}, err
	}

	return &Result{
		Token: tokenString,
	}, err
}

// New initializes a new simple token of a specified kind.
func New(kind, text string) *Token {
	return &Token{
		Kind: kind,
		Text: text,
	}
}

// Parse can parse the authorization information from a request.
func Parse(r *http.Request, fn SecretFunc) (*Token, error) {
	token := &Token{}

	raw, err := request.OAuth2Extractor.ExtractToken(r)

	if err != nil {
		return nil, err
	}

	parsed, err := jwt.Parse(raw, keyFunc(token, fn))

	if err != nil {
		return nil, err
	} else if !parsed.Valid {
		return nil, jwt.ValidationError{}
	}

	return token, nil
}

// Direct can parse the token directly without a request.
func Direct(val string, fn SecretFunc) (*Token, error) {
	token := &Token{}

	parsed, err := jwt.Parse(val, keyFunc(token, fn))

	if err != nil {
		return nil, err
	} else if !parsed.Valid {
		return nil, jwt.ValidationError{}
	}

	return token, nil
}

func keyFunc(token *Token, fn SecretFunc) jwt.Keyfunc {
	return func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != SignerAlgo {
			return nil, jwt.ErrSignatureInvalid
		}

		claims := t.Claims.(jwt.MapClaims)

		kindv, ok := claims["type"]

		if !ok {
			return nil, jwt.ValidationError{}
		}

		token.Kind, _ = kindv.(string)

		textv, ok := claims["text"]

		if !ok {
			return nil, jwt.ValidationError{}
		}

		token.Text, _ = textv.(string)

		return fn(token)
	}
}
