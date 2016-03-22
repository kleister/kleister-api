package controller

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"
	"github.com/solderapp/solder-api/router/middleware/context"
	"github.com/solderapp/solder-api/router/middleware/session"
	"github.com/vincent-petithory/dataurl"
)

// GetVersions retrieves all available versions.
func GetVersions(c *gin.Context) {
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

	c.JSON(
		http.StatusOK,
		records,
	)
}

// GetVersion retrieves a specific version.
func GetVersion(c *gin.Context) {
	record := session.Version(c)

	// c.Request.Host

	c.JSON(
		http.StatusOK,
		record,
	)
}

// GetVersionFile retrieves a file for a specific version.
func GetVersionFile(c *gin.Context) {
	record := session.Version(c)

	if record.File == nil {
		c.AbortWithError(
			http.StatusNotFound,
			fmt.Errorf("No file content available"),
		)

		return
	}

	decoded, err := dataurl.DecodeString(
		record.File.Content,
	)

	if err != nil {
		c.AbortWithError(
			http.StatusInternalServerError,
			fmt.Errorf("Failed to decode file"),
		)

		return
	}

	c.Writer.Header().Set(
		"Content-Type",
		decoded.ContentType(),
	)

	c.Writer.Write(
		decoded.Data,
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
