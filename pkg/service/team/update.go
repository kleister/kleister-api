package team

// import (
// 	"net/http"

// 	"github.com/Sirupsen/logrus"
// 	"github.com/gin-gonic/gin"
// 	"github.com/kleister/kleister-api/pkg/model"
// 	"github.com/kleister/kleister-api/pkg/router/middleware/session"
// 	"github.com/kleister/kleister-api/pkg/store"
// )

// // TeamUpdate updates an existing team.
// func TeamUpdate(c *gin.Context) {
// 	record := session.Team(c)

// 	if err := c.BindJSON(&record); err != nil {
// 		logrus.Warnf("Failed to bind team data. %s", err)

// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "Failed to bind team data",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	err := store.UpdateTeam(
// 		c,
// 		record,
// 	)

// 	if err != nil {
// 		logrus.Warnf("Failed to update team. %s", err)

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
