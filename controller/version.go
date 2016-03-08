package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetVersions retrieves all available versions.
func GetVersions(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{},
	)
}

// GetVersion retrieves a specific version.
func GetVersion(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{},
	)
}

// DeleteVersion removes a specific version.
func DeleteVersion(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{},
	)
}

// PatchVersion updates an existing version.
func PatchVersion(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{},
	)
}

// PostVersion creates a new version.
func PostVersion(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{},
	)
}
