package api

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/kleister/kleister-api/model"
	"github.com/kleister/kleister-api/model/minecraft"
	"github.com/kleister/kleister-api/store"
)

// MinecraftIndex retrieves all available Minecraft versions.
func MinecraftIndex(c *gin.Context) {
	records, err := store.GetMinecrafts(
		c,
	)

	if err != nil {
		logrus.Warnf("Failed to fetch minecraft versions. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to fetch Minecraft versions",
			},
		)

		c.Abort()
		return
	}

	if c.Param("minecraft") != "" {
		records = records.Filter(
			c.Param("minecraft"),
		)
	}

	c.JSON(
		http.StatusOK,
		records,
	)
}

// MinecraftUpdate updates the list of available Minecraft versions.
func MinecraftUpdate(c *gin.Context) {
	result, err := minecraft.Load()

	if err != nil {
		logrus.Warnf("Failed to sync minecraft versions. %s", err)

		c.JSON(
			http.StatusServiceUnavailable,
			gin.H{
				"status":  http.StatusServiceUnavailable,
				"message": "Failed to request Minecraft versions",
			},
		)

		c.Abort()
		return
	}

	for _, version := range result.Versions {
		if version.Invalid() {
			continue
		}

		_, err := store.SyncMinecraft(
			c,
			version,
		)

		if err != nil {
			logrus.Warnf("Failed to store minecraft version. %s", err)

			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"status":  http.StatusInternalServerError,
					"message": "Failed to store Minecraft versions",
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
			"message": "Successfully imported Minecraft versions",
		},
	)
}

// MinecraftBuildIndex retrieves all builds related to a Minecraft version.
func MinecraftBuildIndex(c *gin.Context) {
	records, err := store.GetMinecraftBuilds(
		c,
		&model.MinecraftBuildParams{
			Minecraft: c.Param("minecraft"),
		},
	)

	if err != nil {
		logrus.Warnf("Failed to fetch minecraft builds. %s", err)

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

// MinecraftBuildAppend appends a build to a Minecraft version.
func MinecraftBuildAppend(c *gin.Context) {
	form := &model.MinecraftBuildParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind minecraft build data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind minecraft build data",
			},
		)

		c.Abort()
		return
	}

	form.Minecraft = c.Param("minecraft")

	assigned := store.GetMinecraftHasBuild(
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

	err := store.CreateMinecraftBuild(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to append minecraft build. %s", err)

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

// MinecraftBuildDelete deleted a build from a Minecraft version
func MinecraftBuildDelete(c *gin.Context) {
	form := &model.MinecraftBuildParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind minecraft build data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind minecraft build data",
			},
		)

		c.Abort()
		return
	}

	form.Minecraft = c.Param("minecraft")

	assigned := store.GetMinecraftHasBuild(
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

	err := store.DeleteMinecraftBuild(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to delete minecraft build. %s", err)

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
