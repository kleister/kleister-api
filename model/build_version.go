package model

// BuildVersions is simply a collection of build version structs.
type BuildVersions []*BuildVersion

// BuildVersion represents a build version model definition.
type BuildVersion struct {
	BuildID   int      `json:"build_id" sql:"index"`
	Build     *Build   `json:"build,omitempty"`
	VersionID int      `json:"version_id" sql:"index"`
	Version   *Version `json:"version,omitempty"`
}
