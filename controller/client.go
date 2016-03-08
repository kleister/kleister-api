package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder/model"
	"github.com/solderapp/solder/router/middleware/context"
)

// GetClients retrieves all available clients.
func GetClients(c *gin.Context) {
	records := &model.Clients{}

	err := context.Store(c).Scopes(
		model.ClientDefaultOrder,
	).Find(
		&records,
	).Error

	if err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to fetch clients",
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

// GetClient retrieves a specific client.
func GetClient(c *gin.Context) {
	record := &model.Client{
		Slug: c.Param("client"),
	}

	err := context.Store(c).Find(
		&record,
	).Error

	if err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "Failed to find client",
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

// DeleteClient removes a specific client.
func DeleteClient(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// PatchClient creates a new client.
func PatchClient(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// PostClient updates an existing client.
func PostClient(c *gin.Context) {
	record := &model.Client{}
	record.Defaults()

	if err := c.BindJSON(&record); err != nil {
		c.IndentedJSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind client form data",
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
