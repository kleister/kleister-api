package recovery

import (
	"github.com/gin-gonic/gin"
)

// SetRecovery initializes the recovery middleware.
func SetRecovery() gin.HandlerFunc {
	return gin.Recovery()
}
