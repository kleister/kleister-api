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
		current := Current(c)

		if current.Admin {
			c.Next()
			return
		}

		switch {
		case action == "display":
			if allowClientDisplay(current, c) {
				c.Next()
				return
			}
		case action == "change":
			if allowClientChange(current, c) {
				c.Next()
				return
			}
		case action == "delete":
			if allowClientDelete(current, c) {
				c.Next()
				return
			}
		}

		AbortUnauthorized(c)
	}
}

// allowClientDisplay checks if the given user is allowed to display the resource.
func allowClientDisplay(current *model.User, c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowClientChange checks if the given user is allowed to change the resource.
func allowClientChange(current *model.User, c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowClientDelete checks if the given user is allowed to delete the resource.
func allowClientDelete(current *model.User, c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// MustClientPacks validates the minecraft packs access.
func MustClientPacks(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		current := Current(c)

		if current.Admin {
			c.Next()
			return
		}

		switch {
		case action == "display":
			if allowClientPackDisplay(current, c) {
				c.Next()
				return
			}
		case action == "change":
			if allowClientPackChange(current, c) {
				c.Next()
				return
			}
		}

		AbortUnauthorized(c)
	}
}

// allowClientPackDisplay checks if the given user is allowed to display the resource.
func allowClientPackDisplay(current *model.User, c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowClientPackChange checks if the given user is allowed to change the resource.
func allowClientPackChange(current *model.User, c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}
