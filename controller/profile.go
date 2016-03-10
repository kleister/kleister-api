package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/solderapp/solder-api.v0/router/middleware/context"
	"gopkg.in/solderapp/solder-api.v0/router/middleware/session"
)

// GetProfile displays the current profile.
func GetProfile(c *gin.Context) {
	record := session.Current(c)

	c.JSON(
		http.StatusOK,
		record,
	)
}

// PatchProfile updates the current profile.
func PatchProfile(c *gin.Context) {
	record := session.Current(c)

	if err := c.BindJSON(&record); err != nil {
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

	err := context.Store(c).Save(
		&record,
	).Error

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
