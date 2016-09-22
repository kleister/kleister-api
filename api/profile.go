package api

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/kleister/kleister-api/router/middleware/session"
	"github.com/kleister/kleister-api/shared/token"
	"github.com/kleister/kleister-api/store"
)

// ProfileShow displays the current profile.
func ProfileShow(c *gin.Context) {
	record := session.Current(c)

	c.JSON(
		http.StatusOK,
		record,
	)
}

// ProfileUpdate updates the current profile.
func ProfileUpdate(c *gin.Context) {
	record := session.Current(c)

	if err := c.BindJSON(&record); err != nil {
		logrus.Warnf("Failed to bind profile data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind profile data",
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
		logrus.Warnf("Failed to update profile. %s", err)

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

// ProfileToken displays the users token.
func ProfileToken(c *gin.Context) {
	record := session.Current(c)

	token := token.New(token.UserToken, record.Username)
	result, err := token.SignUnlimited(record.Hash)

	if err != nil {
		logrus.Warnf("Failed to generate token. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to generate token",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		result,
	)
}
