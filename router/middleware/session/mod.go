package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kleister/kleister-api/model"
	"github.com/kleister/kleister-api/store"
)

const (
	// ModContextKey defines the context key that stores the mod.
	ModContextKey = "mod"
)

// Mod gets the mod from the context.
func Mod(c *gin.Context) *model.Mod {
	v, ok := c.Get(ModContextKey)

	if !ok {
		return nil
	}

	r, ok := v.(*model.Mod)

	if !ok {
		return nil
	}

	return r
}

// SetMod injects the mod into the context.
func SetMod() gin.HandlerFunc {
	return func(c *gin.Context) {
		record, res := store.GetMod(
			c,
			c.Param("mod"),
		)

		if res.Error != nil || res.RecordNotFound() {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  http.StatusNotFound,
					"message": "Failed to find mod",
				},
			)

			c.Abort()
		} else {
			c.Set(ModContextKey, record)
			c.Next()
		}
	}
}

// MustMods validates the mods access.
func MustMods(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		current := Current(c)

		if current.Admin {
			c.Next()
			return
		}

		switch {
		case action == "display":
			if allowModDisplay(current, c) {
				c.Next()
				return
			}
		case action == "change":
			if allowModChange(current, c) {
				c.Next()
				return
			}
		case action == "delete":
			if allowModDelete(current, c) {
				c.Next()
				return
			}
		}

		AbortUnauthorized(c)
	}
}

// allowModDisplay checks if the given user is allowed to display the resource.
func allowModDisplay(current *model.User, c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowModChange checks if the given user is allowed to change the resource.
func allowModChange(current *model.User, c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowModDelete checks if the given user is allowed to delete the resource.
func allowModDelete(current *model.User, c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// MustModUsers validates the minecraft mods access.
func MustModUsers(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		current := Current(c)

		if current.Admin {
			c.Next()
			return
		}

		switch {
		case action == "display":
			if allowModUserDisplay(current, c) {
				c.Next()
				return
			}
		case action == "change":
			if allowModUserChange(current, c) {
				c.Next()
				return
			}
		}

		AbortUnauthorized(c)
	}
}

// allowModUserDisplay checks if the given user is allowed to display the resource.
func allowModUserDisplay(current *model.User, c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowModUserChange checks if the given user is allowed to change the resource.
func allowModUserChange(current *model.User, c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// MustModTeams validates the minecraft mods access.
func MustModTeams(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		current := Current(c)

		if current.Admin {
			c.Next()
			return
		}

		switch {
		case action == "display":
			if allowModTeamDisplay(current, c) {
				c.Next()
				return
			}
		case action == "change":
			if allowModTeamChange(current, c) {
				c.Next()
				return
			}
		}

		AbortUnauthorized(c)
	}
}

// allowModTeamDisplay checks if the given user is allowed to display the resource.
func allowModTeamDisplay(current *model.User, c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowModTeamChange checks if the given user is allowed to change the resource.
func allowModTeamChange(current *model.User, c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}
