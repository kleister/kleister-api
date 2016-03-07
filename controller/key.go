package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetKeys retrieves all available keys.
func GetKeys(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// GetKey retrieves a specific key.
func GetKey(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// DeleteKey removes a specific key.
func DeleteKey(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// PatchKey creates a new key.
func PatchKey(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// PostKey updates an existing key.
func PostKey(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}
