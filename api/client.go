package api

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/kleister/kleister-api/model"
	"github.com/kleister/kleister-api/router/middleware/session"
	"github.com/kleister/kleister-api/store"
)

// ClientIndex retrieves all available clients.
func ClientIndex(c *gin.Context) {
	records, err := store.GetClients(
		c,
	)

	if err != nil {
		logrus.Warnf("Failed to fetch clients. %s", err)

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

// ClientShow retrieves a specific client.
func ClientShow(c *gin.Context) {
	record := session.Client(c)

	c.JSON(
		http.StatusOK,
		record,
	)
}

// ClientDelete removes a specific client.
func ClientDelete(c *gin.Context) {
	record := session.Client(c)

	err := store.DeleteClient(
		c,
		record,
	)

	if err != nil {
		logrus.Warnf("Failed to delete client. %s", err)

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Failed to delete client",
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

// ClientUpdate updates an existing client.
func ClientUpdate(c *gin.Context) {
	record := session.Client(c)

	if err := c.BindJSON(&record); err != nil {
		logrus.Warnf("Failed to bind client data. %s", err)

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

	err := store.UpdateClient(
		c,
		record,
	)

	if err != nil {
		logrus.Warnf("Failed to update client. %s", err)

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

// ClientCreate creates a new client.
func ClientCreate(c *gin.Context) {
	record := &model.Client{}

	if err := c.BindJSON(&record); err != nil {
		logrus.Warnf("Failed to bind client data. %s", err)

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

	err := store.CreateClient(
		c,
		record,
	)

	if err != nil {
		logrus.Warnf("Failed to create client. %s", err)

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

// ClientPackIndex retrieves all packs related to a client.
func ClientPackIndex(c *gin.Context) {
	records, err := store.GetClientPacks(
		c,
		&model.ClientPackParams{
			Client: c.Param("client"),
		},
	)

	if err != nil {
		logrus.Warnf("Failed to fetch client packs. %s", err)

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

// ClientPackAppend appends a pack to a client.
func ClientPackAppend(c *gin.Context) {
	form := &model.ClientPackParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind client pack data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind client pack data",
			},
		)

		c.Abort()
		return
	}

	form.Client = c.Param("client")

	assigned := store.GetClientHasPack(
		c,
		form,
	)

	if assigned {
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

	err := store.CreateClientPack(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to append client pack. %s", err)

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

// ClientPackDelete deleted a pack from a client
func ClientPackDelete(c *gin.Context) {
	form := &model.ClientPackParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind client pack data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind client pack data",
			},
		)

		c.Abort()
		return
	}

	form.Client = c.Param("client")

	assigned := store.GetClientHasPack(
		c,
		form,
	)

	if !assigned {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Pack is not assigned",
			},
		)

		c.Abort()
		return
	}

	err := store.DeleteClientPack(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to delete client pack. %s", err)

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
