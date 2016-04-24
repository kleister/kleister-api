package api

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"
	"github.com/solderapp/solder-api/router/middleware/session"
	"github.com/solderapp/solder-api/store"
)

// GetUsers retrieves all available users.
func GetUsers(c *gin.Context) {
	records, err := store.GetUsers(
		c,
	)

	if err != nil {
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

// GetUser retrieves a specific user.
func GetUser(c *gin.Context) {
	record := session.User(c)

	c.JSON(
		http.StatusOK,
		record,
	)
}

// DeleteUser removes a specific user.
func DeleteUser(c *gin.Context) {
	record := session.User(c)

	err := store.DeleteUser(
		c,
		record,
	)

	if err != nil {
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
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully deleted user",
		},
	)
}

// PatchUser updates an existing user.
func PatchUser(c *gin.Context) {
	record := session.User(c)

	if err := c.BindJSON(&record); err != nil {
		logrus.Warn("Failed to bind user data")
		logrus.Warn(err)

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

// PostUser creates a new user.
func PostUser(c *gin.Context) {
	record := &model.User{
		Permission: &model.Permission{},
	}

	if err := c.BindJSON(&record); err != nil {
		logrus.Warn("Failed to bind user data")
		logrus.Warn(err)

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

// GetUserMods retrieves all mods related to a user.
func GetUserMods(c *gin.Context) {
	records, err := store.GetUserMods(
		c,
		&model.UserModParams{
			User: c.Param("user"),
		},
	)

	if err != nil {
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

// PatchUserMod appends a mod to a user.
func PatchUserMod(c *gin.Context) {
	assigned := store.GetUserHasMod(
		c,
		&model.UserModParams{
			User: c.Param("user"),
			Mod:  c.Param("mod"),
		},
	)

	if assigned == true {
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
		&model.UserModParams{
			User: c.Param("user"),
			Mod:  c.Param("mod"),
		},
	)

	if err != nil {
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

// DeleteUserMod deleted a mod from a user
func DeleteUserMod(c *gin.Context) {
	assigned := store.GetUserHasMod(
		c,
		&model.UserModParams{
			User: c.Param("user"),
			Mod:  c.Param("mod"),
		},
	)

	if assigned == false {
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
		&model.UserModParams{
			User: c.Param("user"),
			Mod:  c.Param("mod"),
		},
	)

	if err != nil {
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

// GetUserPacks retrieves all packs related to a user.
func GetUserPacks(c *gin.Context) {
	records, err := store.GetUserPacks(
		c,
		&model.UserPackParams{
			User: c.Param("user"),
			Pack: c.Param("pack"),
		},
	)

	if err != nil {
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

// PatchUserPack appends a pack to a user.
func PatchUserPack(c *gin.Context) {
	assigned := store.GetUserHasPack(
		c,
		&model.UserPackParams{
			User: c.Param("user"),
			Pack: c.Param("pack"),
		},
	)

	if assigned == true {
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
		&model.UserPackParams{
			User: c.Param("user"),
			Pack: c.Param("pack"),
		},
	)

	if err != nil {
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

// DeleteUserPack deleted a pack from a user
func DeleteUserPack(c *gin.Context) {
	assigned := store.GetUserHasPack(
		c,
		&model.UserPackParams{
			User: c.Param("user"),
			Pack: c.Param("pack"),
		},
	)

	if assigned == false {
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
		&model.UserPackParams{
			User: c.Param("user"),
			Pack: c.Param("pack"),
		},
	)

	if err != nil {
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
