package neoforge

import (
	"errors"

	"github.com/rs/zerolog/log"
)

var (
	// ErrSyncUnavailable defines the error of the versions definition is unavailable.
	ErrSyncUnavailable = errors.New("fabric version service is unavailable")
)

// FetchRemote is just a wrapper to get a syncable list of versions.
func FetchRemote() (Versions, error) {
	result, err := FromDefault()

	if err != nil {
		log.Error().
			Err(err).
			Str("service", "neoforge").
			Str("method", "fetch").
			Msg("Failed to sync versions")

		return nil, ErrSyncUnavailable
	}

	ByVersion(
		result.Versions,
	).Sort()

	return result.Versions, nil
}
