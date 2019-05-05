package profile

// import (
// 	"net/http"

// 	"github.com/Sirupsen/logrus"
// 	"github.com/gin-gonic/gin"
// 	"github.com/kleister/kleister-api/pkg/router/middleware/session"
// 	"github.com/kleister/kleister-api/pkg/store"
// 	"github.com/kleister/kleister-api/pkg/token"
// )

// // ProfileToken displays the users token.
// func ProfileToken(c *gin.Context) {
// 	record := session.Current(c)

// 	token := token.New(token.UserToken, record.Username)
// 	result, err := token.SignUnlimited(record.Hash)

// 	if err != nil {
// 		logrus.Warnf("Failed to generate token. %s", err)

// 		c.JSON(
// 			http.StatusInternalServerError,
// 			gin.H{
// 				"status":  http.StatusInternalServerError,
// 				"message": "Failed to generate token",
// 			},
// 		)

// 		c.Abort()
// 		return
// 	}

// 	c.JSON(
// 		http.StatusOK,
// 		result,
// 	)
// }
