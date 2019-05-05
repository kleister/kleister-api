package pack

// import (
// 	"net/http"

// 	"github.com/Sirupsen/logrus"
// 	"github.com/gin-gonic/gin"
// 	"github.com/kleister/kleister-api/pkg/model"
// 	"github.com/kleister/kleister-api/pkg/router/middleware/session"
// 	"github.com/kleister/kleister-api/pkg/store"
// )

// // PackUpdate updates an existing pack.
// func PackUpdate(c *gin.Context) {
// 	record := session.Pack(c)

// 	if err := c.BindJSON(&record); err != nil {
// 		logrus.Warnf("Failed to bind pack data. %s", err)

// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "Failed to bind pack data",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	err := store.UpdatePack(
// 		c,
// 		record,
// 	)

// 	if err != nil {
// 		logrus.Warnf("Failed to update pack. %s", err)

// 		c.JSON(
// 			http.StatusBadRequest,
// 			gin.H{
// 				"status":  http.StatusBadRequest,
// 				"message": err.Error(),
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	c.JSON(
// 		http.StatusOK,
// 		record,
// 	)
// }
