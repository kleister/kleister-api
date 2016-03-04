package controller

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder/model"
	"github.com/solderapp/solder/model/forge"
	"github.com/solderapp/solder/router/middleware/context"
)

func GetForge(c *gin.Context) {
	records := &model.Forges{}

	context.Store(c).Order(
		"minecraft DESC, name DESC",
	).Find(
		&records,
	)

	c.IndentedJSON(
		200,
		records,
	)
}

func CompleteForge(c *gin.Context) {
	records := &model.Forges{}

	context.Store(c).Where(
		"name LIKE ?",
		fmt.Sprintf("%%%s%%", c.Param("filter")),
	).Order(
		"minecraft DESC, name DESC",
	).Find(
		&records,
	)

	c.IndentedJSON(
		200,
		records,
	)
}

func PatchForge(c *gin.Context) {
	result, err := forge.Load()

	if err != nil {
		logrus.Warn(err)

		c.IndentedJSON(
			422,
			gin.H{
				"status":  422,
				"message": "Failed to request Forge versions",
			},
		)

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
