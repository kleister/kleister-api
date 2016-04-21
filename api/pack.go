package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"
	"github.com/solderapp/solder-api/router/middleware/location"
	"github.com/solderapp/solder-api/router/middleware/session"
	"github.com/solderapp/solder-api/store"
)

// GetPacks retrieves all available packs.
func GetPacks(c *gin.Context) {
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

// GetPack retrieves a specific pack.
func GetPack(c *gin.Context) {
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

// DeletePack removes a specific pack.
func DeletePack(c *gin.Context) {
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

// PatchPack updates an existing pack.
func PatchPack(c *gin.Context) {
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

// PostPack creates a new pack.
func PostPack(c *gin.Context) {
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

// GetPackClients retrieves all clients related to a pack.
func GetPackClients(c *gin.Context) {
	pack := session.Pack(c)

	records, err := store.GetPackClients(
		c,
		pack.ID,
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

// PatchPackClient appends a client to a pack.
func PatchPackClient(c *gin.Context) {
	pack := session.Pack(c)
	client := session.Client(c)

	assigned := store.GetPackHasClient(
		c,
		pack.ID,
		client.ID,
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
		pack.ID,
		client.ID,
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

// DeletePackClient deleted a client from a pack
func DeletePackClient(c *gin.Context) {
	pack := session.Pack(c)
	client := session.Client(c)

	assigned := store.GetPackHasClient(
		c,
		pack.ID,
		client.ID,
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
		pack.ID,
		client.ID,
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

// GetPackUsers retrieves all users related to a pack.
func GetPackUsers(c *gin.Context) {
	pack := session.Pack(c)

	records, err := store.GetPackUsers(
		c,
		pack.ID,
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

// PatchPackUser appends a user to a pack.
func PatchPackUser(c *gin.Context) {
	pack := session.Pack(c)
	user := session.User(c)

	assigned := store.GetPackHasUser(
		c,
		pack.ID,
		user.ID,
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
		pack.ID,
		user.ID,
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

// DeletePackUser deleted a user from a pack
func DeletePackUser(c *gin.Context) {
	pack := session.Pack(c)
	user := session.User(c)

	assigned := store.GetPackHasUser(
		c,
		pack.ID,
		user.ID,
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
		pack.ID,
		user.ID,
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
