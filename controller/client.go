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
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
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
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}
