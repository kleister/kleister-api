package context

import (
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/config"
	"github.com/solderapp/solder-api/model"
	"github.com/solderapp/solder-api/store"
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
	users := 0

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

	store.Model(
		&model.Build{},
	).AddUniqueIndex(
		"uix_builds_pack_id_slug",
		"pack_id",
		"slug",
	)

	store.Model(
		&model.Build{},
	).AddUniqueIndex(
		"uix_builds_pack_id_name",
		"pack_id",
		"name",
	)

	store.Model(
		&model.Version{},
	).AddUniqueIndex(
		"uix_versions_mod_id_slug",
		"mod_id",
		"slug",
	)

	store.Model(
		&model.Version{},
	).AddUniqueIndex(
		"uix_versions_mod_id_name",
		"mod_id",
		"name",
	)

	store.Model(
		&model.User{},
	).Count(
		&users,
	)

	if users == 0 {
		record := &model.User{
			Username: "admin",
			Password: "admin",
			Email:    "admin@example.com",
			Permission: &model.Permission{
				DisplayUsers:   true,
				ChangeUsers:    true,
				DeleteUsers:    true,
				DisplayKeys:    true,
				ChangeKeys:     true,
				DeleteKeys:     true,
				DisplayClients: true,
				ChangeClients:  true,
				DeleteClients:  true,
				DisplayPacks:   true,
				ChangePacks:    true,
				DeletePacks:    true,
				DisplayMods:    true,
				ChangeMods:     true,
				DeleteMods:     true,
			},
		}

		err := store.Create(&record).Error

		if err != nil {
			logrus.Errorf(
				"Failed to create initial user. %s",
				err.Error(),
			)
		}
	}

	return func(c *gin.Context) {
		c.Set("store", store)
		c.Next()
	}
}
