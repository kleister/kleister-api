package store

import (
	"github.com/jinzhu/gorm"
	"github.com/kleister/kleister-api/model"
	"golang.org/x/net/context"
)

// GetKeys retrieves all available keys from the database.
func GetKeys(c context.Context) (*model.Keys, error) {
	return FromContext(c).GetKeys()
}

// CreateKey creates a new key.
func CreateKey(c context.Context, record *model.Key) error {
	return FromContext(c).CreateKey(record, Current(c))
}

// UpdateKey updates a key.
func UpdateKey(c context.Context, record *model.Key) error {
	return FromContext(c).UpdateKey(record, Current(c))
}

// DeleteKey deletes a key.
func DeleteKey(c context.Context, record *model.Key) error {
	return FromContext(c).DeleteKey(record, Current(c))
}

// GetKey retrieves a specific key from the database.
func GetKey(c context.Context, id string) (*model.Key, *gorm.DB) {
	return FromContext(c).GetKey(id)
}

// GetKeyByValue retrieves a specific key by value from the database.
func GetKeyByValue(c context.Context, id string) (*model.Key, *gorm.DB) {
	return FromContext(c).GetKeyByValue(id)
}
