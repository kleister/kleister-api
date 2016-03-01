package logger

import (
	"github.com/gin-gonic/gin"
)

func SetLogger() gin.HandlerFunc {
	return gin.Logger()
}
