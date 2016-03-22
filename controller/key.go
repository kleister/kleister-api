package controller

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"
	"github.com/solderapp/solder-api/router/middleware/context"
	"github.com/solderapp/solder-api/router/middleware/session"
)

// GetKeys retrieves all available keys.
func GetKeys(c *gin.Context) {
	records, err := context.Store(c).GetKeys()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to fetch keys",
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

// GetKey retrieves a specific key.
func GetKey(c *gin.Context) {
	record := session.Key(c)

	c.JSON(
		http.StatusOK,
		record,
	)
}

// DeleteKey removes a specific key.
func DeleteKey(c *gin.Context) {
	record := session.Key(c)

	err := context.Store(c).Delete(
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
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully deleted key",
		},
	)
}

// PatchKey updates an existing key.
func PatchKey(c *gin.Context) {
	record := session.Key(c)

	if err := c.Bind(&record); err != nil {
		logrus.Warn("Failed to bind key data")
		logrus.Warn(err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind key data",
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

// PostKey creates a new key.
func PostKey(c *gin.Context) {
	record := &model.Key{}

	if err := c.Bind(&record); err != nil {
		logrus.Warn("Failed to bind key data")
		logrus.Warn(err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind key data",
			},
		)

		c.Abort()
		return
	}

	err := context.Store(c).Create(
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
