package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kleister/kleister-api/model"
	"github.com/kleister/kleister-api/store"
)

const (
	// MinecraftContextKey defines the context key that stores the minecraft.
	MinecraftContextKey = "minecraft"
)

// Minecraft gets the minecraft from the context.
func Minecraft(c *gin.Context) *model.Minecraft {
	v, ok := c.Get(MinecraftContextKey)

	if !ok {
		return nil
	}

	r, ok := v.(*model.Minecraft)

	if !ok {
		return nil
	}

	return r
}

// SetMinecraft injects the minecraft into the context.
func SetMinecraft() gin.HandlerFunc {
	return func(c *gin.Context) {
		record, res := store.GetMinecraft(
			c,
			c.Param("minecraft"),
		)

		if res.Error != nil || res.RecordNotFound() {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  http.StatusNotFound,
					"message": "Failed to find Minecraft version",
				},
			)

			c.Abort()
		} else {
			c.Set(MinecraftContextKey, record)
			c.Next()
		}
	}
}

// MustMinecraftBuilds validates the minecraft builds access.
func MustMinecraftBuilds(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		current := Current(c)

		if current.Admin {
			c.Next()
			return
		}

		switch {
		case action == "display":
			if allowMinecraftBuildDisplay(c) {
				c.Next()
				return
			}
		case action == "change":
			if allowMinecraftBuildChange(c) {
				c.Next()
				return
			}
		}

		AbortUnauthorized(c)
	}
}

// allowMinecraftBuildDisplay checks if the given user is allowed to display the resource.
func allowMinecraftBuildDisplay(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowMinecraftBuildChange checks if the given user is allowed to change the resource.
func allowMinecraftBuildChange(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}
