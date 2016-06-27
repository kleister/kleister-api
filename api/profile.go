package api

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/router/middleware/session"
	"github.com/solderapp/solder-api/store"
)

// ProfileShow displays the current profile.
func ProfileShow(c *gin.Context) {
	record := session.Current(c)

	c.JSON(
		http.StatusOK,
		record,
	)
}

// PatchProfile updates the current profile.
func ProfileUpdate(c *gin.Context) {
	record := session.Current(c)

	if err := c.BindJSON(&record); err != nil {
		logrus.Warn("Failed to bind profile data")
		logrus.Warn(err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind profile data",
			},
		)

		c.Abort()
		return
	}

	err := store.UpdateUser(
		c,
		record,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		record,
	)
}
