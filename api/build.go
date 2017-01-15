package api

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/kleister/kleister-api/model"
	"github.com/kleister/kleister-api/router/middleware/session"
	"github.com/kleister/kleister-api/store"
)

// BuildIndex retrieves all available builds.
func BuildIndex(c *gin.Context) {
	pack := session.Pack(c)

	records, err := store.GetBuilds(
		c,
		pack.ID,
	)

	if err != nil {
		logrus.Warnf("Failed to fetch builds. %s", err)

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

// BuildShow retrieves a specific build.
func BuildShow(c *gin.Context) {
	record := session.Build(c)

	c.JSON(
		http.StatusOK,
		record,
	)
}

// BuildDelete removes a specific build.
func BuildDelete(c *gin.Context) {
	pack := session.Pack(c)
	record := session.Build(c)

	err := store.DeleteBuild(
		c,
		pack.ID,
		record,
	)

	if err != nil {
		logrus.Warnf("Failed to delete build. %s", err)

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Failed to delete build",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully deleted build",
		},
	)
}

// BuildUpdate updates an existing build.
func BuildUpdate(c *gin.Context) {
	pack := session.Pack(c)
	record := session.Build(c)

	if err := c.BindJSON(&record); err != nil {
		logrus.Warnf("Failed to bind build data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind build data",
			},
		)

		c.Abort()
		return
	}

	err := store.UpdateBuild(
		c,
		pack.ID,
		record,
	)

	if err != nil {
		logrus.Warnf("Failed to update build. %s", err)

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

// BuildCreate creates a new build.
func BuildCreate(c *gin.Context) {
	pack := session.Pack(c)
	record := &model.Build{}

	if err := c.BindJSON(&record); err != nil {
		logrus.Warnf("Failed to bind build data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind build data",
			},
		)

		c.Abort()
		return
	}

	err := store.CreateBuild(
		c,
		pack.ID,
		record,
	)

	if err != nil {
		logrus.Warnf("Failed to create build. %s", err)

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

// BuildVersionIndex retrieves all versions related to a build.
func BuildVersionIndex(c *gin.Context) {
	records, err := store.GetBuildVersions(
		c,
		&model.BuildVersionParams{
			Pack:  c.Param("pack"),
			Build: c.Param("build"),
		},
	)

	if err != nil {
		logrus.Warnf("Failed to fetch build versions. %s", err)

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

// BuildVersionAppend appends a version to a build.
func BuildVersionAppend(c *gin.Context) {
	form := &model.BuildVersionParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind build version data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind build version data",
			},
		)

		c.Abort()
		return
	}

	form.Pack = c.Param("pack")
	form.Build = c.Param("build")

	assigned := store.GetBuildHasVersion(
		c,
		form,
	)

	if assigned {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Version is already appended",
			},
		)

		c.Abort()
		return
	}

	err := store.CreateBuildVersion(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to append build version. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to append version",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully appended version",
		},
	)
}

// BuildVersionDelete deleted a version from a build
func BuildVersionDelete(c *gin.Context) {
	form := &model.BuildVersionParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind build version data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind build version data",
			},
		)

		c.Abort()
		return
	}

	form.Pack = c.Param("pack")
	form.Build = c.Param("build")

	assigned := store.GetBuildHasVersion(
		c,
		form,
	)

	if !assigned {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Version is not assigned",
			},
		)

		c.Abort()
		return
	}

	err := store.DeleteBuildVersion(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to delete build version. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to unlink version",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully unlinked version",
		},
	)
}
