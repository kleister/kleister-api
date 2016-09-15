package data

import (
	"regexp"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/model"
)

// GetKeys retrieves all available keys from the database.
func (db *data) GetKeys() (*model.Keys, error) {
	records := &model.Keys{}

	err := db.Order(
		"name ASC",
	).Find(
		records,
	).Error

	return records, err
}

// CreateKey creates a new key.
func (db *data) CreateKey(record *model.Key, current *model.User) error {
	return db.Create(
		record,
	).Error
}

// UpdateKey updates a key.
func (db *data) UpdateKey(record *model.Key, current *model.User) error {
	return db.Save(
		record,
	).Error
}

// DeleteKey deletes a key.
func (db *data) DeleteKey(record *model.Key, current *model.User) error {
	return db.Delete(
		record,
	).Error
}

// GetKey retrieves a specific key from the database.
func (db *data) GetKey(id string) (*model.Key, *gorm.DB) {
	var (
		record = &model.Key{}
		query  *gorm.DB
	)

	if match, _ := regexp.MatchString("^([0-9]+)$", id); match {
		val, _ := strconv.ParseInt(id, 10, 64)

		query = db.Where(
			&model.Key{
				ID: val,
			},
		)
	} else {
		query = db.Where(
			&model.Key{
				Slug: id,
			},
		)
	}

	res := query.Model(
		record,
	).First(
		record,
	)

	return record, res
}

// GetKeyByValue retrieves a specific key by value from the database.
func (db *data) GetKeyByValue(id string) (*model.Key, *gorm.DB) {
	record := &model.Key{}

	res := db.Where(
		&model.Key{
			Value: id,
		},
	).Model(
		record,
	).First(
		record,
	)

	return record, res
}
