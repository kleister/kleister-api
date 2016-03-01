package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPacks(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

func GetPack(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

func DeletePack(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

func PatchPack(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

func PostPack(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}
