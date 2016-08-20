package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kleister/kleister-api/model"
	"github.com/kleister/kleister-api/store"
)

const (
	// ForgeContextKey defines the context key that stores the forge.
	ForgeContextKey = "forge"
)

// Forge gets the forge from the context.
func Forge(c *gin.Context) *model.Forge {
	v, ok := c.Get(ForgeContextKey)

	if !ok {
		return nil
	}

	r, ok := v.(*model.Forge)

	if !ok {
		return nil
	}

	return r
}

// SetForge injects the forge into the context.
func SetForge() gin.HandlerFunc {
	return func(c *gin.Context) {
		record, res := store.GetForge(
			c,
			c.Param("forge"),
		)

		if res.Error != nil || res.RecordNotFound() {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  http.StatusNotFound,
					"message": "Failed to find Forge version",
				},
			)

			c.Abort()
		} else {
			c.Set(ForgeContextKey, record)
			c.Next()
		}
	}
}

// MustForgeBuilds validates the forge builds access.
func MustForgeBuilds(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		current := Current(c)

		if current.Admin {
			c.Next()
			return
		}

		switch {
		case action == "display":
			if allowForgeBuildDisplay(c) {
				c.Next()
				return
			}
		case action == "change":
			if allowForgeBuildChange(c) {
				c.Next()
				return
			}
		}

		AbortUnauthorized(c)
	}
}

// allowForgeBuildDisplay checks if the given user is allowed to display the resource.
func allowForgeBuildDisplay(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowForgeBuildChange checks if the given user is allowed to change the resource.
func allowForgeBuildChange(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}
