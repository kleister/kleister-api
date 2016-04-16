package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/router/middleware/context"
)

// GetIndex represents the index page.
func GetIndex(c *gin.Context) {
	cfg := context.Config(c)

	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"Root": cfg.Server.Root,
		},
	)
}
