package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMinecraft(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

func PatchMinecraft(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}
