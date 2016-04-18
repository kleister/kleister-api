package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetSolderPack retrieves the pack compatible to Technic Platform.
func GetSolderPack(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{},
	)
}

// GetSolderBuild retrieves the build compatible to Technic Platform.
func GetSolderBuild(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{},
	)
}

// GetSolderMod retrieves the mod compatible to Technic Platform.
func GetSolderMod(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{},
	)
}

// GetSolderVersion retrieves the version compatible to Technic Platform.
func GetSolderVersion(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{},
	)
}
