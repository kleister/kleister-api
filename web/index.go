package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/config"
)

// GetIndex represents the index page.
func GetIndex(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"Root": config.Server.Root,
		},
	)
}
