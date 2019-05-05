package minecraft

// import (
// 	"net/http"

// 	"github.com/Sirupsen/logrus"
// 	"github.com/gin-gonic/gin"
// 	"github.com/kleister/kleister-api/pkg/model"
// 	"github.com/kleister/kleister-api/pkg/store"
// )

// // MinecraftBuildIndex retrieves all builds related to a Minecraft version.
// func MinecraftBuildIndex(c *gin.Context) {
// 	records, err := store.GetMinecraftBuilds(
// 		c,
// 		&model.MinecraftBuildParams{
// 			Minecraft: c.Param("minecraft"),
// 		},
// 	)

// 	if err != nil {
// 		logrus.Warnf("Failed to fetch minecraft builds. %s", err)

// 		c.JSON(
// 			http.StatusInternalServerError,
// 			gin.H{
// 				"status":  http.StatusInternalServerError,
// 				"message": "Failed to fetch builds",
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

// // MinecraftBuildAppend appends a build to a Minecraft version.
// func MinecraftBuildAppend(c *gin.Context) {
// 	form := &model.MinecraftBuildParams{}

// 	if err := c.BindJSON(&form); err != nil {
// 		logrus.Warnf("Failed to bind minecraft build data. %s", err)

// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "Failed to bind minecraft build data",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	form.Minecraft = c.Param("minecraft")

// 	assigned := store.GetMinecraftHasBuild(
// 		c,
// 		form,
// 	)

// 	if assigned {
// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "Build is already appended",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	err := store.CreateMinecraftBuild(
// 		c,
// 		form,
// 	)

// 	if err != nil {
// 		logrus.Warnf("Failed to append minecraft build. %s", err)

// 		c.JSON(
// 			http.StatusInternalServerError,
// 			gin.H{
// 				"status":  http.StatusInternalServerError,
// 				"message": "Failed to append build",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	c.JSON(
// 		http.StatusOK,
// 		gin.H{
// 			"status":  http.StatusOK,
// 			"message": "Successfully appended build",
// 		},
// 	)
// }

// // MinecraftBuildDelete deleted a build from a Minecraft version
// func MinecraftBuildDelete(c *gin.Context) {
// 	form := &model.MinecraftBuildParams{}

// 	if err := c.BindJSON(&form); err != nil {
// 		logrus.Warnf("Failed to bind minecraft build data. %s", err)

// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "Failed to bind minecraft build data",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	form.Minecraft = c.Param("minecraft")

// 	assigned := store.GetMinecraftHasBuild(
// 		c,
// 		form,
// 	)

// 	if !assigned {
// 		c.JSON(
// 			http.StatusPreconditionFailed,
// 			gin.H{
// 				"status":  http.StatusPreconditionFailed,
// 				"message": "Build is not assigned",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	err := store.DeleteMinecraftBuild(
// 		c,
// 		form,
// 	)

// 	if err != nil {
// 		logrus.Warnf("Failed to delete minecraft build. %s", err)

// 		c.JSON(
// 			http.StatusInternalServerError,
// 			gin.H{
// 				"status":  http.StatusInternalServerError,
// 				"message": "Failed to unlink build",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	c.JSON(
// 		http.StatusOK,
// 		gin.H{
// 			"status":  http.StatusOK,
// 			"message": "Successfully unlinked build",
// 		},
// 	)
// }
