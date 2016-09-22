package api

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/kleister/kleister-api/model"
	"github.com/kleister/kleister-api/router/middleware/session"
	"github.com/kleister/kleister-api/store"
)

// KeyIndex retrieves all available keys.
func KeyIndex(c *gin.Context) {
	records, err := store.GetKeys(
		c,
	)

	if err != nil {
		logrus.Warnf("Failed to fetch keys. %s", err)

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

// KeyShow retrieves a specific key.
func KeyShow(c *gin.Context) {
	record := session.Key(c)

	c.JSON(
		http.StatusOK,
		record,
	)
}

// KeyDelete removes a specific key.
func KeyDelete(c *gin.Context) {
	record := session.Key(c)

	err := store.DeleteKey(
		c,
		record,
	)

	if err != nil {
		logrus.Warnf("Failed to delete key. %s", err)

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Failed to delete key",
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

// KeyUpdate updates an existing key.
func KeyUpdate(c *gin.Context) {
	record := session.Key(c)

	if err := c.BindJSON(&record); err != nil {
		logrus.Warnf("Failed to bind key data. %s", err)

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

	err := store.UpdateKey(
		c,
		record,
	)

	if err != nil {
		logrus.Warnf("Failed to update key. %s", err)

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

// KeyCreate creates a new key.
func KeyCreate(c *gin.Context) {
	record := &model.Key{}

	if err := c.BindJSON(&record); err != nil {
		logrus.Warnf("Failed to bind key data. %s", err)

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

	err := store.CreateKey(
		c,
		record,
	)

	if err != nil {
		logrus.Warnf("Failed to create key. %s", err)

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

// KeyVerify is a handler to verify a key for Technic.
func KeyVerify(c *gin.Context) {
	record, res := store.GetKeyByValue(
		c,
		c.Param("key"),
	)

	if res.Error != nil || res.RecordNotFound() {
		c.JSON(
			http.StatusOK,
			gin.H{
				"error": "Invalid key provided",
			},
		)
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{
				"valid":      "Valid key provided",
				"name":       record.Name,
				"created_at": record.CreatedAt,
			},
		)
	}
}
