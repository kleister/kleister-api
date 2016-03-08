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
	record := &model.Client{}

	res := context.Store(c).Where(
		"clients.id = ?",
		c.Param("client"),
	).Or(
		"clients.slug = ?",
		c.Param("client"),
	).First(
		&record,
	)

	if res.RecordNotFound() {
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
	record := &model.Client{}

	res := context.Store(c).Where(
		"clients.id = ?",
		c.Param("client"),
	).Or(
		"clients.slug = ?",
		c.Param("client"),
	).First(
		&record,
	)

	if res.RecordNotFound() {
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

	err := context.Store(c).Delete(
		&record,
	).Error

	if err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			},
		)

		c.Abort()
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully deleted client",
		},
	)
}

// PatchClient updates an existing client.
func PatchClient(c *gin.Context) {
	record := &model.Client{}

	res := context.Store(c).Where(
		"clients.id = ?",
		c.Param("client"),
	).Or(
		"clients.slug = ?",
		c.Param("client"),
	).First(
		&record,
	)

	if res.RecordNotFound() {
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

	if err := c.BindJSON(&record); err != nil {
		c.IndentedJSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind client data",
			},
		)

		c.Abort()
		return
	}

	err := context.Store(c).Save(
		&record,
	).Error

	if err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
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

// PostClient creates a new client.
func PostClient(c *gin.Context) {
	record := &model.Client{}
	record.Defaults()

	if err := c.BindJSON(&record); err != nil {
		c.IndentedJSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind client data",
			},
		)

		c.Abort()
		return
	}

	err := context.Store(c).Create(
		&record,
	).Error

	if err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
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
