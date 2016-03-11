package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"
	"github.com/solderapp/solder-api/router/middleware/context"
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
		record := &model.Minecraft{}

		res := context.Store(c).Where(
			"minecrafts.id = ?",
			c.Param("minecraft"),
		).Or(
			"minecrafts.slug = ?",
			c.Param("minecraft"),
		).First(
			&record,
		)

		if res.RecordNotFound() {
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
