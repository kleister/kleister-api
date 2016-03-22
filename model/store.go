package model

import (
	"github.com/jinzhu/gorm"
)

// Store is a basic struct to represent the database handle.
type Store struct {
	*gorm.DB
}

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

func (s Store) GetClients() (*Clients, error) {
	records := &Clients{}

	err := s.Order(
		"name ASC",
	).Find(
		&records,
	).Error

	return records, err
}

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

func (s Store) GetKeys() (*Keys, error) {
	records := &Keys{}

	err := s.Order(
		"name ASC",
	).Find(
		&records,
	).Error

	return records, err
}

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

func (s Store) GetPacks() (*Packs, error) {
	records := &Packs{}

	err := s.Order(
		"name ASC",
	).Preload(
		"Builds",
	).Preload(
		"Forge",
	).Preload(
		"Minecraft",
	).Find(
		&records,
	).Error

	return records, err
}

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
	).First(
		&record,
	)

	return record, res
}

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
