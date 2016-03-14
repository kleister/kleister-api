package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"
	"github.com/solderapp/solder-api/router/middleware/context"
	"github.com/solderapp/solder-api/router/middleware/session"
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
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to fetch clients",
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

// GetClient retrieves a specific client.
func GetClient(c *gin.Context) {
	record := session.Client(c)

	c.JSON(
		http.StatusOK,
		record,
	)
}

// DeleteClient removes a specific client.
func DeleteClient(c *gin.Context) {
	record := session.Client(c)

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
			"message": "Successfully deleted client",
		},
	)
}

// PatchClient updates an existing client.
func PatchClient(c *gin.Context) {
	record := session.Client(c)

	if err := c.BindJSON(&record); err != nil {
		c.JSON(
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

// PostClient creates a new client.
func PostClient(c *gin.Context) {
	record := &model.Client{}
	record.Defaults()

	if err := c.BindJSON(&record); err != nil {
		c.JSON(
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

// GetClientPacks retrieves all packs related to a client.
func GetClientPacks(c *gin.Context) {
	client := session.Client(c)
	records := &model.Packs{}

	err := context.Store(c).Model(
		&client,
	).Association(
		"Packs",
	).Find(
		&records,
	).Error

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to fetch packs",
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

// PatchClientPack appends a pack to a client.
func PatchClientPack(c *gin.Context) {
	client := session.Client(c)
	pack := session.Pack(c)

	count := context.Store(c).Model(
		&client,
	).Association(
		"Packs",
	).Find(
		&pack,
	).Count()

	if count > 0 {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Pack is already appended",
			},
		)

		c.Abort()
		return
	}

	err := context.Store(c).Model(
		&client,
	).Association(
		"Packs",
	).Append(
		&pack,
	).Error

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to append pack",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully appended pack",
		},
	)
}

// DeleteClientPack deleted a pack from a client
func DeleteClientPack(c *gin.Context) {
	client := session.Client(c)
	pack := session.Pack(c)

	count := context.Store(c).Model(
		&client,
	).Association(
		"Packs",
	).Find(
		&pack,
	).Count()

	if count < 1 {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "Pack is not assigned",
			},
		)

		c.Abort()
		return
	}

	err := context.Store(c).Model(
		&client,
	).Association(
		"Packs",
	).Delete(
		&pack,
	).Error

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to unlink pack",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully unlinked pack",
		},
	)
}
