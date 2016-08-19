package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kleister/kleister-api/model"
	"github.com/kleister/kleister-api/store"
)

const (
	// ClientContextKey defines the context key that stores the client.
	ClientContextKey = "client"
)

// Client gets the client from the context.
func Client(c *gin.Context) *model.Client {
	v, ok := c.Get(ClientContextKey)

	if !ok {
		return nil
	}

	r, ok := v.(*model.Client)

	if !ok {
		return nil
	}

	return r
}

// SetClient injects the client into the context.
func SetClient() gin.HandlerFunc {
	return func(c *gin.Context) {
		record, res := store.GetClient(
			c,
			c.Param("client"),
		)

		if res.Error != nil || res.RecordNotFound() {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  http.StatusNotFound,
					"message": "Failed to find client",
				},
			)

			c.Abort()
		} else {
			c.Set(ClientContextKey, record)
			c.Next()
		}
	}
}

// MustClients validates the clients access.
func MustClients(action string) gin.HandlerFunc {
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
			case action == "display": // && user.Permission.DisplayClients:
				c.Next()
			case action == "change": // && user.Permission.ChangeClients:
				c.Next()
			case action == "delete": // && user.Permission.DeleteClients:
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
