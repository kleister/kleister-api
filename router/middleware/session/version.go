package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder/model"
	"github.com/solderapp/solder/router/middleware/context"
)

const (
	// VersionContextKey defines the context key that stores the version.
	VersionContextKey = "version"
)

// Version gets the version from the context.
func Version(c *gin.Context) *model.Version {
	v, ok := c.Get(VersionContextKey)

	if !ok {
		return nil
	}

	r, ok := v.(*model.Version)

	if !ok {
		return nil
	}

	return r
}

// SetVersion injects the version into the context.
func SetVersion() gin.HandlerFunc {
	return func(c *gin.Context) {
		record := &model.Version{}

		res := context.Store(c).Where(
			"versions.id = ?",
			c.Param("version"),
		).Or(
			"versions.slug = ?",
			c.Param("version"),
		).First(
			&record,
		)

		if res.RecordNotFound() {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  http.StatusNotFound,
					"message": "Failed to find version",
				},
			)

			c.Abort()
		} else {
			c.Set(VersionContextKey, record)
			c.Next()
		}
	}
}
