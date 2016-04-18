package config

import (
	"fmt"
)

var (
	// VersionMajor is the current major version
	VersionMajor = 0

	// VersionMinor is the current minor version
	VersionMinor = 0

	// VersionPatch is the current patch version
	VersionPatch = 0

	// VersionDev indicates the current commit
	VersionDev = "dev"

	// Version is the version of the current implementation.
	Version = fmt.Sprintf(
		"%d.%d.%d+%s",
		VersionMajor,
		VersionMinor,
		VersionPatch,
		VersionDev,
	)
)
