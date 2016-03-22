package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"
	"github.com/solderapp/solder-api/model/forge"
	"github.com/solderapp/solder-api/router/middleware/context"
	"github.com/solderapp/solder-api/router/middleware/session"
)

// GetForge retrieves all available Forge versions.
func GetForge(c *gin.Context) {
	records, err := context.Store(c).GetForges()

	if err != nil {
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

	c.JSON(
		http.StatusOK,
		records,
	)
}

// CompleteForge returns filtered Forge versions for autocompletion.
func CompleteForge(c *gin.Context) {
	records, err := context.Store(c).GetForges()

	if err != nil {
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

	c.JSON(
		http.StatusOK,
		records.Filter(c.Param("forge")),
	)
}

// PatchForge updates the list of available Forge versions.
func PatchForge(c *gin.Context) {
	result, err := forge.Load()

	if err != nil {
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
		record := &model.Forge{}

		if number.Invalid() {
			continue
		}

		err := context.Store(c).Where(
			model.Forge{
				Name: number.ID,
			},
		).Attrs(
			model.Forge{
				Minecraft: number.Minecraft,
			},
		).FirstOrCreate(
			&record,
		).Error

		if err != nil {
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

// GetForgeBuilds retrieves all builds related to a Forge version.
func GetForgeBuilds(c *gin.Context) {
	forge := session.Forge(c)
	records := &model.Builds{}

	err := context.Store(c).Model(
		&forge,
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

// PatchForgeBuild appends a build to a Forge version.
func PatchForgeBuild(c *gin.Context) {
	forge := session.Forge(c)
	build := session.Build(c)

	count := context.Store(c).Model(
		&forge,
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

	build.Forge = forge

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

// DeleteForgeBuild deleted a build from a Forge version
func DeleteForgeBuild(c *gin.Context) {
	forge := session.Forge(c)
	build := session.Build(c)

	count := context.Store(c).Model(
		&forge,
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

	build.ForgeID = 0

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
