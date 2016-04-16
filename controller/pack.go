package controller

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"
	"github.com/solderapp/solder-api/router/middleware/context"
	"github.com/solderapp/solder-api/router/middleware/session"
)

// GetPacks retrieves all available packs.
func GetPacks(c *gin.Context) {
	root := context.Root(c)
	records, err := context.Store(c).GetPacks()

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
					root,
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
					root,
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
					root,
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

// GetPack retrieves a specific pack.
func GetPack(c *gin.Context) {
	root := context.Root(c)
	record := session.Pack(c)

	if record.Icon != nil {
		record.Icon.URL = strings.Join(
			[]string{
				root,
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
				root,
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
				root,
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

// DeletePack removes a specific pack.
func DeletePack(c *gin.Context) {
	config := context.Config(c)
	record := session.Pack(c)

	tx := context.Store(c).Begin()
	failed := false

	if record.Icon != nil {
		record.Icon.Path = config.Server.Storage

		err := tx.Delete(
			&record.Icon,
		).Error

		if err != nil {
			failed = true
		}
	}

	if record.Background != nil {
		record.Background.Path = config.Server.Storage

		err := tx.Delete(
			&record.Background,
		).Error

		if err != nil {
			failed = true
		}
	}

	if record.Logo != nil {
		record.Logo.Path = config.Server.Storage

		err := tx.Delete(
			&record.Logo,
		).Error

		if err != nil {
			failed = true
		}
	}

	err := tx.Delete(
		&record,
	).Error

	if failed || err != nil {
		tx.Rollback()

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

	tx.Commit()

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully deleted pack",
		},
	)
}

// PatchPack updates an existing pack.
func PatchPack(c *gin.Context) {
	config := context.Config(c)
	record := session.Pack(c)

	if err := c.BindJSON(&record); err != nil {
		logrus.Warn("Failed to bind pack data")
		logrus.Warn(err)

		b, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(
				http.StatusPreconditionFailed,
				gin.H{
					"status":  http.StatusPreconditionFailed,
					"message": err,
				},
			)

			c.Abort()
			return
		}

		logrus.Warnf("%s", b)

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

	if record.Icon != nil {
		record.Icon.Path = config.Server.Storage
	}

	if record.Background != nil {
		record.Background.Path = config.Server.Storage
	}

	if record.Logo != nil {
		record.Logo.Path = config.Server.Storage
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

// PostPack creates a new pack.
func PostPack(c *gin.Context) {
	config := context.Config(c)
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

	if record.Icon != nil {
		record.Icon.Path = config.Server.Storage
	}

	if record.Background != nil {
		record.Background.Path = config.Server.Storage
	}

	if record.Logo != nil {
		record.Logo.Path = config.Server.Storage
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

// GetPackClients retrieves all clients related to a pack.
func GetPackClients(c *gin.Context) {
	pack := session.Pack(c)
	records := &model.Clients{}

	err := context.Store(c).Model(
		&pack,
	).Association(
		"Clients",
	).Find(
		&records,
	).Error

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

// PatchPackClient appends a client to a pack.
func PatchPackClient(c *gin.Context) {
	pack := session.Pack(c)
	client := session.Client(c)

	count := context.Store(c).Model(
		&pack,
	).Association(
		"Clients",
	).Find(
		&client,
	).Count()

	if count > 0 {
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

	err := context.Store(c).Model(
		&pack,
	).Association(
		"Clients",
	).Append(
		&client,
	).Error

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

// DeletePackClient deleted a client from a pack
func DeletePackClient(c *gin.Context) {
	pack := session.Pack(c)
	client := session.Client(c)

	count := context.Store(c).Model(
		&pack,
	).Association(
		"Clients",
	).Find(
		&client,
	).Count()

	if count < 1 {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "Client is not assigned",
			},
		)

		c.Abort()
		return
	}

	err := context.Store(c).Model(
		&pack,
	).Association(
		"Clients",
	).Delete(
		&client,
	).Error

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
