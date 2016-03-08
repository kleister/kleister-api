package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder/model"
	"github.com/solderapp/solder/model/minecraft"
	"github.com/solderapp/solder/router/middleware/context"
)

// GetMinecraft retrieves all available Minecraft versions.
func GetMinecraft(c *gin.Context) {
	records := &model.Minecrafts{}

	err := context.Store(c).Scopes(
		model.MinecraftDefaultOrder,
	).Find(
		&records,
	).Error

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
	records := &model.Minecrafts{}

	err := context.Store(c).Where(
		"name LIKE ?",
		fmt.Sprintf("%%%s%%", c.Param("filter")),
	).Scopes(
		model.MinecraftDefaultOrder,
	).Find(
		&records,
	).Error

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to filter Minecraft versions",
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
