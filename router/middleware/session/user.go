package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder/model"
)

// User gets the user from the context.
func User(c *gin.Context) *model.User {
	v, ok := c.Get("user")

	if !ok {
		return nil
	}

	u, ok := v.(*model.User)

	if !ok {
		return nil
	}

	return u
}

// SetUser injects the user into the context.
func SetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO(must): Set user in the context
		c.Next()
	}
}

// MustAdmin validates the the user is an admin.
func MustAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := User(c)

		switch {
		case user == nil:
			c.AbortWithStatus(http.StatusUnauthorized)
		case !user.Permission.Admin:
			c.AbortWithStatus(http.StatusForbidden)
		default:
			c.Next()
		}
	}
}

// MustUser validates the the user is authed.
func MustUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := User(c)

		switch {
		case user == nil:
			c.AbortWithStatus(http.StatusUnauthorized)
		default:
			c.Next()
		}
	}
}
