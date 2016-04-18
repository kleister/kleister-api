package store

import (
	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/store"
	"github.com/solderapp/solder-api/store/data"
)

// SetStore injects the storage into the context.
func SetStore() gin.HandlerFunc {
	db := data.Load()

	return func(c *gin.Context) {
		store.ToContext(c, db)
		c.Next()
	}
}
