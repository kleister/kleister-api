package recover

import (
	"github.com/gin-gonic/gin"
)

// SetRecover initializes the recovery middleware.
func SetRecover() gin.HandlerFunc {
	return gin.Recovery()
}
