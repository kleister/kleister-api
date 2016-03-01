package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder/router/middleware/context"
)

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

func GetAPI(c *gin.Context) {
	config := context.Config(c)

	c.IndentedJSON(
		http.StatusOK,
		gin.H{
			"stream":  "reloaded",
			"api":     "SolderNG",
			"version": config.Version,
		},
	)
}
