package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder/model"
	"github.com/solderapp/solder/model/forge"
	"github.com/solderapp/solder/router/middleware/context"
)

// GetForge retrieves all available Forge versions.
func GetForge(c *gin.Context) {
	records := &model.Forges{}

	err := context.Store(c).Scopes(
		model.ForgeDefaultOrder,
	).Find(
		&records,
	).Error

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
	records := &model.Forges{}

	err := context.Store(c).Where(
		"name LIKE ?",
		fmt.Sprintf("%%%s%%", c.Param("filter")),
	).Scopes(
		model.ForgeDefaultOrder,
	).Find(
		&records,
	).Error

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to filter Forge versions",
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
