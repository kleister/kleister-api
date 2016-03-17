package controller

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"
	"github.com/solderapp/solder-api/router/middleware/context"
	"github.com/solderapp/solder-api/router/middleware/session"
)

// GetMods retrieves all available mods.
func GetMods(c *gin.Context) {
	records := &model.Mods{}

	err := context.Store(c).Scopes(
		model.ModDefaultOrder,
	).Preload(
		"Versions",
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

// GetMod retrieves a specific mod.
func GetMod(c *gin.Context) {
	record := session.Mod(c)

	c.JSON(
		http.StatusOK,
		record,
	)
}

// DeleteMod removes a specific mod.
func DeleteMod(c *gin.Context) {
	record := session.Mod(c)

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
			"message": "Successfully deleted mod",
		},
	)
}

// PatchMod updates an existing mod.
func PatchMod(c *gin.Context) {
	record := session.Mod(c)

	if err := c.BindJSON(&record); err != nil {
		logrus.Warn("Failed to bind mod data")
		logrus.Warn(err)

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

// PostMod creates a new mod.
func PostMod(c *gin.Context) {
	record := &model.Mod{}

	if err := c.BindJSON(&record); err != nil {
		logrus.Warn("Failed to bind mod data")
		logrus.Warn(err)

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

// GetModUsers retrieves all users related to a mod.
func GetModUsers(c *gin.Context) {
	mod := session.Mod(c)
	records := &model.Users{}

	err := context.Store(c).Model(
		&mod,
	).Association(
		"Users",
	).Find(
		&records,
	).Error

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

// PatchModUser appends a user to a mod.
func PatchModUser(c *gin.Context) {
	mod := session.Mod(c)
	user := session.User(c)

	count := context.Store(c).Model(
		&mod,
	).Association(
		"Users",
	).Find(
		&user,
	).Count()

	if count > 0 {
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

	err := context.Store(c).Model(
		&mod,
	).Association(
		"Users",
	).Append(
		&user,
	).Error

	if err != nil {
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

// DeleteModUser deleted a user from a mod
func DeleteModUser(c *gin.Context) {
	mod := session.Mod(c)
	user := session.User(c)

	count := context.Store(c).Model(
		&mod,
	).Association(
		"Users",
	).Find(
		&user,
	).Count()

	if count < 1 {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "User is not assigned",
			},
		)

		c.Abort()
		return
	}

	err := context.Store(c).Model(
		&mod,
	).Association(
		"Users",
	).Delete(
		&user,
	).Error

	if err != nil {
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
