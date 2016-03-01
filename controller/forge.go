package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetForge(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

func PatchForge(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}
