package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder/model"
	"github.com/solderapp/solder/router/middleware/context"
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
		record := &model.Client{}

		res := context.Store(c).Where(
			"clients.id = ?",
			c.Param("client"),
		).Or(
			"clients.slug = ?",
			c.Param("client"),
		).First(
			&record,
		)

		if res.RecordNotFound() {
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
