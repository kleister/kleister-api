package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder/model"
	"github.com/solderapp/solder/router/middleware/context"
)

// GetKeys retrieves all available keys.
func GetKeys(c *gin.Context) {
	records := &model.Keys{}

	err := context.Store(c).Scopes(
		model.KeyDefaultOrder,
	).Find(
		&records,
	).Error

	if err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to fetch keys",
			},
		)

		c.Abort()
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		records,
	)
}

// GetKey retrieves a specific key.
func GetKey(c *gin.Context) {
	record := &model.Key{
		Slug: c.Param("key"),
	}

	err := context.Store(c).Find(
		&record,
	).Error

	if err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "Failed to find key",
			},
		)

		c.Abort()
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		record,
	)
}

// DeleteKey removes a specific key.
func DeleteKey(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// PatchKey updates an existing key.
func PatchKey(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// PostKey creates a new key.
func PostKey(c *gin.Context) {
	record := &model.Client{}
	record.Defaults()

	if err := c.Bind(&record); err != nil {
		c.IndentedJSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind key form data",
			},
		)

		c.Abort()
		return
	}

	res := context.Store(c).Create(
		&record,
	)

	if res.Error != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": res.Error.Error(),
			},
		)

		c.Abort()
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		record,
	)
}
