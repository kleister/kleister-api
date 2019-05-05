package team

// import (
// 	"net/http"

// 	"github.com/Sirupsen/logrus"
// 	"github.com/gin-gonic/gin"
// 	"github.com/kleister/kleister-api/pkg/model"
// 	"github.com/kleister/kleister-api/pkg/router/middleware/session"
// 	"github.com/kleister/kleister-api/pkg/store"
// )

// // TeamModIndex retrieves all mods related to a team.
// func TeamModIndex(c *gin.Context) {
// 	records, err := store.GetTeamMods(
// 		c,
// 		&model.TeamModParams{
// 			Team: c.Param("team"),
// 		},
// 	)

// 	if err != nil {
// 		logrus.Warnf("Failed to fetch team mods. %s", err)

// 		c.JSON(
// 			http.StatusInternalServerError,
// 			gin.H{
// 				"status":  http.StatusInternalServerError,
// 				"message": "Failed to fetch mods",
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

// // TeamModAppend appends a mod to a team.
// func TeamModAppend(c *gin.Context) {
// 	form := &model.TeamModParams{}

// 	if err := c.BindJSON(&form); err != nil {
// 		logrus.Warnf("Failed to bind team mod data. %d", err)

// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "Failed to bind team mod data",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	assigned := store.GetTeamHasMod(
// 		c,
// 		form,
// 	)

// 	if assigned {
// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "Mod is already appended",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	err := store.CreateTeamMod(
// 		c,
// 		form,
// 	)

// 	if err != nil {
// 		logrus.Warnf("Failed to append team mod. %s", err)

// 		c.JSON(
// 			http.StatusInternalServerError,
// 			gin.H{
// 				"status":  http.StatusInternalServerError,
// 				"message": "Failed to append mod",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	c.JSON(
// 		http.StatusOK,
// 		gin.H{
// 			"status":  http.StatusOK,
// 			"message": "Successfully appended mod",
// 		},
// 	)
// }

// // TeamModPerm updates the team mod permission.
// func TeamModPerm(c *gin.Context) {
// 	form := &model.TeamModParams{}

// 	if err := c.BindJSON(&form); err != nil {
// 		logrus.Warnf("Failed to bind team mod data. %s", err)

// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "Failed to bind team mod data",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	assigned := store.GetTeamHasMod(
// 		c,
// 		form,
// 	)

// 	if !assigned {
// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "Mod is not assigned",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	err := store.UpdateTeamMod(
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

// // TeamModDelete deleted a mod from a team
// func TeamModDelete(c *gin.Context) {
// 	form := &model.TeamModParams{}

// 	if err := c.BindJSON(&form); err != nil {
// 		logrus.Warnf("Failed to bind team mod data. %s", err)

// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "Failed to bind team mod data",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	assigned := store.GetTeamHasMod(
// 		c,
// 		form,
// 	)

// 	if !assigned {
// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "Mod is not assigned",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	err := store.DeleteTeamMod(
// 		c,
// 		form,
// 	)

// 	if err != nil {
// 		logrus.Warnf("Failed to delete team mod. %s", err)

// 		c.JSON(
// 			http.StatusInternalServerError,
// 			gin.H{
// 				"status":  http.StatusInternalServerError,
// 				"message": "Failed to unlink mod",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	c.JSON(
// 		http.StatusOK,
// 		gin.H{
// 			"status":  http.StatusOK,
// 			"message": "Successfully unlinked mod",
// 		},
// 	)
// }
