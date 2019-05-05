package client

// import (
// 	"net/http"

// 	"github.com/Sirupsen/logrus"
// 	"github.com/gin-gonic/gin"
// 	"github.com/kleister/kleister-api/pkg/model"
// 	"github.com/kleister/kleister-api/pkg/router/middleware/session"
// 	"github.com/kleister/kleister-api/pkg/store"
// )

// // ClientPackIndex retrieves all packs related to a client.
// func ClientPackIndex(c *gin.Context) {
// 	records, err := store.GetClientPacks(
// 		c,
// 		&model.ClientPackParams{
// 			Client: c.Param("client"),
// 		},
// 	)

// 	if err != nil {
// 		logrus.Warnf("Failed to fetch client packs. %s", err)

// 		c.JSON(
// 			http.StatusInternalServerError,
// 			gin.H{
// 				"status":  http.StatusInternalServerError,
// 				"message": "Failed to fetch packs",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	c.JSON(
// 		http.StatusOK,
// 		records,
// 	)
// }

// // ClientPackAppend appends a pack to a client.
// func ClientPackAppend(c *gin.Context) {
// 	form := &model.ClientPackParams{}

// 	if err := c.BindJSON(&form); err != nil {
// 		logrus.Warnf("Failed to bind client pack data. %s", err)

// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "Failed to bind client pack data",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	form.Client = c.Param("client")

// 	assigned := store.GetClientHasPack(
// 		c,
// 		form,
// 	)

// 	if assigned {
// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "Pack is already appended",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	err := store.CreateClientPack(
// 		c,
// 		form,
// 	)

// 	if err != nil {
// 		logrus.Warnf("Failed to append client pack. %s", err)

// 		c.JSON(
// 			http.StatusInternalServerError,
// 			gin.H{
// 				"status":  http.StatusInternalServerError,
// 				"message": "Failed to append pack",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	c.JSON(
// 		http.StatusOK,
// 		gin.H{
// 			"status":  http.StatusOK,
// 			"message": "Successfully appended pack",
// 		},
// 	)
// }

// // ClientPackDelete deleted a pack from a client
// func ClientPackDelete(c *gin.Context) {
// 	form := &model.ClientPackParams{}

// 	if err := c.BindJSON(&form); err != nil {
// 		logrus.Warnf("Failed to bind client pack data. %s", err)

// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "Failed to bind client pack data",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	form.Client = c.Param("client")

// 	assigned := store.GetClientHasPack(
// 		c,
// 		form,
// 	)

// 	if !assigned {
// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "Pack is not assigned",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	err := store.DeleteClientPack(
// 		c,
// 		form,
// 	)

// 	if err != nil {
// 		logrus.Warnf("Failed to delete client pack. %s", err)

// 		c.JSON(
// 			http.StatusInternalServerError,
// 			gin.H{
// 				"status":  http.StatusInternalServerError,
// 				"message": "Failed to unlink pack",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	c.JSON(
// 		http.StatusOK,
// 		gin.H{
// 			"status":  http.StatusOK,
// 			"message": "Successfully unlinked pack",
// 		},
// 	)
// }
