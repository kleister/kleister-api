package session

import (
	"github.com/gin-gonic/gin"
	"github.com/kleister/kleister-api/store"
)

// SetSolder injects the client and key into the context.
func SetSolder() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Query("k") != "" {
			record, res := store.GetKeyByValue(
				c,
				c.Query("k"),
			)

			if !res.RecordNotFound() {
				c.Set(KeyContextKey, record)
			}
		}

		if c.Query("cid") != "" {
			record, res := store.GetClientByValue(
				c,
				c.Query("cid"),
			)

			if !res.RecordNotFound() {
				c.Set(ClientContextKey, record)
			}
		}

		c.Next()
	}
}
