package api

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/kleister/kleister-api/model"
	"github.com/kleister/kleister-api/router/middleware/session"
	"github.com/kleister/kleister-api/store"
)

// ModIndex retrieves all available mods.
func ModIndex(c *gin.Context) {
	records, err := store.GetMods(
		c,
	)

	if err != nil {
		logrus.Warnf("Failed to fetch mods. %s", err)

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

// ModShow retrieves a specific mod.
func ModShow(c *gin.Context) {
	record := session.Mod(c)

	c.JSON(
		http.StatusOK,
		record,
	)
}

// ModDelete removes a specific mod.
func ModDelete(c *gin.Context) {
	record := session.Mod(c)

	err := store.DeleteMod(
		c,
		record,
	)

	if err != nil {
		logrus.Warnf("Failed to delete mod. %s", err)

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Failed to delete mod",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully deleted mod",
		},
	)
}

// ModUpdate updates an existing mod.
func ModUpdate(c *gin.Context) {
	record := session.Mod(c)

	if err := c.BindJSON(&record); err != nil {
		logrus.Warnf("Failed to bind mod data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind mod data",
			},
		)

		c.Abort()
		return
	}

	err := store.UpdateMod(
		c,
		record,
	)

	if err != nil {
		logrus.Warnf("Failed to update mod. %s", err)

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

// ModCreate creates a new mod.
func ModCreate(c *gin.Context) {
	record := &model.Mod{}

	if err := c.BindJSON(&record); err != nil {
		logrus.Warnf("Failed to bind mod data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind mod data",
			},
		)

		c.Abort()
		return
	}

	err := store.CreateMod(
		c,
		record,
	)

	if err != nil {
		logrus.Warnf("Failed to create mod. %s", err)

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

// ModUserIndex retrieves all users related to a mod.
func ModUserIndex(c *gin.Context) {
	records, err := store.GetModUsers(
		c,
		&model.ModUserParams{
			Mod: c.Param("mod"),
		},
	)

	if err != nil {
		logrus.Warnf("Failed to fetch mod users. %s", err)

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

// ModUserAppend appends a user to a mod.
func ModUserAppend(c *gin.Context) {
	form := &model.ModUserParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind mod user data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind mod user data",
			},
		)

		c.Abort()
		return
	}

	form.Mod = c.Param("mod")

	assigned := store.GetModHasUser(
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

	err := store.CreateModUser(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to append mod user. %s", err)

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

// ModUserPerm updates the mod user permission.
func ModUserPerm(c *gin.Context) {
	form := &model.ModUserParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind mod user data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind mod user data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetModHasUser(
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

	err := store.UpdateModUser(
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

// ModUserDelete deleted a user from a mod
func ModUserDelete(c *gin.Context) {
	form := &model.ModUserParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind mod user data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind mod user data",
			},
		)

		c.Abort()
		return
	}

	form.Mod = c.Param("mod")

	assigned := store.GetModHasUser(
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

	err := store.DeleteModUser(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to delete mod user. %s", err)

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

// ModTeamIndex retrieves all teams related to a mod.
func ModTeamIndex(c *gin.Context) {
	records, err := store.GetModTeams(
		c,
		&model.ModTeamParams{
			Mod: c.Param("mod"),
		},
	)

	if err != nil {
		logrus.Warnf("Failed to fetch mod teams. %s", err)

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

// ModTeamAppend appends a team to a mod.
func ModTeamAppend(c *gin.Context) {
	form := &model.ModTeamParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind mod team data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind mod team data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetModHasTeam(
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

	err := store.CreateModTeam(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to append mod team. %s", err)

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

// ModTeamPerm updates the mod team permission.
func ModTeamPerm(c *gin.Context) {
	form := &model.ModTeamParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind mod team data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind mod team data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetModHasTeam(
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

	err := store.UpdateModTeam(
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

// ModTeamDelete deleted a team from a mod
func ModTeamDelete(c *gin.Context) {
	form := &model.ModTeamParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind mod team data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind mod team data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetModHasTeam(
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

	err := store.DeleteModTeam(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to delete mod team. %s", err)

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
