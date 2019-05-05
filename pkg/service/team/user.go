package team

// import (
// 	"net/http"

// 	"github.com/Sirupsen/logrus"
// 	"github.com/gin-gonic/gin"
// 	"github.com/kleister/kleister-api/pkg/model"
// 	"github.com/kleister/kleister-api/pkg/router/middleware/session"
// 	"github.com/kleister/kleister-api/pkg/store"
// )

// // TeamUserIndex retrieves all users related to a team.
// func TeamUserIndex(c *gin.Context) {
// 	records, err := store.GetTeamUsers(
// 		c,
// 		&model.TeamUserParams{
// 			Team: c.Param("team"),
// 		},
// 	)

// 	if err != nil {
// 		logrus.Warnf("Failed to fetch team users. %s", err)

// 		c.JSON(
// 			http.StatusInternalServerError,
// 			gin.H{
// 				"status":  http.StatusInternalServerError,
// 				"message": "Failed to fetch users",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	c.JSON(
// 		http.StatusOK,
// 		records,
// 	)
// }

// // TeamUserAppend appends a user to a team.
// func TeamUserAppend(c *gin.Context) {
// 	form := &model.TeamUserParams{}

// 	if err := c.BindJSON(&form); err != nil {
// 		logrus.Warnf("Failed to bind team user data. %s", err)

// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "Failed to bind team user data",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	assigned := store.GetTeamHasUser(
// 		c,
// 		form,
// 	)

// 	if assigned {
// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "User is already appended",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	err := store.CreateTeamUser(
// 		c,
// 		form,
// 	)

// 	if err != nil {
// 		logrus.Warnf("Failed to append team user. %s", err)

// 		c.JSON(
// 			http.StatusInternalServerError,
// 			gin.H{
// 				"status":  http.StatusInternalServerError,
// 				"message": "Failed to append user",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	c.JSON(
// 		http.StatusOK,
// 		gin.H{
// 			"status":  http.StatusOK,
// 			"message": "Successfully appended user",
// 		},
// 	)
// }

// // TeamUserPerm updates the team user permission.
// func TeamUserPerm(c *gin.Context) {
// 	form := &model.TeamUserParams{}

// 	if err := c.BindJSON(&form); err != nil {
// 		logrus.Warnf("Failed to bind team user data. %s", err)

// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "Failed to bind team user data",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	assigned := store.GetTeamHasUser(
// 		c,
// 		form,
// 	)

// 	if !assigned {
// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "User is not assigned",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	err := store.UpdateTeamUser(
// 		c,
// 		form,
// 	)

// 	if err != nil {
// 		logrus.Warnf("Failed to update permissions. %s", err)

// 		c.JSON(
// 			http.StatusInternalServerError,
// 			gin.H{
// 				"status":  http.StatusInternalServerError,
// 				"message": "Failed to update permissions",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	c.JSON(
// 		http.StatusOK,
// 		gin.H{
// 			"status":  http.StatusOK,
// 			"message": "Successfully updated permissions",
// 		},
// 	)
// }

// // TeamUserDelete deleted a user from a team
// func TeamUserDelete(c *gin.Context) {
// 	form := &model.TeamUserParams{}

// 	if err := c.BindJSON(&form); err != nil {
// 		logrus.Warnf("Failed to bind team user data. %s", err)

// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "Failed to bind team user data",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	assigned := store.GetTeamHasUser(
// 		c,
// 		form,
// 	)

// 	if !assigned {
// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "User is not assigned",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	err := store.DeleteTeamUser(
// 		c,
// 		form,
// 	)

// 	if err != nil {
// 		logrus.Warnf("Failed to delete team user. %s", err)

// 		c.JSON(
// 			http.StatusInternalServerError,
// 			gin.H{
// 				"status":  http.StatusInternalServerError,
// 				"message": "Failed to unlink user",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	c.JSON(
// 		http.StatusOK,
// 		gin.H{
// 			"status":  http.StatusOK,
// 			"message": "Successfully unlinked user",
// 		},
// 	)
// }
