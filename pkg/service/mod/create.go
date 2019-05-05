package mod

// import (
// 	"net/http"

// 	"github.com/Sirupsen/logrus"
// 	"github.com/gin-gonic/gin"
// 	"github.com/kleister/kleister-api/pkg/model"
// 	"github.com/kleister/kleister-api/pkg/router/middleware/session"
// 	"github.com/kleister/kleister-api/pkg/store"
// )

// // ModCreate creates a new mod.
// func ModCreate(c *gin.Context) {
// 	record := &model.Mod{}

// 	if err := c.BindJSON(&record); err != nil {
// 		logrus.Warnf("Failed to bind mod data. %s", err)

// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "Failed to bind mod data",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	err := store.CreateMod(
// 		c,
// 		record,
// 	)

// 	if err != nil {
// 		logrus.Warnf("Failed to create mod. %s", err)

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
