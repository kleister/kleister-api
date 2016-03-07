package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetVersions retrieves all available versions.
func GetVersions(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// GetVersion retrieves a specific version.
func GetVersion(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// DeleteVersion removes a specific version.
func DeleteVersion(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// PatchVersion creates a new version.
func PatchVersion(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

// PostVersion updates an existing version.
func PostVersion(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}
