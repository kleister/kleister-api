package key

// import (
// 	"net/http"

// 	"github.com/Sirupsen/logrus"
// 	"github.com/gin-gonic/gin"
// 	"github.com/kleister/kleister-api/pkg/model"
// 	"github.com/kleister/kleister-api/pkg/router/middleware/session"
// 	"github.com/kleister/kleister-api/pkg/store"
// )

// // KeyCreate creates a new key.
// func KeyCreate(c *gin.Context) {
// 	record := &model.Key{}

// 	if err := c.BindJSON(&record); err != nil {
// 		logrus.Warnf("Failed to bind key data. %s", err)

// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "Failed to bind key data",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	err := store.CreateKey(
// 		c,
// 		record,
// 	)

// 	if err != nil {
// 		logrus.Warnf("Failed to create key. %s", err)

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
