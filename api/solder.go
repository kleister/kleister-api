package api

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/kleister/kleister-api/model/solder"
	"github.com/kleister/kleister-api/router/middleware/session"
	"github.com/kleister/kleister-api/store"
)

// SolderPacks retrieves the packs compatible to Technic Platform.
func SolderPacks(c *gin.Context) {
	records, _ := store.GetSolderPacks(
		c,
	)

	c.JSON(
		http.StatusOK,
		solder.NewPacksFromList(
			records,
			session.Client(c),
			session.Key(c),
			c.Query("include"),
		),
	)
}

// SolderPack retrieves the pack compatible to Technic Platform.
func SolderPack(c *gin.Context) {
	record, err := store.GetSolderPack(
		c,
		c.Param("pack"),
	)

	if err != nil {
		logrus.Warnf("Failed to fetch solder pack. %s", err)

		c.JSON(
			http.StatusOK,
			gin.H{
				"error": "Modpack does not exist",
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		solder.NewPackFromModel(
			record,
			session.Client(c),
			session.Key(c),
			c.Query("include"),
		),
	)
}

// SolderBuild retrieves the build compatible to Technic Platform.
func SolderBuild(c *gin.Context) {
	record, err := store.GetSolderBuild(
		c,
		c.Param("pack"),
		c.Param("build"),
	)

	if err != nil {
		logrus.Warnf("Failed to fetch solder build. %s", err)

		c.JSON(
			http.StatusOK,
			gin.H{
				"error": "Build does not exist",
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		solder.NewBuildFromModel(
			record,
			session.Client(c),
			session.Key(c),
			c.Query("include"),
		),
	)
}
