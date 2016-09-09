package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kleister/kleister-api/model"
	"github.com/kleister/kleister-api/store"
)

const (
	// KeyContextKey defines the context key that stores the key.
	KeyContextKey = "key"
)

// Key gets the key from the context.
func Key(c *gin.Context) *model.Key {
	v, ok := c.Get(KeyContextKey)

	if !ok {
		return nil
	}

	r, ok := v.(*model.Key)

	if !ok {
		return nil
	}

	return r
}

// SetKey injects the key into the context.
func SetKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		record, res := store.GetKey(
			c,
			c.Param("key"),
		)

		if res.Error != nil || res.RecordNotFound() {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  http.StatusNotFound,
					"message": "Failed to find key",
				},
			)

			c.Abort()
		} else {
			c.Set(KeyContextKey, record)
			c.Next()
		}
	}
}

// MustKeys validates the keys access.
func MustKeys(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		current := Current(c)

		if current.Admin {
			c.Next()
			return
		}

		switch {
		case action == "display":
			if allowKeyDisplay(c) {
				c.Next()
				return
			}
		case action == "change":
			if allowKeyChange(c) {
				c.Next()
				return
			}
		case action == "delete":
			if allowKeyDelete(c) {
				c.Next()
				return
			}
		}

		AbortUnauthorized(c)
	}
}

// allowKeyDisplay checks if the given user is allowed to display the resource.
func allowKeyDisplay(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowKeyChange checks if the given user is allowed to change the resource.
func allowKeyChange(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowKeyDelete checks if the given user is allowed to delete the resource.
func allowKeyDelete(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}
