package api

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/kleister/kleister-api/model"
	"github.com/kleister/kleister-api/router/middleware/session"
	"github.com/kleister/kleister-api/store"
)

// TeamIndex retrieves all available teams.
func TeamIndex(c *gin.Context) {
	records, err := store.GetTeams(
		c,
	)

	if err != nil {
		logrus.Warnf("Failed to fetch teams. %s", err)

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

// TeamShow retrieves a specific team.
func TeamShow(c *gin.Context) {
	record := session.Team(c)

	c.JSON(
		http.StatusOK,
		record,
	)
}

// TeamDelete removes a specific team.
func TeamDelete(c *gin.Context) {
	record := session.Team(c)

	err := store.DeleteTeam(
		c,
		record,
	)

	if err != nil {
		logrus.Warnf("Failed to delete team. %s", err)

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Failed to delete team",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully deleted team",
		},
	)
}

// TeamUpdate updates an existing team.
func TeamUpdate(c *gin.Context) {
	record := session.Team(c)

	if err := c.BindJSON(&record); err != nil {
		logrus.Warnf("Failed to bind team data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind team data",
			},
		)

		c.Abort()
		return
	}

	err := store.UpdateTeam(
		c,
		record,
	)

	if err != nil {
		logrus.Warnf("Failed to update team. %s", err)

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

// TeamCreate creates a new user.
func TeamCreate(c *gin.Context) {
	record := &model.Team{}

	if err := c.BindJSON(&record); err != nil {
		logrus.Warnf("Failed to bind team data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind team data",
			},
		)

		c.Abort()
		return
	}

	err := store.CreateTeam(
		c,
		record,
	)

	if err != nil {
		logrus.Warnf("Failed to create team. %s", err)

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

// TeamUserIndex retrieves all users related to a team.
func TeamUserIndex(c *gin.Context) {
	records, err := store.GetTeamUsers(
		c,
		&model.TeamUserParams{
			Team: c.Param("team"),
		},
	)

	if err != nil {
		logrus.Warnf("Failed to fetch team users. %s", err)

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

// TeamUserAppend appends a user to a team.
func TeamUserAppend(c *gin.Context) {
	form := &model.TeamUserParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind team user data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind team user data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetTeamHasUser(
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

	err := store.CreateTeamUser(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to append team user. %s", err)

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

// TeamUserPerm updates the team user permission.
func TeamUserPerm(c *gin.Context) {
	form := &model.TeamUserParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind team user data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind team user data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetTeamHasUser(
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

	err := store.UpdateTeamUser(
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

// TeamUserDelete deleted a user from a team
func TeamUserDelete(c *gin.Context) {
	form := &model.TeamUserParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind team user data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind team user data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetTeamHasUser(
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

	err := store.DeleteTeamUser(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to delete team user. %s", err)

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

// TeamPackIndex retrieves all packs related to a team.
func TeamPackIndex(c *gin.Context) {
	records, err := store.GetTeamPacks(
		c,
		&model.TeamPackParams{
			Team: c.Param("team"),
		},
	)

	if err != nil {
		logrus.Warnf("Failed to fetch team packs. %s", err)

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

// TeamPackAppend appends a pack to a team.
func TeamPackAppend(c *gin.Context) {
	form := &model.TeamPackParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind team pack data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind team pack data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetTeamHasPack(
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

	err := store.CreateTeamPack(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to append team pack. %s", err)

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

// TeamPackPerm updates the team pack permission.
func TeamPackPerm(c *gin.Context) {
	form := &model.TeamPackParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind team pack data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind team pack data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetTeamHasPack(
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

	err := store.UpdateTeamPack(
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

// TeamPackDelete deleted a pack from a team
func TeamPackDelete(c *gin.Context) {
	form := &model.TeamPackParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind team pack data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind team pack data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetTeamHasPack(
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

	err := store.DeleteTeamPack(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to delete team pack. %s", err)

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

// TeamModIndex retrieves all mods related to a team.
func TeamModIndex(c *gin.Context) {
	records, err := store.GetTeamMods(
		c,
		&model.TeamModParams{
			Team: c.Param("team"),
		},
	)

	if err != nil {
		logrus.Warnf("Failed to fetch team mods. %s", err)

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

// TeamModAppend appends a mod to a team.
func TeamModAppend(c *gin.Context) {
	form := &model.TeamModParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind team mod data. %d", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind team mod data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetTeamHasMod(
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

	err := store.CreateTeamMod(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to append team mod. %s", err)

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

// TeamModPerm updates the team mod permission.
func TeamModPerm(c *gin.Context) {
	form := &model.TeamModParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind team mod data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind team mod data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetTeamHasMod(
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

	err := store.UpdateTeamMod(
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

// TeamModDelete deleted a mod from a team
func TeamModDelete(c *gin.Context) {
	form := &model.TeamModParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind team mod data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind team mod data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetTeamHasMod(
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

	err := store.DeleteTeamMod(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to delete team mod. %s", err)

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
