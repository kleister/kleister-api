package context

import (
	"net/url"
	"strings"

	"github.com/drone/gin-location"
	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/config"
	"github.com/solderapp/solder-api/model"
)

// Config gets the config from the context.
func Config(c *gin.Context) config.Config {
	return c.MustGet("config").(config.Config)
}

// SetConfig injects the config into the context.
func SetConfig(config config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("config", config)
		c.Next()
	}
}

// Store gets the storage from the context.
func Store(c *gin.Context) model.Store {
	return c.MustGet("store").(model.Store)
}

// SetStore injects the storage into the context.
func SetStore(store model.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("store", store)
		c.Next()
	}
}

// Location gets the location from the context.
func Location(c *gin.Context) *url.URL {
	return location.Get(c)
}

// SetLocation injects the location into the context.
func SetLocation() gin.HandlerFunc {
	return location.New(location.Config{
		Host: "localhost:8080",
		Base: strings.TrimSuffix(
			config.Values.Server.Root,
			"/",
		),
	})
}
