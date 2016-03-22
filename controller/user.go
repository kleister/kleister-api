package controller

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"
	"github.com/solderapp/solder-api/router/middleware/context"
	"github.com/solderapp/solder-api/router/middleware/session"
)

// GetUsers retrieves all available users.
func GetUsers(c *gin.Context) {
	records, err := context.Store(c).GetUsers()

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

	err := context.Store(c).Delete(
		&record,
	).Error

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

	err := context.Store(c).Save(
		&record,
	).Error

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

	err := context.Store(c).Create(
		&record,
	).Error

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
	user := session.User(c)
	records := &model.Mods{}

	err := context.Store(c).Model(
		&user,
	).Association(
		"Mods",
	).Find(
		&records,
	).Error

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
	user := session.User(c)
	mod := session.Mod(c)

	count := context.Store(c).Model(
		&user,
	).Association(
		"Mods",
	).Find(
		&mod,
	).Count()

	if count > 0 {
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

	err := context.Store(c).Model(
		&user,
	).Association(
		"Mods",
	).Append(
		model.Mod{
			ID: mod.ID,
		},
	).Error

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
	user := session.User(c)
	mod := session.Mod(c)

	count := context.Store(c).Model(
		&user,
	).Association(
		"Mods",
	).Find(
		&mod,
	).Count()

	if count < 1 {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "Mod is not assigned",
			},
		)

		c.Abort()
		return
	}

	err := context.Store(c).Model(
		&user,
	).Association(
		"Mods",
	).Delete(
		model.Mod{
			ID: mod.ID,
		},
	).Error

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
