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
		user := Current(c)

		if user == nil {
			c.JSON(
				http.StatusUnauthorized,
				gin.H{
					"status":  http.StatusUnauthorized,
					"message": "You have to be authenticated",
				},
			)

			c.Abort()
		} else {
			switch {
			case action == "display": // && user.Permission.DisplayTeams:
				c.Next()
			case action == "change": // && user.Permission.ChangeTeams:
				c.Next()
			case action == "delete": // && user.Permission.DeleteTeams:
				c.Next()
			default:
				c.JSON(
					http.StatusForbidden,
					gin.H{
						"status":  http.StatusForbidden,
						"message": "You are not authorized to request this resource",
					},
				)

				c.Abort()
			}
		}
	}
}
