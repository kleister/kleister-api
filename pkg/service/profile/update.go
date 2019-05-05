package profile

// import (
// 	"net/http"

// 	"github.com/Sirupsen/logrus"
// 	"github.com/gin-gonic/gin"
// 	"github.com/kleister/kleister-api/pkg/router/middleware/session"
// 	"github.com/kleister/kleister-api/pkg/store"
// 	"github.com/kleister/kleister-api/pkg/token"
// )

// // ProfileUpdate updates the current profile.
// func ProfileUpdate(c *gin.Context) {
// 	record := session.Current(c)

// 	if err := c.BindJSON(&record); err != nil {
// 		logrus.Warnf("Failed to bind profile data. %s", err)

// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "Failed to bind profile data",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	err := store.UpdateUser(
// 		c,
// 		record,
// 	)

// 	if err != nil {
// 		logrus.Warnf("Failed to update profile. %s", err)

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
