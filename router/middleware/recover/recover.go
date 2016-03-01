package recover

import (
	"github.com/gin-gonic/gin"
)

func SetRecover() gin.HandlerFunc {
	return gin.Recovery()
}
