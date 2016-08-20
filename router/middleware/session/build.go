package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kleister/kleister-api/model"
	"github.com/kleister/kleister-api/store"
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

// MustBuilds validates the builds access.
func MustBuilds(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		current := Current(c)

		if current.Admin {
			c.Next()
			return
		}

		switch {
		case action == "display":
			if allowBuildDisplay(c) {
				c.Next()
				return
			}
		case action == "change":
			if allowBuildChange(c) {
				c.Next()
				return
			}
		case action == "delete":
			if allowBuildDelete(c) {
				c.Next()
				return
			}
		}

		AbortUnauthorized(c)
	}
}

// allowBuildDisplay checks if the given user is allowed to display the resource.
func allowBuildDisplay(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowBuildChange checks if the given user is allowed to change the resource.
func allowBuildChange(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowBuildDelete checks if the given user is allowed to delete the resource.
func allowBuildDelete(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// MustBuildVersions validates the minecraft builds access.
func MustBuildVersions(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		current := Current(c)

		if current.Admin {
			c.Next()
			return
		}

		switch {
		case action == "display":
			if allowBuildVersionDisplay(c) {
				c.Next()
				return
			}
		case action == "change":
			if allowBuildVersionChange(c) {
				c.Next()
				return
			}
		}

		AbortUnauthorized(c)
	}
}

// allowBuildVersionDisplay checks if the given user is allowed to display the resource.
func allowBuildVersionDisplay(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowBuildVersionChange checks if the given user is allowed to change the resource.
func allowBuildVersionChange(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}
