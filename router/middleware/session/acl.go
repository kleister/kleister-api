package session

import (
	"encoding/base32"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/config"
	"github.com/kleister/kleister-api/model"
	"github.com/kleister/kleister-api/shared/token"
	"github.com/kleister/kleister-api/store"
)

const (
	// CurrentContextKey defines the context key that stores the user.
	CurrentContextKey = "current"

	// TokenContextKey defines the context key that stores the token.
	TokenContextKey = "token"
)

// Current gets the user from the context.
func Current(c *gin.Context) *model.User {
	v, ok := c.Get(CurrentContextKey)

	if !ok {
		return nil
	}

	r, ok := v.(*model.User)

	if !ok {
		return nil
	}

	return r
}

// SetCurrent injects the user into the context.
func SetCurrent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			record *model.User
		)

		parsed, err := token.Parse(
			c.Request,
			func(t *token.Token) ([]byte, error) {
				var (
					res *gorm.DB
				)

				record, res = store.GetUser(
					c,
					t.Text,
				)

				signingKey, _ := base32.StdEncoding.DecodeString(record.Hash)
				return signingKey, res.Error
			},
		)

		if err == nil {
			c.Set(TokenContextKey, parsed)
			c.Set(CurrentContextKey, record)
		}

		c.Next()
	}
}

// MustCurrent validates the user access.
func MustCurrent() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := Current(c)

		if user == nil {
			c.JSON(
				http.StatusUnauthorized,
				gin.H{
					"status":  http.StatusUnauthorized,
					"message": "You have to be authenticated",
				},
			)

			c.Abort()
			return
		}

		c.Next()
	}
}

// MustAdmin validates the admin access.
func MustAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := Current(c)

		if user == nil || !user.Admin || !isAdmin(user.Username) {
			c.JSON(
				http.StatusUnauthorized,
				gin.H{
					"status":  http.StatusUnauthorized,
					"message": "You have to be an admin user",
				},
			)

			c.Abort()
			return
		}

		c.Next()
	}
}

// MustNobody validates anonymous users.
func MustNobody() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := Current(c)

		if user != nil {
			c.JSON(
				http.StatusUnauthorized,
				gin.H{
					"status":  http.StatusUnauthorized,
					"message": "You have to be a guest user",
				},
			)

			c.Abort()
			return
		}

		c.Next()
	}
}

// AbortUnauthorized stops the middleware execution with JSON error.
func AbortUnauthorized(c *gin.Context) {
	c.JSON(
		http.StatusForbidden,
		gin.H{
			"status":  http.StatusForbidden,
			"message": "You are not authorized to request this resource",
		},
	)

	c.Abort()
}

func isAdmin(username string) bool {
	for _, admin := range config.Admin.Users {
		if admin == username {
			return true
		}
	}

	return false
}
