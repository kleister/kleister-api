package forge

// import (
// 	"net/http"

// 	"github.com/Sirupsen/logrus"
// 	"github.com/gin-gonic/gin"
// 	"github.com/kleister/kleister-api/pkg/model"
// 	"github.com/kleister/go-forge/version"
// 	"github.com/kleister/kleister-api/pkg/store"
// )

// // ForgeUpdate updates the list of available Forge versions.
// func ForgeUpdate(c *gin.Context) {
// 	result, err := version.FromDefault()

// 	if err != nil {
// 		logrus.Warnf("Failed to sync forge versions. %s", err)

// 		c.JSON(
// 			http.StatusServiceUnavailable,
// 			gin.H{
// 				"status":  http.StatusServiceUnavailable,
// 				"message": "Failed to request Forge versions",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	for _, ver := range result.Releases {
// 		_, err := store.SyncForge(
// 			c,
// 			ver,
// 		)

// 		if err != nil {
// 			logrus.Warnf("Failed to store forge version. %s", err)

// 			c.JSON(
// 				http.StatusInternalServerError,
// 				gin.H{
// 					"status":  http.StatusInternalServerError,
// 					"message": "Failed to store Forge versions",
// 				},
// 			)

// 			c.Abort()
// 			return
// 		}
// 	}

// 	c.JSON(
// 		http.StatusOK,
// 		gin.H{
// 			"status":  http.StatusOK,
// 			"message": "Successfully imported Forge versions",
// 		},
// 	)
// }
