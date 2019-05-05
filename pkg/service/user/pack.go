package user

// import (
// 	"net/http"

// 	"github.com/Sirupsen/logrus"
// 	"github.com/gin-gonic/gin"
// 	"github.com/kleister/kleister-api/pkg/model"
// 	"github.com/kleister/kleister-api/pkg/router/middleware/session"
// 	"github.com/kleister/kleister-api/pkg/store"
// )

// // UserPackIndex retrieves all packs related to a user.
// func UserPackIndex(c *gin.Context) {
// 	records, err := store.GetUserPacks(
// 		c,
// 		&model.UserPackParams{
// 			User: c.Param("user"),
// 			Pack: c.Param("pack"),
// 		},
// 	)

// 	if err != nil {
// 		logrus.Warnf("Failed to fetch user packs. %s", err)

// 		c.JSON(
// 			http.StatusInternalServerError,
// 			gin.H{
// 				"status":  http.StatusInternalServerError,
// 				"message": "Failed to fetch packs",
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

// // UserPackAppend appends a pack to a user.
// func UserPackAppend(c *gin.Context) {
// 	form := &model.UserPackParams{}

// 	if err := c.BindJSON(&form); err != nil {
// 		logrus.Warnf("Failed to bind user pack data. %s", err)

// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "Failed to bind user pack data",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	form.User = c.Param("user")

// 	assigned := store.GetUserHasPack(
// 		c,
// 		form,
// 	)

// 	if assigned {
// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "Pack is already appended",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	err := store.CreateUserPack(
// 		c,
// 		form,
// 	)

// 	if err != nil {
// 		logrus.Warnf("Failed to append user pack. %s", err)

// 		c.JSON(
// 			http.StatusInternalServerError,
// 			gin.H{
// 				"status":  http.StatusInternalServerError,
// 				"message": "Failed to append pack",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	c.JSON(
// 		http.StatusOK,
// 		gin.H{
// 			"status":  http.StatusOK,
// 			"message": "Successfully appended pack",
// 		},
// 	)
// }

// // UserPackPerm updates the user pack permission.
// func UserPackPerm(c *gin.Context) {
// 	form := &model.UserPackParams{}

// 	if err := c.BindJSON(&form); err != nil {
// 		logrus.Warnf("Failed to bind user pack data. %s", err)

// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "Failed to bind user pack data",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	assigned := store.GetUserHasPack(
// 		c,
// 		form,
// 	)

// 	if !assigned {
// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "Pack is not assigned",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	err := store.UpdateUserPack(
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

// // UserPackDelete deleted a pack from a user
// func UserPackDelete(c *gin.Context) {
// 	form := &model.UserPackParams{}

// 	if err := c.BindJSON(&form); err != nil {
// 		logrus.Warnf("Failed to bind user pack data. %s", err)

// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "Failed to bind user pack data",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	form.User = c.Param("user")

// 	assigned := store.GetUserHasPack(
// 		c,
// 		form,
// 	)

// 	if !assigned {
// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "Pack is not assigned",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	err := store.DeleteUserPack(
// 		c,
// 		form,
// 	)

// 	if err != nil {
// 		logrus.Warnf("Failed to delete user pack. %s", err)

// 		c.JSON(
// 			http.StatusInternalServerError,
// 			gin.H{
// 				"status":  http.StatusInternalServerError,
// 				"message": "Failed to unlink pack",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	c.JSON(
// 		http.StatusOK,
// 		gin.H{
// 			"status":  http.StatusOK,
// 			"message": "Successfully unlinked pack",
// 		},
// 	)
// }
