package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProfile displays the current profile.
func GetProfile(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{},
	)
}

// PatchProfile updates the current profile.
func PatchProfile(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{},
	)
}
