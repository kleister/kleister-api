package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder/model"
	"github.com/solderapp/solder/router/middleware/context"
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
		record := &model.Key{}

		res := context.Store(c).Where(
			"keys.id = ?",
			c.Param("key"),
		).Or(
			"keys.slug = ?",
			c.Param("key"),
		).First(
			&record,
		)

		if res.RecordNotFound() {
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
			case action == "display" && user.Permission.DisplayKeys:
				c.Next()
			case action == "change" && user.Permission.ChangeKeys:
				c.Next()
			case action == "delete" && user.Permission.DeleteKeys:
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
