package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kleister/kleister-api/config"
)

// Index represents the index page.
func Index(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"Root": config.Server.Root,
		},
	)
}
