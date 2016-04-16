package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/assets"
)

// GetFavicon represents the favicon.
func GetFavicon(c *gin.Context) {
	c.Data(
		http.StatusOK,
		"image/x-icon",
		assets.MustAsset(
			"images/favicon.ico",
		),
	)
}
