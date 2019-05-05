package minecraft

// import (
// 	"net/http"

// 	"github.com/Sirupsen/logrus"
// 	"github.com/gin-gonic/gin"
// 	"github.com/kleister/kleister-api/pkg/model"
// 	"github.com/kleister/go-minecraft/version"
// 	"github.com/kleister/kleister-api/pkg/store"
// )

// // MinecraftUpdate updates the list of available Minecraft versions.
// func MinecraftUpdate(c *gin.Context) {
// 	result, err := version.FromDefault()

// 	if err != nil {
// 		logrus.Warnf("Failed to sync minecraft versions. %s", err)

// 		c.JSON(
// 			http.StatusServiceUnavailable,
// 			gin.H{
// 				"status":  http.StatusServiceUnavailable,
// 				"message": "Failed to request Minecraft versions",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	for _, ver := range result.Releases {
// 		_, err := store.SyncMinecraft(
// 			c,
// 			ver,
// 		)

// 		if err != nil {
// 			logrus.Warnf("Failed to store minecraft version. %s", err)

// 			c.JSON(
// 				http.StatusInternalServerError,
// 				gin.H{
// 					"status":  http.StatusInternalServerError,
// 					"message": "Failed to store Minecraft versions",
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
// 			"message": "Successfully imported Minecraft versions",
// 		},
// 	)
// }
