package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model/solder"
	"github.com/solderapp/solder-api/router/middleware/location"
	"github.com/solderapp/solder-api/store"
)

// GetSolderPacks retrieves the packs compatible to Technic Platform.
func GetSolderPacks(c *gin.Context) {
	records, _ := store.GetSolderPacks(
		c,
	)

	result := make(
		map[string]string,
	)

	for _, record := range *records {
		result[record.Slug] = record.Name
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"modpacks": result,
		},
	)
}

// GetSolderPack retrieves the pack compatible to Technic Platform.
func GetSolderPack(c *gin.Context) {
	record, err := store.GetSolderPack(
		c,
		c.Param("pack"),
		location.Location(c).String(),
	)

	if err != nil {
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
		),
	)
}

// GetSolderBuild retrieves the build compatible to Technic Platform.
func GetSolderBuild(c *gin.Context) {
	record, err := store.GetSolderBuild(
		c,
		c.Param("pack"),
		c.Param("build"),
		location.Location(c).String(),
	)

	if err != nil {
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
		),
	)
}

// GetSolderMods retrieves the mods compatible to Technic Platform.
func GetSolderMods(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"error": "No mod requested",
		},
	)
}

// GetSolderMod retrieves the mod compatible to Technic Platform.
func GetSolderMod(c *gin.Context) {
	record, res := store.GetMod(
		c,
		c.Param("mod"),
	)

	if res.RecordNotFound() {
		c.JSON(
			http.StatusOK,
			gin.H{
				"error": "Mod does not exist",
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		solder.NewModFromModel(
			record,
		),
	)
}

// GetSolderVersion retrieves the version compatible to Technic Platform.
func GetSolderVersion(c *gin.Context) {
	location := location.Location(c)

	parent, res := store.GetMod(
		c,
		c.Param("mod"),
	)

	if res.RecordNotFound() {
		c.JSON(
			http.StatusOK,
			gin.H{
				"error": "Mod does not exist",
			},
		)

		return
	}

	record, res := store.GetVersion(
		c,
		parent.ID,
		c.Param("version"),
	)

	if res.RecordNotFound() {
		c.JSON(
			http.StatusOK,
			gin.H{
				"error": "Version does not exist",
			},
		)

		return
	}

	if record.File != nil {
		record.File.URL = strings.Join(
			[]string{
				location.String(),
				"storage",
				"version",
				strconv.Itoa(record.ID),
			},
			"/",
		)
	}

	c.JSON(
		http.StatusOK,
		solder.NewVersionFromModel(
			record,
		),
	)
}
