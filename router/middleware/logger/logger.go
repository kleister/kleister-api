package logger

import (
	"github.com/gin-gonic/gin"
)

// SetLogger initializes the logging middleware.
func SetLogger() gin.HandlerFunc {
	return gin.Logger()
}
