package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kleister/kleister-api/pkg/version"
)

// IndexInfo represents the API index.
func IndexInfo(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"api":     "Kleister API",
			"version": version.Version.String(),
			"stream":  "master",
		},
	)
}
