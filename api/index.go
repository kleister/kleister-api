package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/config"
)

// GetIndex represents the API index.
func GetIndex(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"api":     "Solder",
			"version": config.Version,
		},
	)
}
