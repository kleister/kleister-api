package user

import (
	"github.com/kleister/kleister-api/pkg/storage"
)

type Service interface {
	// ListUser
	ListUser()

	// CreateUser
	CreateUser()

	// UpdateUser
	UpdateUser()

	// DeleteUser
	DeleteUser()

	// ShowUser
	ShowUser()
}

type ServiceOptions struct {
	Store storage.Store
}

func NewService(opts ServiceOptions) Service {
	return &service{
		store: opts.Store,
	}
}

type service struct {
	store storage.Store
}

func (s *service) ListUser() {
	// records, err := store.GetUsers(
	// 	c,
	// )

	// if err != nil {
	// 	logrus.Warnf("Failed to fetch users. %s", err)

	// 	c.JSON(
	// 		http.StatusInternalServerError,
	// 		gin.H{
	// 			"status":  http.StatusInternalServerError,
	// 			"message": "Failed to fetch users",
	// 		},
	// 	)

	// 	c.Abort()
	// 	return
	// }

	// c.JSON(
	// 	http.StatusOK,
	// 	records,
	// )
}

func (s *service) CreateUser() {
	// record := &model.User{}

	// if err := c.BindJSON(&record); err != nil {
	// 	logrus.Warnf("Failed to bind user data. %s", err)

	// 	c.JSON(
	// 		http.StatusPreconditionFailed,
	// 		gin.H{
	// 			"status":  http.StatusPreconditionFailed,
	// 			"message": "Failed to bind user data",
	// 		},
	// 	)

	// 	c.Abort()
	// 	return
	// }

	// err := store.CreateUser(
	// 	c,
	// 	record,
	// )

	// if err != nil {
	// 	logrus.Warnf("Failed to create user. %s", err)

	// 	c.JSON(
	// 		http.StatusBadRequest,
	// 		gin.H{
	// 			"status":  http.StatusBadRequest,
	// 			"message": err.Error(),
	// 		},
	// 	)

	// 	c.Abort()
	// 	return
	// }

	// c.JSON(
	// 	http.StatusOK,
	// 	record,
	// )
}

func (s *service) UpdateUser() {
	// record := session.User(c)

	// if err := c.BindJSON(&record); err != nil {
	// 	logrus.Warnf("Failed to bind user data. %s", err)

	// 	c.JSON(
	// 		http.StatusPreconditionFailed,
	// 		gin.H{
	// 			"status":  http.StatusPreconditionFailed,
	// 			"message": "Failed to bind user data",
	// 		},
	// 	)

	// 	c.Abort()
	// 	return
	// }

	// err := store.UpdateUser(
	// 	c,
	// 	record,
	// )

	// if err != nil {
	// 	logrus.Warnf("Failed to update user. %s", err)

	// 	c.JSON(
	// 		http.StatusBadRequest,
	// 		gin.H{
	// 			"status":  http.StatusBadRequest,
	// 			"message": err.Error(),
	// 		},
	// 	)

	// 	c.Abort()
	// 	return
	// }

	// c.JSON(
	// 	http.StatusOK,
	// 	record,
	// )
}

func (s *service) DeleteUser() {
	// record := session.User(c)

	// err := store.DeleteUser(
	// 	c,
	// 	record,
	// )

	// if err != nil {
	// 	logrus.Warnf("Failed to delete user. %s", err)

	// 	c.JSON(
	// 		http.StatusBadRequest,
	// 		gin.H{
	// 			"status":  http.StatusBadRequest,
	// 			"message": "Failed to delete user",
	// 		},
	// 	)

	// 	c.Abort()
	// 	return
	// }

	// c.JSON(
	// 	http.StatusOK,
	// 	gin.H{
	// 		"status":  http.StatusOK,
	// 		"message": "Successfully deleted user",
	// 	},
	// )
}

func (s *service) ShowUser() {
	// record := session.User(c)

	// c.JSON(
	// 	http.StatusOK,
	// 	record,
	// )
}
