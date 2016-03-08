package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder/model"
	"github.com/solderapp/solder/router/middleware/context"
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
		record := &model.Build{}

		res := context.Store(c).Where(
			"builds.id = ?",
			c.Param("build"),
		).Or(
			"builds.slug = ?",
			c.Param("build"),
		).First(
			&record,
		)

		if res.RecordNotFound() {
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
