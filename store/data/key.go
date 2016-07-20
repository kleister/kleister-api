package data

import (
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
func (db *data) CreateKey(record *model.Key) error {
	return db.Create(
		record,
	).Error
}

// UpdateKey updates a key.
func (db *data) UpdateKey(record *model.Key) error {
	return db.Save(
		record,
	).Error
}

// DeleteKey deletes a key.
func (db *data) DeleteKey(record *model.Key) error {
	return db.Delete(
		record,
	).Error
}

// GetKey retrieves a specific key from the database.
func (db *data) GetKey(id string) (*model.Key, *gorm.DB) {
	record := &model.Key{}

	res := db.Where(
		"id = ?",
		id,
	).Or(
		"slug = ?",
		id,
	).First(
		record,
	)

	return record, res
}
