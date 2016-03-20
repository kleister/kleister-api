package context

import (
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
