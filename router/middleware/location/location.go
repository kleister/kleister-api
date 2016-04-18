package location

import (
	"net/url"
	"strings"

	"github.com/drone/gin-location"
	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/config"
)

// Location gets the location from the context.
func Location(c *gin.Context) *url.URL {
	return location.Get(c)
}

// SetLocation injects the location into the context.
func SetLocation() gin.HandlerFunc {
	return location.New(location.Config{
		Host: "localhost:8080",
		Base: strings.TrimSuffix(
			config.Server.Root,
			"/",
		),
	})
}
