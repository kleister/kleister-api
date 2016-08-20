package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kleister/kleister-api/model"
	"github.com/kleister/kleister-api/store"
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

		record, res := store.GetVersion(
			c,
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

// MustVersions validates the versions access.
func MustVersions(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		current := Current(c)

		if current.Admin {
			c.Next()
			return
		}

		switch {
		case action == "display":
			if allowVersionDisplay(c) {
				c.Next()
				return
			}
		case action == "change":
			if allowVersionChange(c) {
				c.Next()
				return
			}
		case action == "delete":
			if allowVersionDelete(c) {
				c.Next()
				return
			}
		}

		AbortUnauthorized(c)
	}
}

// allowVersionDisplay checks if the given user is allowed to display the resource.
func allowVersionDisplay(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowVersionChange checks if the given user is allowed to change the resource.
func allowVersionChange(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowVersionDelete checks if the given user is allowed to delete the resource.
func allowVersionDelete(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// MustVersionBuilds validates the minecraft builds access.
func MustVersionBuilds(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		current := Current(c)

		if current.Admin {
			c.Next()
			return
		}

		switch {
		case action == "display":
			if allowVersionBuildDisplay(c) {
				c.Next()
				return
			}
		case action == "change":
			if allowVersionBuildChange(c) {
				c.Next()
				return
			}
		}

		AbortUnauthorized(c)
	}
}

// allowVersionBuildDisplay checks if the given user is allowed to display the resource.
func allowVersionBuildDisplay(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowVersionBuildChange checks if the given user is allowed to change the resource.
func allowVersionBuildChange(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}
