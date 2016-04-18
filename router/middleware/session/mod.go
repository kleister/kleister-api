package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"
	"github.com/solderapp/solder-api/store"
)

const (
	// ModContextKey defines the context key that stores the mod.
	ModContextKey = "mod"
)

// Mod gets the mod from the context.
func Mod(c *gin.Context) *model.Mod {
	v, ok := c.Get(ModContextKey)

	if !ok {
		return nil
	}

	r, ok := v.(*model.Mod)

	if !ok {
		return nil
	}

	return r
}

// SetMod injects the mod into the context.
func SetMod() gin.HandlerFunc {
	return func(c *gin.Context) {
		record, res := store.GetMod(
			c,
			c.Param("mod"),
		)

		if res.Error != nil || res.RecordNotFound() {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  http.StatusNotFound,
					"message": "Failed to find mod",
				},
			)

			c.Abort()
		} else {
			c.Set(ModContextKey, record)
			c.Next()
		}
	}
}

// MustMods validates the mods access.
func MustMods(action string) gin.HandlerFunc {
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
			case action == "display" && user.Permission.DisplayMods:
				c.Next()
			case action == "change" && user.Permission.ChangeMods:
				c.Next()
			case action == "delete" && user.Permission.DeleteMods:
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
