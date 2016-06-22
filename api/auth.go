package api

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/config"
	"github.com/solderapp/solder-api/model"
	"github.com/solderapp/solder-api/router/middleware/session"
	"github.com/solderapp/solder-api/shared/token"
	"github.com/solderapp/solder-api/store"
)

// AuthLogout represents the logout handler.
func AuthLogout(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{},
	)
}

// AuthRefresh represents the refresh handler.
func AuthRefresh(c *gin.Context) {
	record := session.Current(c)

	token := token.New(token.SessToken, record.Username)
	result, err := token.SignExpiring(record.Hash, config.Session.Expire)

	if err != nil {
		logrus.Warnf("Failed to refresh token: %s", err)

		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Failed to refresh token",
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

// AuthLogin represents the login handler.
func AuthLogin(c *gin.Context) {
	auth := &model.Auth{}

	if err := c.BindJSON(&auth); err != nil {
		logrus.Warn("Failed to bind login data: %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind login data",
			},
		)

		c.Abort()
		return
	}

	user, res := store.GetUser(
		c,
		auth.Username,
	)

	if res.Error != nil || res.RecordNotFound() {
		logrus.Warnf("Failed to fetch requested user: %s", res.Error)

		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Wrong username or password",
			},
		)

		c.Abort()
		return
	}

	if err := user.MatchPassword(auth.Password); err != nil {
		logrus.Warnf("Failed to match passwords: %s", err)

		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Wrong username or password",
			},
		)

		c.Abort()
		return
	}

	token := token.New(token.SessToken, user.Username)
	result, err := token.SignExpiring(user.Hash, config.Session.Expire)

	if err != nil {
		logrus.Warnf("Failed to generate token: %s", err)

		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Wrong username or password",
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
