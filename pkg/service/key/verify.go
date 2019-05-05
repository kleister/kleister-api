package key

// import (
// 	"net/http"

// 	"github.com/Sirupsen/logrus"
// 	"github.com/gin-gonic/gin"
// 	"github.com/kleister/kleister-api/pkg/model"
// 	"github.com/kleister/kleister-api/pkg/router/middleware/session"
// 	"github.com/kleister/kleister-api/pkg/store"
// )

// // KeyVerify is a handler to verify a key for Technic.
// func KeyVerify(c *gin.Context) {
// 	record, res := store.GetKeyByValue(
// 		c,
// 		c.Param("key"),
// 	)

// 	if res.Error != nil || res.RecordNotFound() {
// 		c.JSON(
// 			http.StatusOK,
// 			gin.H{
// 				"error": "Invalid key provided",
// 			},
// 		)
// 	} else {
// 		c.JSON(
// 			http.StatusOK,
// 			gin.H{
// 				"valid":      "Valid key provided",
// 				"name":       record.Name,
// 				"created_at": record.CreatedAt,
// 			},
// 		)
// 	}
// }
