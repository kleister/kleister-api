package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"
	"github.com/solderapp/solder-api/router/middleware/location"
	"github.com/solderapp/solder-api/router/middleware/session"
	"github.com/solderapp/solder-api/store"
)

// GetVersions retrieves all available versions.
func GetVersions(c *gin.Context) {
	location := location.Location(c)
	mod := session.Mod(c)

	records, err := store.GetVersions(
		c,
		mod.ID,
	)

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

	for _, record := range *records {
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
	}

	c.JSON(
		http.StatusOK,
		records,
	)
}

// GetVersion retrieves a specific version.
func GetVersion(c *gin.Context) {
	location := location.Location(c)
	record := session.Version(c)

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
		record,
	)
}

// DeleteVersion removes a specific version.
func DeleteVersion(c *gin.Context) {
	mod := session.Mod(c)
	record := session.Version(c)

	err := store.DeleteVersion(
		c,
		mod.ID,
		record,
	)

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
	mod := session.Mod(c)
	record := session.Version(c)

	if err := c.BindJSON(&record); err != nil {
		logrus.Warn("Failed to bind version data")
		logrus.Warn(err)

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

	err := store.UpdateVersion(
		c,
		mod.ID,
		record,
	)

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
	record := &model.Version{}

	if err := c.BindJSON(&record); err != nil {
		logrus.Warn("Failed to bind version data")
		logrus.Warn(err)

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

	err := store.CreateVersion(
		c,
		mod.ID,
		record,
	)

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
	records, err := store.GetVersionBuilds(
		c,
		&model.VersionBuildParams{
			Mod:     c.Param("mod"),
			Version: c.Param("version"),
		},
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to fetch builds",
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

// PatchVersionBuild appends a build to a version.
func PatchVersionBuild(c *gin.Context) {
	form := &model.VersionBuildParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warn("Failed to bind post data")
		logrus.Warn(err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind form data",
			},
		)

		c.Abort()
		return
	}

	form.Mod = c.Param("mod")
	form.Version = c.Param("version")

	assigned := store.GetVersionHasBuild(
		c,
		form,
	)

	if assigned == true {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Build is already appended",
			},
		)

		c.Abort()
		return
	}

	err := store.CreateVersionBuild(
		c,
		form,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to append build",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully appended build",
		},
	)
}

// DeleteVersionBuild deleted a build from a version
func DeleteVersionBuild(c *gin.Context) {
	form := &model.VersionBuildParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warn("Failed to bind post data")
		logrus.Warn(err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind form data",
			},
		)

		c.Abort()
		return
	}

	form.Mod = c.Param("mod")
	form.Version = c.Param("version")

	assigned := store.GetVersionHasBuild(
		c,
		form,
	)

	if assigned == false {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Build is not assigned",
			},
		)

		c.Abort()
		return
	}

	err := store.DeleteVersionBuild(
		c,
		form,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to unlink build",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully unlinked build",
		},
	)
}
