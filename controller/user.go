package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

func GetUser(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

func DeleteUser(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

func PatchUser(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

func PostUser(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}
