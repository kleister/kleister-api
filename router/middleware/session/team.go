package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kleister/kleister-api/model"
	"github.com/kleister/kleister-api/store"
)

const (
	// TeamContextKey defines the context key that stores the team.
	TeamContextKey = "team"
)

// Team gets the team from the context.
func Team(c *gin.Context) *model.Team {
	v, ok := c.Get(TeamContextKey)

	if !ok {
		return nil
	}

	r, ok := v.(*model.Team)

	if !ok {
		return nil
	}

	return r
}

// SetTeam injects the team into the context.
func SetTeam() gin.HandlerFunc {
	return func(c *gin.Context) {
		record, res := store.GetTeam(
			c,
			c.Param("team"),
		)

		if res.Error != nil || res.RecordNotFound() {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  http.StatusNotFound,
					"message": "Failed to find team",
				},
			)

			c.Abort()
		} else {
			c.Set(TeamContextKey, record)
			c.Next()
		}
	}
}

// MustTeams validates the teams access.
func MustTeams(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		current := Current(c)

		if current.Admin {
			c.Next()
			return
		}

		switch {
		case action == "display":
			if allowTeamDisplay(c) {
				c.Next()
				return
			}
		case action == "change":
			if allowTeamChange(c) {
				c.Next()
				return
			}
		case action == "delete":
			if allowTeamDelete(c) {
				c.Next()
				return
			}
		}

		AbortUnauthorized(c)
	}
}

// allowTeamDisplay checks if the given user is allowed to display the resource.
func allowTeamDisplay(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowTeamChange checks if the given user is allowed to change the resource.
func allowTeamChange(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowTeamDelete checks if the given user is allowed to delete the resource.
func allowTeamDelete(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// MustTeamUsers validates the minecraft builds access.
func MustTeamUsers(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		current := Current(c)

		if current.Admin {
			c.Next()
			return
		}

		switch {
		case action == "display":
			if allowTeamUserDisplay(c) {
				c.Next()
				return
			}
		case action == "change":
			if allowTeamUserChange(c) {
				c.Next()
				return
			}
		}

		AbortUnauthorized(c)
	}
}

// allowTeamUserDisplay checks if the given user is allowed to display the resource.
func allowTeamUserDisplay(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowTeamUserChange checks if the given user is allowed to change the resource.
func allowTeamUserChange(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// MustTeamPacks validates the minecraft builds access.
func MustTeamPacks(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		current := Current(c)

		if current.Admin {
			c.Next()
			return
		}

		switch {
		case action == "display":
			if allowTeamPackDisplay(c) {
				c.Next()
				return
			}
		case action == "change":
			if allowTeamPackChange(c) {
				c.Next()
				return
			}
		}

		AbortUnauthorized(c)
	}
}

// allowTeamPackDisplay checks if the given user is allowed to display the resource.
func allowTeamPackDisplay(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowTeamPackChange checks if the given user is allowed to change the resource.
func allowTeamPackChange(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// MustTeamMods validates the minecraft builds access.
func MustTeamMods(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		current := Current(c)

		if current.Admin {
			c.Next()
			return
		}

		switch {
		case action == "display":
			if allowTeamModDisplay(c) {
				c.Next()
				return
			}
		case action == "change":
			if allowTeamModChange(c) {
				c.Next()
				return
			}
		}

		AbortUnauthorized(c)
	}
}

// allowTeamModDisplay checks if the given user is allowed to display the resource.
func allowTeamModDisplay(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowTeamModChange checks if the given user is allowed to change the resource.
func allowTeamModChange(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}
