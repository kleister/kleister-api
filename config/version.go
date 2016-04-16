package config

import (
	"fmt"
)

const (
	// VersionMajor is the current major version
	VersionMajor = 0

	// VersionMinor is the current minor version
	VersionMinor = 0

	// VersionPatch is the current patch version
	VersionPatch = 0

	// VersionDev indicates the current commit
	VersionDev = "dev"
)

// Version is the version of the current implementation.
var Version = fmt.Sprintf(
	"%d.%d.%d+%s",
	VersionMajor,
	VersionMinor,
	VersionPatch,
	VersionDev,
)

// StrippedVersion is the version without the commit SHA.
var StrippedVersion = fmt.Sprintf(
	"%d.%d.%d",
	VersionMajor,
	VersionMinor,
	VersionPatch,
)
