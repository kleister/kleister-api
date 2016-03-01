package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBuilds(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

func GetBuild(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

func DeleteBuild(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

func PatchBuild(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}

func PostBuild(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}
