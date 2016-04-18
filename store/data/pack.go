package data

import (
	"github.com/jinzhu/gorm"
	"github.com/solderapp/solder-api/model"
)

// GetPacks retrieves all available packs from the database.
func (db *data) GetPacks() (*model.Packs, error) {
	records := &model.Packs{}

	err := db.Order(
		"name ASC",
	).Preload(
		"Builds",
	).Preload(
		"Icon",
	).Preload(
		"Background",
	).Preload(
		"Logo",
	).Preload(
		"Builds.Forge",
	).Preload(
		"Builds.Minecraft",
	).Find(
		records,
	).Error

	return records, err
}

// CreatePack creates a new pack.
func (db *data) CreatePack(record *model.Pack) error {
	return db.Create(
		record,
	).Error
}

// UpdatePack updates a pack.
func (db *data) UpdatePack(record *model.Pack) error {
	return db.Save(
		record,
	).Error
}

// DeletePack deletes a pack.
func (db *data) DeletePack(record *model.Pack) error {

	// tx := context.Store(c).Begin()
	// failed := false

	// if record.Icon != nil {
	// 	err := tx.Delete(
	// 		&record.Icon,
	// 	).Error

	// 	if err != nil {
	// 		failed = true
	// 	}
	// }

	// if record.Background != nil {
	// 	err := tx.Delete(
	// 		&record.Background,
	// 	).Error

	// 	if err != nil {
	// 		failed = true
	// 	}
	// }

	// if record.Logo != nil {
	// 	err := tx.Delete(
	// 		&record.Logo,
	// 	).Error

	// 	if err != nil {
	// 		failed = true
	// 	}
	// }

	// err := tx.Delete(
	// 	&record,
	// ).Error

	// if failed || err != nil {
	// 	tx.Rollback()

	// 	c.JSON(
	// 		http.StatusBadRequest,
	// 		gin.H{
	// 			"status":  http.StatusBadRequest,
	// 			"message": "Failed to delete pack",
	// 		},
	// 	)

	// 	c.Abort()
	// 	return
	// }

	// tx.Commit()

	return db.Delete(
		record,
	).Error
}

// GetPack retrieves a specific pack from the database.
func (db *data) GetPack(id string) (*model.Pack, *gorm.DB) {
	record := &model.Pack{}

	res := db.Where(
		"packs.id = ?",
		id,
	).Or(
		"packs.slug = ?",
		id,
	).Model(
		record,
	).Preload(
		"Builds",
	).Preload(
		"Icon",
	).Preload(
		"Background",
	).Preload(
		"Logo",
	).First(
		record,
	)

	return record, res
}

// GetPackClients retrieves clients for a pack.
func (db *data) GetPackClients(id int) (*model.Clients, error) {
	records := &model.Clients{}

	err := db.Model(
		&model.Pack{
			ID: id,
		},
	).Association(
		"Clients",
	).Find(
		records,
	).Error

	return records, err
}

// GetPackHasClient checks if a specific client is assigned to a pack.
func (db *data) GetPackHasClient(parent, id int) bool {
	record := &model.Client{
		ID: id,
	}

	count := db.Model(
		&model.Pack{
			ID: parent,
		},
	).Association(
		"Clients",
	).Find(
		record,
	).Count()

	return count > 0
}
