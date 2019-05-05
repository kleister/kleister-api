package key

// import (
// 	"net/http"

// 	"github.com/Sirupsen/logrus"
// 	"github.com/gin-gonic/gin"
// 	"github.com/kleister/kleister-api/pkg/model"
// 	"github.com/kleister/kleister-api/pkg/router/middleware/session"
// 	"github.com/kleister/kleister-api/pkg/store"
// )

// // KeyDelete removes a specific key.
// func KeyDelete(c *gin.Context) {
// 	record := session.Key(c)

// 	err := store.DeleteKey(
// 		c,
// 		record,
// 	)

// 	if err != nil {
// 		logrus.Warnf("Failed to delete key. %s", err)

// 		c.JSON(
// 			http.StatusBadRequest,
// 			gin.H{
// 				"status":  http.StatusBadRequest,
// 				"message": "Failed to delete key",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	c.JSON(
// 		http.StatusOK,
// 		gin.H{
// 			"status":  http.StatusOK,
// 			"message": "Successfully deleted key",
// 		},
// 	)
// }
