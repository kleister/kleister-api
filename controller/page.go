package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/router/middleware/context"
	"github.com/solderapp/solder-api/static"
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

// GetFavicon represents the favicon.
func GetFavicon(c *gin.Context) {
	c.Data(
		http.StatusOK,
		"image/x-icon",
		static.MustAsset(
			"images/favicon.ico",
		),
	)
}
