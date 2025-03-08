package minecraft

import (
	"errors"

	"github.com/kleister/go-minecraft/version"
	"github.com/rs/zerolog/log"
)

var (
	// ErrSyncUnavailable defines the error of the versions definition is unavailable.
	ErrSyncUnavailable = errors.New("fabric version service is unavailable")
)

// FetchRemote is just a wrapper to get a syncable list of versions.
func FetchRemote() (version.Versions, error) {
	result, err := version.FromDefault()

	if err != nil {
		log.Error().
			Err(err).
			Str("service", "minecraft").
			Str("method", "fetch").
			Msg("Failed to fetch versions")

		return nil, ErrSyncUnavailable
	}

	version.ByVersion(
		result.Releases,
	).Sort()

	return result.Releases.Filter(
		&version.Filter{
			Version: ">=1.7.10",
		},
	), nil
}
