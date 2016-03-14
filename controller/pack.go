package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"
	"github.com/solderapp/solder-api/router/middleware/context"
	"github.com/solderapp/solder-api/router/middleware/session"
)

// GetPacks retrieves all available packs.
func GetPacks(c *gin.Context) {
	records := &model.Packs{}

	err := context.Store(c).Scopes(
		model.PackDefaultOrder,
	).Preload(
		"Builds",
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

// GetPack retrieves a specific pack.
func GetPack(c *gin.Context) {
	record := session.Pack(c)

	c.JSON(
		http.StatusOK,
		record,
	)
}

// DeletePack removes a specific pack.
func DeletePack(c *gin.Context) {
	record := session.Pack(c)

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
			"message": "Successfully deleted pack",
		},
	)
}

// PatchPack updates an existing pack.
func PatchPack(c *gin.Context) {
	record := session.Pack(c)

	if err := c.BindJSON(&record); err != nil {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind pack data",
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

// PostPack creates a new pack.
func PostPack(c *gin.Context) {
	record := &model.Pack{}
	record.Defaults()

	if err := c.BindJSON(&record); err != nil {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind pack data",
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

// GetPackClients retrieves all clients related to a pack.
func GetPackClients(c *gin.Context) {
	pack := session.Pack(c)
	records := &model.Clients{}

	err := context.Store(c).Model(
		&pack,
	).Association(
		"Clients",
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

// PatchPackClient appends a client to a pack.
func PatchPackClient(c *gin.Context) {
	pack := session.Pack(c)
	client := session.Client(c)

	count := context.Store(c).Model(
		&pack,
	).Association(
		"Clients",
	).Find(
		&client,
	).Count()

	if count > 0 {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Client is already appended",
			},
		)

		c.Abort()
		return
	}

	err := context.Store(c).Model(
		&pack,
	).Association(
		"Clients",
	).Append(
		&client,
	).Error

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to append client",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully appended client",
		},
	)
}

// DeletePackClient deleted a client from a pack
func DeletePackClient(c *gin.Context) {
	pack := session.Pack(c)
	client := session.Client(c)

	count := context.Store(c).Model(
		&pack,
	).Association(
		"Clients",
	).Find(
		&client,
	).Count()

	if count < 1 {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "Client is not assigned",
			},
		)

		c.Abort()
		return
	}

	err := context.Store(c).Model(
		&pack,
	).Association(
		"Clients",
	).Delete(
		&client,
	).Error

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to unlink client",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully unlinked client",
		},
	)
}
