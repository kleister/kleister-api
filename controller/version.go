package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"
	"github.com/solderapp/solder-api/router/middleware/context"
	"github.com/solderapp/solder-api/router/middleware/session"
)

// GetVersions retrieves all available versions.
func GetVersions(c *gin.Context) {
	mod := session.Mod(c)
	records := &model.Versions{}

	err := context.Store(c).Scopes(
		model.VersionDefaultOrder,
	).Where(
		"versions.mod_id = ?",
		mod.ID,
	).Preload(
		"Mod",
	).Find(
		&records,
	).Error

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to fetch versions",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		records,
	)
}

// GetVersion retrieves a specific version.
func GetVersion(c *gin.Context) {
	record := session.Version(c)

	c.JSON(
		http.StatusOK,
		record,
	)
}

// DeleteVersion removes a specific version.
func DeleteVersion(c *gin.Context) {
	record := session.Version(c)

	err := context.Store(c).Delete(
		&record,
	).Error

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully deleted version",
		},
	)
}

// PatchVersion updates an existing version.
func PatchVersion(c *gin.Context) {
	record := session.Version(c)

	if err := c.BindJSON(&record); err != nil {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind version data",
			},
		)

		c.Abort()
		return
	}

	err := context.Store(c).Save(
		&record,
	).Error

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		record,
	)
}

// PostVersion creates a new version.
func PostVersion(c *gin.Context) {
	mod := session.Mod(c)

	record := &model.Version{
		ModID: mod.ID,
	}

	record.Defaults()

	if err := c.BindJSON(&record); err != nil {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind version data",
			},
		)

		c.Abort()
		return
	}

	err := context.Store(c).Create(
		&record,
	).Error

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		record,
	)
}

// GetVersionBuilds retrieves all builds related to a version.
func GetVersionBuilds(c *gin.Context) {

}

// PatchVersionBuild appends a build to a version.
func PatchVersionBuild(c *gin.Context) {

}

// DeleteVersionBuild deleted a build from a version
func DeleteVersionBuild(c *gin.Context) {

}
