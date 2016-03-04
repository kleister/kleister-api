package controller

import (
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder/model"
	"github.com/solderapp/solder/model/minecraft"
	"github.com/solderapp/solder/router/middleware/context"
)

func GetMinecraft(c *gin.Context) {
	records := &model.Minecrafts{}

	context.Store(c).Order(
		"name DESC",
	).Find(
		&records,
	)

	c.IndentedJSON(
		200,
		records,
	)
}

func PatchMinecraft(c *gin.Context) {
	result, err := minecraft.Load()

	if err != nil {
		logrus.Warn(err)

		c.IndentedJSON(
			422,
			gin.H{
				"status":  422,
				"message": "Failed to request Minecraft versions",
			},
		)

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
