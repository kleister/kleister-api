package forge

// import (
// 	"net/http"

// 	"github.com/Sirupsen/logrus"
// 	"github.com/gin-gonic/gin"
// 	"github.com/kleister/kleister-api/pkg/model"
// 	"github.com/kleister/kleister-api/pkg/store"
// )

// // ForgeIndex retrieves all available Forge versions.
// func ForgeIndex(c *gin.Context) {
// 	records, err := store.GetForges(
// 		c,
// 	)

// 	if err != nil {
// 		logrus.Warnf("Failed to fetch forge versions. %s", err)

// 		c.JSON(
// 			http.StatusInternalServerError,
// 			gin.H{
// 				"status":  http.StatusInternalServerError,
// 				"message": "Failed to fetch Forge versions",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	if c.Param("forge") != "" {
// 		records = records.Filter(
// 			c.Param("forge"),
// 		)
// 	}

// 	c.JSON(
// 		http.StatusOK,
// 		records,
// 	)
// }
