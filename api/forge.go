package api

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/kleister/kleister-api/model"
	"github.com/kleister/kleister-api/model/forge"
	"github.com/kleister/kleister-api/store"
)

// ForgeIndex retrieves all available Forge versions.
func ForgeIndex(c *gin.Context) {
	records, err := store.GetForges(
		c,
	)

	if err != nil {
		logrus.Warnf("Failed to fetch forge versions. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to fetch Forge versions",
			},
		)

		c.Abort()
		return
	}

	if c.Param("forge") != "" {
		records = records.Filter(
			c.Param("forge"),
		)
	}

	c.JSON(
		http.StatusOK,
		records,
	)
}

// ForgeUpdate updates the list of available Forge versions.
func ForgeUpdate(c *gin.Context) {
	result, err := forge.Load()

	if err != nil {
		logrus.Warnf("Failed to sync forge versions. %s", err)

		c.JSON(
			http.StatusServiceUnavailable,
			gin.H{
				"status":  http.StatusServiceUnavailable,
				"message": "Failed to request Forge versions",
			},
		)

		c.Abort()
		return
	}

	for _, number := range result.Numbers {
		if number.Invalid() {
			continue
		}

		_, err := store.SyncForge(
			c,
			number,
		)

		if err != nil {
			logrus.Warnf("Failed to store forge version. %s", err)

			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"status":  http.StatusInternalServerError,
					"message": "Failed to store Forge versions",
				},
			)

			c.Abort()
			return
		}
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully imported Forge versions",
		},
	)
}

// ForgeBuildIndex retrieves all builds related to a Forge version.
func ForgeBuildIndex(c *gin.Context) {
	records, err := store.GetForgeBuilds(
		c,
		&model.ForgeBuildParams{
			Forge: c.Param("forge"),
		},
	)

	if err != nil {
		logrus.Warnf("Failed to fetch forge builds. %s", err)

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

// ForgeBuildAppend appends a build to a Forge version.
func ForgeBuildAppend(c *gin.Context) {
	form := &model.ForgeBuildParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind forge build data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind forge build data",
			},
		)

		c.Abort()
		return
	}

	form.Forge = c.Param("forge")

	assigned := store.GetForgeHasBuild(
		c,
		form,
	)

	if assigned {
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

	err := store.CreateForgeBuild(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to append forge build. %s", err)

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

// ForgeBuildDelete deleted a build from a Forge version
func ForgeBuildDelete(c *gin.Context) {
	form := &model.ForgeBuildParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind forge build data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind forge build data",
			},
		)

		c.Abort()
		return
	}

	form.Forge = c.Param("forge")

	assigned := store.GetForgeHasBuild(
		c,
		form,
	)

	if !assigned {
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

	err := store.DeleteForgeBuild(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to delete forge build. %s", err)

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
