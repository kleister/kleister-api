package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetClients retrieves all available clients.
func GetClients(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// GetClient retrieves a specific client.
func GetClient(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// DeleteClient removes a specific client.
func DeleteClient(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// PatchClient creates a new client.
func PatchClient(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// PostClient updates an existing client.
func PostClient(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}
