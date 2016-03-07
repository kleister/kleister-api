package context

import (
	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder/config"
	"github.com/solderapp/solder/model"
	"github.com/solderapp/solder/store"
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
func Store(c *gin.Context) store.Store {
	return c.MustGet("store").(store.Store)
}

// SetStore injects the storage into the context.
func SetStore(store store.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		store.AutoMigrate(
			&model.Attachment{},
			&model.Build{},
			&model.Client{},
			&model.Forge{},
			&model.Key{},
			&model.Minecraft{},
			&model.Mod{},
			&model.Pack{},
			&model.Permission{},
			&model.User{},
			&model.Version{},
		)

		c.Set("store", store)
		c.Next()
	}
}
