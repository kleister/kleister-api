package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"
	"github.com/solderapp/solder-api/router/middleware/context"
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
		mod := Mod(c)

		if mod == nil {
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

		record, res := context.Store(c).GetVersion(
			mod.ID,
			c.Param("version"),
		)

		if res.Error != nil || res.RecordNotFound() {
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
