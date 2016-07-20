package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/kleister/kleister-api/model"
	"github.com/kleister/kleister-api/router/middleware/location"
	"github.com/kleister/kleister-api/router/middleware/session"
	"github.com/kleister/kleister-api/store"
)

// VersionIndex retrieves all available versions.
func VersionIndex(c *gin.Context) {
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

// VersionShow retrieves a specific version.
func VersionShow(c *gin.Context) {
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

// VersionDelete removes a specific version.
func VersionDelete(c *gin.Context) {
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

// VersionUpdate updates an existing version.
func VersionUpdate(c *gin.Context) {
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

// VersionCreate creates a new version.
func VersionCreate(c *gin.Context) {
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

// VersionBuildIndex retrieves all builds related to a version.
func VersionBuildIndex(c *gin.Context) {
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

// VersionBuildAppend appends a build to a version.
func VersionBuildAppend(c *gin.Context) {
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

// VersionBuildDelete deleted a build from a version
func VersionBuildDelete(c *gin.Context) {
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
