package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder/model"
	"github.com/solderapp/solder/model/minecraft"
	"github.com/solderapp/solder/router/middleware/context"
)

// GetMinecraft retrieves all available Minecraft versions.
func GetMinecraft(c *gin.Context) {
	records := &model.Minecrafts{}

	err := context.Store(c).Find(
		&records,
	).Error

	if err != nil {
		c.IndentedJSON(
			422,
			gin.H{
				"status":  422,
				"message": "Failed to fetch Minecraft versions",
			},
		)

		c.Abort()
		return
	}

	c.IndentedJSON(
		200,
		records,
	)
}

// CompleteMinecraft returns filtered Minecraft versions for autocompletion.
func CompleteMinecraft(c *gin.Context) {
	records := &model.Minecrafts{}

	err := context.Store(c).Where(
		"name LIKE ?",
		fmt.Sprintf("%%%s%%", c.Param("filter")),
	).Find(
		&records,
	).Error

	if err != nil {
		c.IndentedJSON(
			422,
			gin.H{
				"status":  422,
				"message": "Failed to filter Minecraft versions",
			},
		)

		c.Abort()
		return
	}

	c.IndentedJSON(
		200,
		records,
	)
}

// PatchMinecraft updates the list of available Minecraft versions.
func PatchMinecraft(c *gin.Context) {
	result, err := minecraft.Load()

	if err != nil {
		c.IndentedJSON(
			422,
			gin.H{
				"status":  422,
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
			c.IndentedJSON(
				422,
				gin.H{
					"status":  422,
					"message": "Failed to store Minecraft versions",
				},
			)

			c.Abort()
			return
		}
	}

	c.IndentedJSON(
		200,
		gin.H{
			"status":  200,
			"message": "Successfully imported Minecraft versions",
		},
	)
}
