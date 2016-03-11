package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"
	"github.com/solderapp/solder-api/router/middleware/context"
	"github.com/solderapp/solder-api/router/middleware/session"
)

// GetBuilds retrieves all available builds.
func GetBuilds(c *gin.Context) {
	pack := session.Pack(c)
	records := &model.Builds{}

	err := context.Store(c).Scopes(
		model.BuildDefaultOrder,
	).Where(
		"builds.pack_id = ?",
		pack.ID,
	).Preload(
		"Pack",
	).Find(
		&records,
	).Error

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to fetch builds",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		records,
	)
}

// GetBuild retrieves a specific build.
func GetBuild(c *gin.Context) {
	record := session.Build(c)

	c.JSON(
		http.StatusOK,
		record,
	)
}

// DeleteBuild removes a specific build.
func DeleteBuild(c *gin.Context) {
	record := session.Build(c)

	err := context.Store(c).Delete(
		&record,
	).Error

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully deleted build",
		},
	)
}

// PatchBuild updates an existing build.
func PatchBuild(c *gin.Context) {
	record := session.Build(c)

	if err := c.BindJSON(&record); err != nil {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind build data",
			},
		)

		c.Abort()
		return
	}

	err := context.Store(c).Save(
		&record,
	).Error

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		record,
	)
}

// PostBuild creates a new build.
func PostBuild(c *gin.Context) {
	pack := session.Pack(c)

	record := &model.Build{
		PackID: pack.ID,
	}

	record.Defaults()

	if err := c.BindJSON(&record); err != nil {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind build data",
			},
		)

		c.Abort()
		return
	}

	err := context.Store(c).Create(
		&record,
	).Error

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		record,
	)
}

// GetBuildVersions retrieves all versions related to a build.
func GetBuildVersions(c *gin.Context) {

}

// PatchBuildVersion appends a version to a build.
func PatchBuildVersion(c *gin.Context) {

}

// DeleteBuildVersion deleted a version from a build
func DeleteBuildVersion(c *gin.Context) {

}
