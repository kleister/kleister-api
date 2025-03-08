package store

import (
	"errors"
)

var (
	// ErrWrongCredentials is returned when credentials are wrong.
	ErrWrongCredentials = errors.New("wrong credentials provided")

	// ErrAlreadyAssigned defines the error if relation is already assigned.
	ErrAlreadyAssigned = errors.New("user pack already exists")

	// ErrNotAssigned defines the error if relation is not assigned.
	ErrNotAssigned = errors.New("user pack is not defined")

	// ErrMinecraftNotFound is returned when a minecraft was not found.
	ErrMinecraftNotFound = errors.New("minecraft not found")

	// ErrForgeNotFound is returned when a forge was not found.
	ErrForgeNotFound = errors.New("forge not found")

	// ErrNeoforgeNotFound is returned when a neoforge was not found.
	ErrNeoforgeNotFound = errors.New("neoforge not found")

	// ErrQuiltNotFound is returned when a quilt was not found.
	ErrQuiltNotFound = errors.New("quilt not found")

	// ErrFabricNotFound is returned when a fabric was not found.
	ErrFabricNotFound = errors.New("fabric not found")

	// ErrModNotFound is returned when a mod was not found.
	ErrModNotFound = errors.New("mod not found")

	// ErrVersionNotFound is returned when a version was not found.
	ErrVersionNotFound = errors.New("version not found")

	// ErrPackNotFound is returned when a pack was not found.
	ErrPackNotFound = errors.New("pack not found")

	// ErrBuildNotFound is returned when a build was not found.
	ErrBuildNotFound = errors.New("build not found")

	// ErrGroupNotFound is returned when a user was not found.
	ErrGroupNotFound = errors.New("group not found")

	// ErrUserNotFound is returned when a user was not found.
	ErrUserNotFound = errors.New("user not found")

	// ErrTokenNotFound is returned when a token was not found.
	ErrTokenNotFound = errors.New("token not found")

	// ErrInvalidUploadEncoding defines the error for invalid upload encodings.
	ErrInvalidUploadEncoding = errors.New("invalid upload encoding")
)
