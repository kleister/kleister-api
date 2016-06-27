package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/assets"
)

// Favicon represents the favicon.
func Favicon(c *gin.Context) {
	c.Data(
		http.StatusOK,
		"image/x-icon",
		assets.MustAsset(
			"images/favicon.ico",
		),
	)
}
