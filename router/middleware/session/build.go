package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"
	"github.com/solderapp/solder-api/store"
)

const (
	// BuildContextKey defines the context key that stores the mod.
	BuildContextKey = "build"
)

// Build gets the build from the context.
func Build(c *gin.Context) *model.Build {
	v, ok := c.Get(BuildContextKey)

	if !ok {
		return nil
	}

	r, ok := v.(*model.Build)

	if !ok {
		return nil
	}

	return r
}

// SetBuild injects the build into the context.
func SetBuild() gin.HandlerFunc {
	return func(c *gin.Context) {
		pack := Pack(c)

		if pack == nil {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  http.StatusNotFound,
					"message": "Failed to find parent",
				},
			)

			c.Abort()
			return
		}

		record, res := store.GetBuild(
			c,
			pack.ID,
			c.Param("build"),
		)

		if res.Error != nil || res.RecordNotFound() {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  http.StatusNotFound,
					"message": "Failed to find build",
				},
			)

			c.Abort()
		} else {
			c.Set(BuildContextKey, record)
			c.Next()
		}
	}
}
