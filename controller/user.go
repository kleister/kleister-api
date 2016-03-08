package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers retrieves all available users.
func GetUsers(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// GetUser retrieves an specific user.
func GetUser(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// DeleteUser removes an specific user.
func DeleteUser(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// PatchUser updates an existing user.
func PatchUser(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// PostUser creates an new user.
func PostUser(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}
