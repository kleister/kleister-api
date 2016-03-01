package context

import (
	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder/config"
	"github.com/solderapp/solder/store"
)

func Config(c *gin.Context) config.Config {
	return c.MustGet("config").(config.Config)
}

func SetConfig(config config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("config", config)
		c.Next()
	}
}

func Store(c *gin.Context) store.Store {
	return c.MustGet("store").(store.Store)
}

func SetStore(store store.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("store", store)
		c.Next()
	}
}
