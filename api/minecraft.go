package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model/minecraft"
	"github.com/solderapp/solder-api/router/middleware/session"
	"github.com/solderapp/solder-api/store"
)

// GetMinecrafts retrieves all available Minecraft versions.
func GetMinecrafts(c *gin.Context) {
	records, err := store.GetMinecrafts(
		c,
	)

	if err != nil {
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

// PatchMinecraft updates the list of available Minecraft versions.
func PatchMinecraft(c *gin.Context) {
	result, err := minecraft.Load()

	if err != nil {
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

// GetMinecraftBuilds retrieves all builds related to a Minecraft version.
func GetMinecraftBuilds(c *gin.Context) {
	minecraft := session.Minecraft(c)

	records, err := store.GetMinecraftBuilds(
		c,
		minecraft.ID,
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

// PatchMinecraftBuild appends a build to a Minecraft version.
func PatchMinecraftBuild(c *gin.Context) {
	minecraft := session.Minecraft(c)
	build := session.Build(c)

	assigned := store.GetMinecraftHasBuild(
		c,
		minecraft.ID,
		build.ID,
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

	err := store.CreateMinecraftBuild(
		c,
		minecraft.ID,
		build.ID,
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

// DeleteMinecraftBuild deleted a build from a Minecraft version
func DeleteMinecraftBuild(c *gin.Context) {
	minecraft := session.Minecraft(c)
	build := session.Build(c)

	assigned := store.GetMinecraftHasBuild(
		c,
		minecraft.ID,
		build.ID,
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

	err := store.DeleteMinecraftBuild(
		c,
		minecraft.ID,
		build.ID,
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
