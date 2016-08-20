package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kleister/kleister-api/model"
	"github.com/kleister/kleister-api/store"
)

const (
	// PackContextKey defines the context key that stores the pack.
	PackContextKey = "pack"
)

// Pack gets the pack from the context.
func Pack(c *gin.Context) *model.Pack {
	v, ok := c.Get(PackContextKey)

	if !ok {
		return nil
	}

	r, ok := v.(*model.Pack)

	if !ok {
		return nil
	}

	return r
}

// SetPack injects the pack into the context.
func SetPack() gin.HandlerFunc {
	return func(c *gin.Context) {
		record, res := store.GetPack(
			c,
			c.Param("pack"),
		)

		if res.Error != nil || res.RecordNotFound() {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  http.StatusNotFound,
					"message": "Failed to find pack",
				},
			)

			c.Abort()
		} else {
			c.Set(PackContextKey, record)
			c.Next()
		}
	}
}

// MustPacks validates the packs access.
func MustPacks(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		current := Current(c)

		if current.Admin {
			c.Next()
			return
		}

		switch {
		case action == "display":
			if allowPackDisplay(c) {
				c.Next()
				return
			}
		case action == "change":
			if allowPackChange(c) {
				c.Next()
				return
			}
		case action == "delete":
			if allowPackDelete(c) {
				c.Next()
				return
			}
		}

		AbortUnauthorized(c)
	}
}

// allowPackDisplay checks if the given user is allowed to display the resource.
func allowPackDisplay(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowPackChange checks if the given user is allowed to change the resource.
func allowPackChange(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowPackDelete checks if the given user is allowed to delete the resource.
func allowPackDelete(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// MustPackClients validates the minecraft builds access.
func MustPackClients(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		current := Current(c)

		if current.Admin {
			c.Next()
			return
		}

		switch {
		case action == "display":
			if allowPackClientDisplay(c) {
				c.Next()
				return
			}
		case action == "change":
			if allowPackClientChange(c) {
				c.Next()
				return
			}
		}

		AbortUnauthorized(c)
	}
}

// allowPackClientDisplay checks if the given user is allowed to display the resource.
func allowPackClientDisplay(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowPackClientChange checks if the given user is allowed to change the resource.
func allowPackClientChange(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// MustPackUsers validates the minecraft builds access.
func MustPackUsers(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		current := Current(c)

		if current.Admin {
			c.Next()
			return
		}

		switch {
		case action == "display":
			if allowPackUserDisplay(c) {
				c.Next()
				return
			}
		case action == "change":
			if allowPackUserChange(c) {
				c.Next()
				return
			}
		}

		AbortUnauthorized(c)
	}
}

// allowPackUserDisplay checks if the given user is allowed to display the resource.
func allowPackUserDisplay(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowPackUserChange checks if the given user is allowed to change the resource.
func allowPackUserChange(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// MustPackTeams validates the minecraft builds access.
func MustPackTeams(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		current := Current(c)

		if current.Admin {
			c.Next()
			return
		}

		switch {
		case action == "display":
			if allowPackTeamDisplay(c) {
				c.Next()
				return
			}
		case action == "change":
			if allowPackTeamChange(c) {
				c.Next()
				return
			}
		}

		AbortUnauthorized(c)
	}
}

// allowPackTeamDisplay checks if the given user is allowed to display the resource.
func allowPackTeamDisplay(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowPackTeamChange checks if the given user is allowed to change the resource.
func allowPackTeamChange(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}
