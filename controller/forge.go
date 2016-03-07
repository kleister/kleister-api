package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder/model"
	"github.com/solderapp/solder/model/forge"
	"github.com/solderapp/solder/router/middleware/context"
)

// GetForge retrieves all available Forge versions.
func GetForge(c *gin.Context) {
	records := &model.Forges{}

	err := context.Store(c).Find(
		&records,
	).Error

	if err != nil {
		c.IndentedJSON(
			500,
			gin.H{
				"status":  500,
				"message": "Failed to fetch Forge versions",
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

// CompleteForge returns filtered Forge versions for autocompletion.
func CompleteForge(c *gin.Context) {
	records := &model.Forges{}

	err := context.Store(c).Where(
		"name LIKE ?",
		fmt.Sprintf("%%%s%%", c.Param("filter")),
	).Find(
		&records,
	).Error

	if err != nil {
		c.IndentedJSON(
			500,
			gin.H{
				"status":  500,
				"message": "Failed to filter Forge versions",
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

// PatchForge updates the list of available Forge versions.
func PatchForge(c *gin.Context) {
	result, err := forge.Load()

	if err != nil {
		c.IndentedJSON(
			422,
			gin.H{
				"status":  422,
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
			c.IndentedJSON(
				422,
				gin.H{
					"status":  422,
					"message": "Failed to store Forge versions",
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
			"message": "Successfully imported Forge versions",
		},
	)
}
