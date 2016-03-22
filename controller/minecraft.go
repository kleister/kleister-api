package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"
	"github.com/solderapp/solder-api/model/minecraft"
	"github.com/solderapp/solder-api/router/middleware/context"
	"github.com/solderapp/solder-api/router/middleware/session"
)

// GetMinecraft retrieves all available Minecraft versions.
func GetMinecraft(c *gin.Context) {
	records, err := context.Store(c).GetMinecrafts()

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

	c.JSON(
		http.StatusOK,
		records,
	)
}

// CompleteMinecraft returns filtered Minecraft versions for autocompletion.
func CompleteMinecraft(c *gin.Context) {
	records, err := context.Store(c).GetMinecrafts()

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

	c.JSON(
		http.StatusOK,
		records.Filter(c.Param("minecraft")),
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
		record := &model.Minecraft{}

		if version.Invalid() {
			continue
		}

		err := context.Store(c).Where(
			model.Minecraft{
				Name: version.ID,
			},
		).Attrs(
			model.Minecraft{
				Type: version.Type,
			},
		).FirstOrCreate(
			&record,
		).Error

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
	records := &model.Builds{}

	err := context.Store(c).Model(
		&minecraft,
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

// PatchMinecraftBuild appends a build to a Minecraft version.
func PatchMinecraftBuild(c *gin.Context) {
	minecraft := session.Minecraft(c)
	build := session.Build(c)

	count := context.Store(c).Model(
		&minecraft,
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

	build.Minecraft = minecraft

	err := context.Store(c).Save(
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

// DeleteMinecraftBuild deleted a build from a Minecraft version
func DeleteMinecraftBuild(c *gin.Context) {
	minecraft := session.Minecraft(c)
	build := session.Build(c)

	count := context.Store(c).Model(
		&minecraft,
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

	build.MinecraftID = 0

	err := context.Store(c).Save(
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
