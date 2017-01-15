package api

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/kleister/kleister-api/model"
	"github.com/kleister/kleister-api/router/middleware/session"
	"github.com/kleister/kleister-api/store"
)

// UserIndex retrieves all available users.
func UserIndex(c *gin.Context) {
	records, err := store.GetUsers(
		c,
	)

	if err != nil {
		logrus.Warnf("Failed to fetch users. %s", err)

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

// UserShow retrieves a specific user.
func UserShow(c *gin.Context) {
	record := session.User(c)

	c.JSON(
		http.StatusOK,
		record,
	)
}

// UserDelete removes a specific user.
func UserDelete(c *gin.Context) {
	record := session.User(c)

	err := store.DeleteUser(
		c,
		record,
	)

	if err != nil {
		logrus.Warnf("Failed to delete user. %s", err)

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Failed to delete user",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully deleted user",
		},
	)
}

// UserUpdate updates an existing user.
func UserUpdate(c *gin.Context) {
	record := session.User(c)

	if err := c.BindJSON(&record); err != nil {
		logrus.Warnf("Failed to bind user data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind user data",
			},
		)

		c.Abort()
		return
	}

	err := store.UpdateUser(
		c,
		record,
	)

	if err != nil {
		logrus.Warnf("Failed to update user. %s", err)

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

// UserCreate creates a new user.
func UserCreate(c *gin.Context) {
	record := &model.User{}

	if err := c.BindJSON(&record); err != nil {
		logrus.Warnf("Failed to bind user data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind user data",
			},
		)

		c.Abort()
		return
	}

	err := store.CreateUser(
		c,
		record,
	)

	if err != nil {
		logrus.Warnf("Failed to create user. %s", err)

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

// UserModIndex retrieves all mods related to a user.
func UserModIndex(c *gin.Context) {
	records, err := store.GetUserMods(
		c,
		&model.UserModParams{
			User: c.Param("user"),
		},
	)

	if err != nil {
		logrus.Warnf("Failed to fetch user mods. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to fetch mods",
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

// UserModAppend appends a mod to a user.
func UserModAppend(c *gin.Context) {
	form := &model.UserModParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind user mod data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind form data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetUserHasMod(
		c,
		form,
	)

	if assigned {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Mod is already appended",
			},
		)

		c.Abort()
		return
	}

	err := store.CreateUserMod(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to append user mod. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to append mod",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully appended mod",
		},
	)
}

// UserModPerm updates the user mod permission.
func UserModPerm(c *gin.Context) {
	form := &model.UserModParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind user mod data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind form data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetUserHasMod(
		c,
		form,
	)

	if !assigned {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Mod is not assigned",
			},
		)

		c.Abort()
		return
	}

	err := store.UpdateUserMod(
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

// UserModDelete deleted a mod from a user
func UserModDelete(c *gin.Context) {
	form := &model.UserModParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind user mod data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind form data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetUserHasMod(
		c,
		form,
	)

	if !assigned {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Mod is not assigned",
			},
		)

		c.Abort()
		return
	}

	err := store.DeleteUserMod(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to delete user mod. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to unlink mod",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully unlinked mod",
		},
	)
}

// UserPackIndex retrieves all packs related to a user.
func UserPackIndex(c *gin.Context) {
	records, err := store.GetUserPacks(
		c,
		&model.UserPackParams{
			User: c.Param("user"),
			Pack: c.Param("pack"),
		},
	)

	if err != nil {
		logrus.Warnf("Failed to fetch user packs. %s", err)

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

// UserPackAppend appends a pack to a user.
func UserPackAppend(c *gin.Context) {
	form := &model.UserPackParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind user pack data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind user pack data",
			},
		)

		c.Abort()
		return
	}

	form.User = c.Param("user")

	assigned := store.GetUserHasPack(
		c,
		form,
	)

	if assigned {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Pack is already appended",
			},
		)

		c.Abort()
		return
	}

	err := store.CreateUserPack(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to append user pack. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to append pack",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully appended pack",
		},
	)
}

// UserPackPerm updates the user pack permission.
func UserPackPerm(c *gin.Context) {
	form := &model.UserPackParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind user pack data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind user pack data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetUserHasPack(
		c,
		form,
	)

	if !assigned {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Pack is not assigned",
			},
		)

		c.Abort()
		return
	}

	err := store.UpdateUserPack(
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

// UserPackDelete deleted a pack from a user
func UserPackDelete(c *gin.Context) {
	form := &model.UserPackParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind user pack data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind user pack data",
			},
		)

		c.Abort()
		return
	}

	form.User = c.Param("user")

	assigned := store.GetUserHasPack(
		c,
		form,
	)

	if !assigned {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Pack is not assigned",
			},
		)

		c.Abort()
		return
	}

	err := store.DeleteUserPack(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to delete user pack. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to unlink pack",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully unlinked pack",
		},
	)
}

// UserTeamIndex retrieves all teams related to a user.
func UserTeamIndex(c *gin.Context) {
	records, err := store.GetUserTeams(
		c,
		&model.UserTeamParams{
			User: c.Param("user"),
		},
	)

	if err != nil {
		logrus.Warnf("Failed to fetch user teams. %s", err)

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

// UserTeamAppend appends a team to a user.
func UserTeamAppend(c *gin.Context) {
	form := &model.UserTeamParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind user team data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind user team data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetUserHasTeam(
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

	err := store.CreateUserTeam(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to append user team. %s", err)

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

// UserTeamPerm updates the user team permission.
func UserTeamPerm(c *gin.Context) {
	form := &model.UserTeamParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind user team data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind user team data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetUserHasTeam(
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

	err := store.UpdateUserTeam(
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

// UserTeamDelete deleted a team from a user
func UserTeamDelete(c *gin.Context) {
	form := &model.UserTeamParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind user team data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind user team data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetUserHasTeam(
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

	err := store.DeleteUserTeam(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to delete user team. %s", err)

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
