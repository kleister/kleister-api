package minecraft

// import (
// 	"net/http"

// 	"github.com/Sirupsen/logrus"
// 	"github.com/gin-gonic/gin"
// 	"github.com/kleister/kleister-api/pkg/model"
// 	"github.com/kleister/kleister-api/pkg/store"
// )

// // MinecraftIndex retrieves all available Minecraft versions.
// func MinecraftIndex(c *gin.Context) {
// 	records, err := store.GetMinecrafts(
// 		c,
// 	)

// 	if err != nil {
// 		logrus.Warnf("Failed to fetch minecraft versions. %s", err)

// 		c.JSON(
// 			http.StatusInternalServerError,
// 			gin.H{
// 				"status":  http.StatusInternalServerError,
// 				"message": "Failed to fetch Minecraft versions",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	if c.Param("minecraft") != "" {
// 		records = records.Filter(
// 			c.Param("minecraft"),
// 		)
// 	}

// 	c.JSON(
// 		http.StatusOK,
// 		records,
// 	)
// }
