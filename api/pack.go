package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/kleister/kleister-api/model"
	"github.com/kleister/kleister-api/router/middleware/location"
	"github.com/kleister/kleister-api/router/middleware/session"
	"github.com/kleister/kleister-api/store"
)

// PackIndex retrieves all available packs.
func PackIndex(c *gin.Context) {
	location := location.Location(c)

	records, err := store.GetPacks(
		c,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to fetch packs",
			},
		)

		c.Abort()
		return
	}

	for _, record := range *records {
		if record.Icon != nil {
			record.Icon.URL = strings.Join(
				[]string{
					location.String(),
					"storage",
					"icon",
					strconv.Itoa(record.ID),
				},
				"/",
			)
		}

		if record.Background != nil {
			record.Background.URL = strings.Join(
				[]string{
					location.String(),
					"storage",
					"background",
					strconv.Itoa(record.ID),
				},
				"/",
			)
		}

		if record.Logo != nil {
			record.Logo.URL = strings.Join(
				[]string{
					location.String(),
					"storage",
					"logo",
					strconv.Itoa(record.ID),
				},
				"/",
			)
		}
	}

	c.JSON(
		http.StatusOK,
		records,
	)
}

// PackShow retrieves a specific pack.
func PackShow(c *gin.Context) {
	location := location.Location(c)
	record := session.Pack(c)

	if record.Icon != nil {
		record.Icon.URL = strings.Join(
			[]string{
				location.String(),
				"storage",
				"icon",
				strconv.Itoa(record.ID),
			},
			"/",
		)
	}

	if record.Background != nil {
		record.Background.URL = strings.Join(
			[]string{
				location.String(),
				"storage",
				"background",
				strconv.Itoa(record.ID),
			},
			"/",
		)
	}

	if record.Logo != nil {
		record.Logo.URL = strings.Join(
			[]string{
				location.String(),
				"storage",
				"logo",
				strconv.Itoa(record.ID),
			},
			"/",
		)
	}

	c.JSON(
		http.StatusOK,
		record,
	)
}

// PackDelete removes a specific pack.
func PackDelete(c *gin.Context) {
	record := session.Pack(c)

	err := store.DeletePack(
		c,
		record,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Failed to delete pack",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully deleted pack",
		},
	)
}

// PackUpdate updates an existing pack.
func PackUpdate(c *gin.Context) {
	record := session.Pack(c)

	if err := c.BindJSON(&record); err != nil {
		logrus.Warn("Failed to bind pack data")
		logrus.Warn(err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind pack data",
			},
		)

		c.Abort()
		return
	}

	err := store.UpdatePack(
		c,
		record,
	)

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

// PackCreate creates a new pack.
func PackCreate(c *gin.Context) {
	record := &model.Pack{}

	if err := c.BindJSON(&record); err != nil {
		logrus.Warn("Failed to bind pack data")
		logrus.Warn(err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind pack data",
			},
		)

		c.Abort()
		return
	}

	err := store.CreatePack(
		c,
		record,
	)

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

// PackClientIndex retrieves all clients related to a pack.
func PackClientIndex(c *gin.Context) {
	records, err := store.GetPackClients(
		c,
		&model.PackClientParams{
			Pack: c.Param("pack"),
		},
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to fetch clients",
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

// PackClientAppend appends a client to a pack.
func PackClientAppend(c *gin.Context) {
	form := &model.PackClientParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warn("Failed to bind post data")
		logrus.Warn(err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind form data",
			},
		)

		c.Abort()
		return
	}

	form.Pack = c.Param("pack")

	assigned := store.GetPackHasClient(
		c,
		form,
	)

	if assigned == true {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Client is already appended",
			},
		)

		c.Abort()
		return
	}

	err := store.CreatePackClient(
		c,
		form,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to append client",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully appended client",
		},
	)
}

// PackClientDelete deleted a client from a pack
func PackClientDelete(c *gin.Context) {
	form := &model.PackClientParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warn("Failed to bind post data")
		logrus.Warn(err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind form data",
			},
		)

		c.Abort()
		return
	}

	form.Pack = c.Param("pack")

	assigned := store.GetPackHasClient(
		c,
		form,
	)

	if assigned == false {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Client is not assigned",
			},
		)

		c.Abort()
		return
	}

	err := store.DeletePackClient(
		c,
		form,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to unlink client",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully unlinked client",
		},
	)
}

// PackUserIndex retrieves all users related to a pack.
func PackUserIndex(c *gin.Context) {
	records, err := store.GetPackUsers(
		c,
		&model.PackUserParams{
			Pack: c.Param("pack"),
		},
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to fetch users",
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

// PackUserAppend appends a user to a pack.
func PackUserAppend(c *gin.Context) {
	form := &model.PackUserParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warn("Failed to bind post data")
		logrus.Warn(err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind form data",
			},
		)

		c.Abort()
		return
	}

	form.Pack = c.Param("pack")

	assigned := store.GetPackHasUser(
		c,
		form,
	)

	if assigned == true {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "User is already appended",
			},
		)

		c.Abort()
		return
	}

	err := store.CreatePackUser(
		c,
		form,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to append user",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully appended user",
		},
	)
}

// PackUserDelete deleted a user from a pack
func PackUserDelete(c *gin.Context) {
	form := &model.PackUserParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warn("Failed to bind post data")
		logrus.Warn(err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind form data",
			},
		)

		c.Abort()
		return
	}

	form.Pack = c.Param("pack")

	assigned := store.GetPackHasUser(
		c,
		form,
	)

	if assigned == false {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "User is not assigned",
			},
		)

		c.Abort()
		return
	}

	err := store.DeletePackUser(
		c,
		form,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to unlink user",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully unlinked user",
		},
	)
}
