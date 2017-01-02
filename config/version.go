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

	// VersionDate indicates the build date
	VersionDate = "20170101"
)

var (
	// Version is the version of the current implementation.
	Version = fmt.Sprintf(
		"%d.%d.%d+git%s.%s",
		VersionMajor,
		VersionMinor,
		VersionPatch,
		VersionDate,
		VersionDev,
	)
)
