package api

import (
	"encoding/base32"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/config"
	"github.com/kleister/kleister-api/model"
	"github.com/kleister/kleister-api/router/middleware/session"
	"github.com/kleister/kleister-api/shared/token"
	"github.com/kleister/kleister-api/store"
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
		logrus.Warnf("Failed to refresh token. %s", err)

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
		logrus.Warn("Failed to bind login data. %s", err)

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
		logrus.Warnf("Failed to fetch requested user. %s", res.Error)

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
		logrus.Warnf("Failed to match passwords. %s", err)

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
		logrus.Warnf("Failed to generate token. %s", err)

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

// AuthVerify is a handler to verify an JWT token.
func AuthVerify(c *gin.Context) {
	var (
		record *model.User
	)

	_, err := token.Direct(
		c.Param("token"),
		func(t *token.Token) ([]byte, error) {
			var (
				res *gorm.DB
			)

			record, res = store.GetUser(
				c,
				t.Text,
			)

			signingKey, _ := base32.StdEncoding.DecodeString(record.Hash)
			return signingKey, res.Error
		},
	)

	if err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"error": "Invalid token provided",
			},
		)
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{
				"valid":      "Valid token provided",
				"name":       record.Username,
				"created_at": record.CreatedAt,
			},
		)
	}
}
