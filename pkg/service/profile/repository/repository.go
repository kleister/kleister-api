package repository

import (
	"errors"
)

var (
	// ErrProfileNotFound defines the error if a profile could not be found.
	ErrProfileNotFound = errors.New("profile not found")
)

// ProfileRepository defines the required functions for the repository.
type ProfileRepository interface {
}
