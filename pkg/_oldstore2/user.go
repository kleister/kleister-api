package storage

import (
	// "fmt"
	// "regexp"
	// "strconv"

	// "github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/pkg/model"
)

// GetUser retrieves a specific user from the database.
func (db *data) GetUser(id string) (*model.User, error) {
	// var (
	// 	record = &model.User{}
	// 	query  *gorm.DB
	// )

	// if match, _ := regexp.MatchString("^([0-9]+)$", id); match {
	// 	val, _ := strconv.ParseInt(id, 10, 64)

	// 	query = db.Where(
	// 		&model.User{
	// 			ID: val,
	// 		},
	// 	)
	// } else {
	// 	query = db.Where(
	// 		&model.User{
	// 			Slug: id,
	// 		},
	// 	)
	// }

	// res := query.Model(
	// 	record,
	// ).Preload(
	// 	"Teams",
	// ).Preload(
	// 	"Packs",
	// ).Preload(
	// 	"Mods",
	// ).First(
	// 	record,
	// )

	return &model.User{}, nil
}
