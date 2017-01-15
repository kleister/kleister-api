package api

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/kleister/kleister-api/model"
	"github.com/kleister/kleister-api/router/middleware/session"
	"github.com/kleister/kleister-api/store"
)

// PackIndex retrieves all available packs.
func PackIndex(c *gin.Context) {
	records, err := store.GetPacks(
		c,
	)

	if err != nil {
		logrus.Warnf("Failed to fetch packs. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to fetch packs",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		records,
	)
}

// PackShow retrieves a specific pack.
func PackShow(c *gin.Context) {
	record := session.Pack(c)

	c.JSON(
		http.StatusOK,
		record,
	)
}

// PackDelete removes a specific pack.
func PackDelete(c *gin.Context) {
	record := session.Pack(c)

	err := store.DeletePack(
		c,
		record,
	)

	if err != nil {
		logrus.Warnf("Failed to delete pack. %s", err)

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Failed to delete pack",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully deleted pack",
		},
	)
}

// PackUpdate updates an existing pack.
func PackUpdate(c *gin.Context) {
	record := session.Pack(c)

	if err := c.BindJSON(&record); err != nil {
		logrus.Warnf("Failed to bind pack data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind pack data",
			},
		)

		c.Abort()
		return
	}

	err := store.UpdatePack(
		c,
		record,
	)

	if err != nil {
		logrus.Warnf("Failed to update pack. %s", err)

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		record,
	)
}

// PackCreate creates a new pack.
func PackCreate(c *gin.Context) {
	record := &model.Pack{}

	if err := c.BindJSON(&record); err != nil {
		logrus.Warnf("Failed to bind pack data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind pack data",
			},
		)

		c.Abort()
		return
	}

	err := store.CreatePack(
		c,
		record,
	)

	if err != nil {
		logrus.Warnf("Failed to create pack. %s", err)

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		record,
	)
}

// PackClientIndex retrieves all clients related to a pack.
func PackClientIndex(c *gin.Context) {
	records, err := store.GetPackClients(
		c,
		&model.PackClientParams{
			Pack: c.Param("pack"),
		},
	)

	if err != nil {
		logrus.Warnf("Failed to fetch pack clients. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to fetch clients",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		records,
	)
}

// PackClientAppend appends a client to a pack.
func PackClientAppend(c *gin.Context) {
	form := &model.PackClientParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind pack client data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind pack client data",
			},
		)

		c.Abort()
		return
	}

	form.Pack = c.Param("pack")

	assigned := store.GetPackHasClient(
		c,
		form,
	)

	if assigned {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Client is already appended",
			},
		)

		c.Abort()
		return
	}

	err := store.CreatePackClient(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to append pack client. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to append client",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully appended client",
		},
	)
}

// PackClientDelete deleted a client from a pack
func PackClientDelete(c *gin.Context) {
	form := &model.PackClientParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind pack client data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind pack client data",
			},
		)

		c.Abort()
		return
	}

	form.Pack = c.Param("pack")

	assigned := store.GetPackHasClient(
		c,
		form,
	)

	if !assigned {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Client is not assigned",
			},
		)

		c.Abort()
		return
	}

	err := store.DeletePackClient(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to delete pack client. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to unlink client",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully unlinked client",
		},
	)
}

// PackUserIndex retrieves all users related to a pack.
func PackUserIndex(c *gin.Context) {
	records, err := store.GetPackUsers(
		c,
		&model.PackUserParams{
			Pack: c.Param("pack"),
		},
	)

	if err != nil {
		logrus.Warnf("Failed to fetch pack users. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to fetch users",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		records,
	)
}

// PackUserAppend appends a user to a pack.
func PackUserAppend(c *gin.Context) {
	form := &model.PackUserParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind pack user data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind pack user data",
			},
		)

		c.Abort()
		return
	}

	form.Pack = c.Param("pack")

	assigned := store.GetPackHasUser(
		c,
		form,
	)

	if assigned {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "User is already appended",
			},
		)

		c.Abort()
		return
	}

	err := store.CreatePackUser(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to append pack user. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to append user",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully appended user",
		},
	)
}

// PackUserPerm updates the pack user permission.
func PackUserPerm(c *gin.Context) {
	form := &model.PackUserParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind pack user data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind pack user data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetPackHasUser(
		c,
		form,
	)

	if !assigned {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "User is not assigned",
			},
		)

		c.Abort()
		return
	}

	err := store.UpdatePackUser(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to update permissions. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to update permissions",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully updated permissions",
		},
	)
}

// PackUserDelete deleted a user from a pack
func PackUserDelete(c *gin.Context) {
	form := &model.PackUserParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind pack user data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind pack user data",
			},
		)

		c.Abort()
		return
	}

	form.Pack = c.Param("pack")

	assigned := store.GetPackHasUser(
		c,
		form,
	)

	if !assigned {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "User is not assigned",
			},
		)

		c.Abort()
		return
	}

	err := store.DeletePackUser(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to delete pack user. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to unlink user",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully unlinked user",
		},
	)
}

// PackTeamIndex retrieves all teams related to a pack.
func PackTeamIndex(c *gin.Context) {
	records, err := store.GetPackTeams(
		c,
		&model.PackTeamParams{
			Pack: c.Param("pack"),
		},
	)

	if err != nil {
		logrus.Warnf("Failed to fetch pack teams. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to fetch teams",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		records,
	)
}

// PackTeamAppend appends a team to a pack.
func PackTeamAppend(c *gin.Context) {
	form := &model.PackTeamParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind pack team data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind pack team data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetPackHasTeam(
		c,
		form,
	)

	if assigned {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Team is already appended",
			},
		)

		c.Abort()
		return
	}

	err := store.CreatePackTeam(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to append pack team. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to append team",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully appended team",
		},
	)
}

// PackTeamPerm updates the pack team permission.
func PackTeamPerm(c *gin.Context) {
	form := &model.PackTeamParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind pack team data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind pack team data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetPackHasTeam(
		c,
		form,
	)

	if !assigned {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Team is not assigned",
			},
		)

		c.Abort()
		return
	}

	err := store.UpdatePackTeam(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to update permissions. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to update permissions",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully updated permissions",
		},
	)
}

// PackTeamDelete deleted a team from a pack
func PackTeamDelete(c *gin.Context) {
	form := &model.PackTeamParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind pack team data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind pack team data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetPackHasTeam(
		c,
		form,
	)

	if !assigned {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Team is not assigned",
			},
		)

		c.Abort()
		return
	}

	err := store.DeletePackTeam(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to delete pack team. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to unlink team",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully unlinked team",
		},
	)
}
