package recovery

import (
	"github.com/gin-gonic/gin"
)

// SetRecover initializes the recovery middleware.
func SetRecovery() gin.HandlerFunc {
	return gin.Recovery()
}
