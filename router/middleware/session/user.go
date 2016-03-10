package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/solderapp/solder-api.v0/model"
	"gopkg.in/solderapp/solder-api.v0/router/middleware/context"
)

const (
	// UserContextKey defines the context key that stores the user.
	UserContextKey = "user"
)

// User gets the user from the context.
func User(c *gin.Context) *model.User {
	v, ok := c.Get(UserContextKey)

	if !ok {
		return nil
	}

	r, ok := v.(*model.User)

	if !ok {
		return nil
	}

	return r
}

// SetUser injects the user into the context.
func SetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		record := &model.User{}

		res := context.Store(c).Where(
			"users.id = ?",
			c.Param("user"),
		).Or(
			"users.slug = ?",
			c.Param("user"),
		).Model(
			&record,
		).Preload(
			"Permission",
		).First(
			&record,
		)

		if res.RecordNotFound() {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  http.StatusNotFound,
					"message": "Failed to find user",
				},
			)

			c.Abort()
		} else {
			c.Set(UserContextKey, record)
			c.Next()
		}
	}
}

// MustUsers validates the users access.
func MustUsers(action string) gin.HandlerFunc {
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
		} else {
			switch {
			case action == "display" && user.Permission.DisplayUsers:
				c.Next()
			case action == "change" && user.Permission.ChangeUsers:
				c.Next()
			case action == "delete" && user.Permission.DeleteUsers:
				c.Next()
			default:
				c.JSON(
					http.StatusForbidden,
					gin.H{
						"status":  http.StatusForbidden,
						"message": "You are not authorized to request this resource",
					},
				)

				c.Abort()
			}
		}
	}
}
