package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/solderapp/solder-api.v0/model"
	"gopkg.in/solderapp/solder-api.v0/router/middleware/context"
)

const (
	// PackContextKey defines the context key that stores the pack.
	PackContextKey = "pack"
)

// Pack gets the pack from the context.
func Pack(c *gin.Context) *model.Pack {
	v, ok := c.Get(PackContextKey)

	if !ok {
		return nil
	}

	r, ok := v.(*model.Pack)

	if !ok {
		return nil
	}

	return r
}

// SetPack injects the pack into the context.
func SetPack() gin.HandlerFunc {
	return func(c *gin.Context) {
		record := &model.Pack{}

		res := context.Store(c).Where(
			"packs.id = ?",
			c.Param("pack"),
		).Or(
			"packs.slug = ?",
			c.Param("pack"),
		).First(
			&record,
		)

		// .Model(
		// 	&model.Pack{},
		// ).Preload(
		// 	"Builds",
		// )

		if res.RecordNotFound() {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  http.StatusNotFound,
					"message": "Failed to find pack",
				},
			)

			c.Abort()
		} else {
			c.Set(PackContextKey, record)
			c.Next()
		}
	}
}

// MustPacks validates the packs access.
func MustPacks(action string) gin.HandlerFunc {
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
			case action == "display" && user.Permission.DisplayPacks:
				c.Next()
			case action == "change" && user.Permission.ChangePacks:
				c.Next()
			case action == "delete" && user.Permission.DeletePacks:
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
