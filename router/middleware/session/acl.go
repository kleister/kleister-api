package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"
	"github.com/solderapp/solder-api/router/middleware/context"
)

const (
	// CurrentContextKey defines the context key that stores the user.
	CurrentContextKey = "current"
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
		record, res := context.Store(c).GetUser(
			"admin",
		)

		if res.Error != nil || res.RecordNotFound() {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  http.StatusNotFound,
					"message": "Failed to find current",
				},
			)

			c.Abort()
		} else {
			c.Set(CurrentContextKey, record)
			c.Next()
		}
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
		} else {
			c.Next()
		}
	}
}
