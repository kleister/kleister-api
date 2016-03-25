package model

import (
	"github.com/jinzhu/gorm"
)

// Store is a basic struct to represent the database handle.
type Store struct {
	*gorm.DB
}

// GetBuilds retrieves all available builds from the database.
func (s Store) GetBuilds(pack int) (*Builds, error) {
	records := &Builds{}

	err := s.Order(
		"name ASC",
	).Where(
		"pack_id = ?",
		pack,
	).Preload(
		"Pack",
	).Preload(
		"Minecraft",
	).Preload(
		"Forge",
	).Find(
		&records,
	).Error

	return records, err
}

// GetBuild retrieves a specific build from the database.
func (s Store) GetBuild(pack int, id string) (*Build, *gorm.DB) {
	record := &Build{}

	res := s.Where(
		"pack_id = ?",
		pack,
	).Where(
		"id = ?",
		id,
	).Or(
		"slug = ?",
		id,
	).Model(
		&record,
	).Preload(
		"Pack",
	).Preload(
		"Minecraft",
	).Preload(
		"Forge",
	).First(
		&record,
	)

	return record, res
}

// GetClients retrieves all available clients from the database.
func (s Store) GetClients() (*Clients, error) {
	records := &Clients{}

	err := s.Order(
		"name ASC",
	).Find(
		&records,
	).Error

	return records, err
}

// GetClient retrieves a specific client from the database.
func (s Store) GetClient(id string) (*Client, *gorm.DB) {
	record := &Client{}

	res := s.Where(
		"id = ?",
		id,
	).Or(
		"slug = ?",
		id,
	).First(
		&record,
	)

	return record, res
}

// GetForges retrieves all available Forge versions from the database.
func (s Store) GetForges() (*Forges, error) {
	records := &Forges{}

	err := s.Order(
		"minecraft DESC",
	).Order(
		"name DESC",
	).Find(
		&records,
	).Error

	return records, err
}

// GetForge retrieves a specific Forge version from the database.
func (s Store) GetForge(id string) (*Forge, *gorm.DB) {
	record := &Forge{}

	res := s.Where(
		"id = ?",
		id,
	).Or(
		"slug = ?",
		id,
	).First(
		&record,
	)

	return record, res
}

// GetKeys retrieves all available keys from the database.
func (s Store) GetKeys() (*Keys, error) {
	records := &Keys{}

	err := s.Order(
		"name ASC",
	).Find(
		&records,
	).Error

	return records, err
}

// GetKey retrieves a specific key from the database.
func (s Store) GetKey(id string) (*Key, *gorm.DB) {
	record := &Key{}

	res := s.Where(
		"id = ?",
		id,
	).Or(
		"slug = ?",
		id,
	).First(
		&record,
	)

	return record, res
}

// GetMinecrafts retrieves all available Minecraft versions from the database.
func (s Store) GetMinecrafts() (*Minecrafts, error) {
	records := &Minecrafts{}

	err := s.Order(
		"type DESC",
	).Order(
		"name DESC",
	).Find(
		&records,
	).Error

	return records, err
}

// GetMinecraft retrieves a specific Minecraft version from the database.
func (s Store) GetMinecraft(id string) (*Minecraft, *gorm.DB) {
	record := &Minecraft{}

	res := s.Where(
		"id = ?",
		id,
	).Or(
		"slug = ?",
		id,
	).First(
		&record,
	)

	return record, res
}

// GetMods retrieves all available mods from the database.
func (s Store) GetMods() (*Mods, error) {
	records := &Mods{}

	err := s.Order(
		"name ASC",
	).Preload(
		"Versions",
	).Find(
		&records,
	).Error

	return records, err
}

// GetMod retrieves a specific mod from the database.
func (s Store) GetMod(id string) (*Mod, *gorm.DB) {
	record := &Mod{}

	res := s.Where(
		"id = ?",
		id,
	).Or(
		"slug = ?",
		id,
	).Model(
		&record,
	).Preload(
		"Versions",
	).First(
		&record,
	)

	return record, res
}

// GetPacks retrieves all available packs from the database.
func (s Store) GetPacks() (*Packs, error) {
	records := &Packs{}

	err := s.Order(
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
		&records,
	).Error

	return records, err
}

// GetPack retrieves a specific pack from the database.
func (s Store) GetPack(id string) (*Pack, *gorm.DB) {
	record := &Pack{}

	res := s.Where(
		"packs.id = ?",
		id,
	).Or(
		"packs.slug = ?",
		id,
	).Model(
		&record,
	).Preload(
		"Builds",
	).Preload(
		"Icon",
	).Preload(
		"Background",
	).Preload(
		"Logo",
	).First(
		&record,
	)

	return record, res
}

// GetUsers retrieves all available users from the database.
func (s Store) GetUsers() (*Users, error) {
	records := &Users{}

	err := s.Order(
		"username ASC",
	).Preload(
		"Permission",
	).Find(
		&records,
	).Error

	return records, err
}

// GetUser retrieves a specific user from the database.
func (s Store) GetUser(id string) (*User, *gorm.DB) {
	record := &User{}

	res := s.Where(
		"id = ?",
		id,
	).Or(
		"slug = ?",
		id,
	).Model(
		&record,
	).Preload(
		"Permission",
	).First(
		&record,
	)

	return record, res
}

// GetVersions retrieves all available versions from the database.
func (s Store) GetVersions(mod int) (*Versions, error) {
	records := &Versions{}

	err := s.Order(
		"name ASC",
	).Where(
		"mod_id = ?",
		mod,
	).Preload(
		"Mod",
	).Preload(
		"File",
	).Find(
		&records,
	).Error

	return records, err
}

// GetVersion retrieves a specific version from the database.
func (s Store) GetVersion(mod int, id string) (*Version, *gorm.DB) {
	record := &Version{}

	res := s.Where(
		"mod_id = ?",
		mod,
	).Where(
		"id = ?",
		id,
	).Or(
		"slug = ?",
		id,
	).Model(
		&record,
	).Preload(
		"Mod",
	).Preload(
		"File",
	).First(
		&record,
	)

	return record, res
}
