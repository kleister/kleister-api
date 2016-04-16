package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"
	"github.com/solderapp/solder-api/router/middleware/context"
	"github.com/solderapp/solder-api/router/middleware/session"
)

// GetVersions retrieves all available versions.
func GetVersions(c *gin.Context) {
	location := context.Location(c)
	mod := session.Mod(c)

	records, err := context.Store(c).GetVersions(
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
	location := context.Location(c)
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
	config := context.Config(c)
	record := session.Version(c)

	tx := context.Store(c).Begin()
	failed := false

	if record.File != nil {
		record.File.Path = config.Server.Storage

		err := tx.Delete(
			&record.File,
		).Error

		if err != nil {
			failed = true
		}
	}

	err := tx.Delete(
		&record,
	).Error

	if failed || err != nil {
		tx.Rollback()

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Failed to delete version",
			},
		)

		c.Abort()
		return
	}

	tx.Commit()

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
	config := context.Config(c)
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

	record.ModID = mod.ID

	if record.File != nil {
		record.File.Path = config.Server.Storage
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
	config := context.Config(c)
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

	record.ModID = mod.ID

	if record.File != nil {
		record.File.Path = config.Server.Storage
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
	version := session.Version(c)
	records := &model.Builds{}

	err := context.Store(c).Model(
		&version,
	).Association(
		"Builds",
	).Find(
		&records,
	).Error

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
	version := session.Version(c)
	build := session.Build(c)

	count := context.Store(c).Model(
		&version,
	).Association(
		"Builds",
	).Find(
		&build,
	).Count()

	if count > 0 {
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

	err := context.Store(c).Model(
		&version,
	).Association(
		"Builds",
	).Append(
		&build,
	).Error

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
	version := session.Version(c)
	build := session.Build(c)

	count := context.Store(c).Model(
		&version,
	).Association(
		"Builds",
	).Find(
		&build,
	).Count()

	if count < 1 {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "Build is not assigned",
			},
		)

		c.Abort()
		return
	}

	err := context.Store(c).Model(
		&version,
	).Association(
		"Builds",
	).Delete(
		&build,
	).Error

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
