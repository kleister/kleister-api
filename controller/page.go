package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/solderapp/solder-api.v0/router/middleware/context"
)

// GetIndex represents the index page.
func GetIndex(c *gin.Context) {
	config := context.Config(c)

	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"Root": config.Server.Root,
		},
	)
}

// GetAPI represents the API index.
func GetAPI(c *gin.Context) {
	config := context.Config(c)

	c.JSON(
		http.StatusOK,
		gin.H{
			"stream":  "reloaded",
			"api":     "SolderNG",
			"version": config.Version,
		},
	)
}
