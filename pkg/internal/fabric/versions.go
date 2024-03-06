package fabric

import (
	"sort"

	"github.com/kleister/kleister-api/pkg/internal/client/maven"
	"github.com/mcuadros/go-version"
)

const (
	// DefaultURL defines the default Forge version URL.
	DefaultURL = "https://maven.fabricmc.net/net/fabricmc/fabric-loader/maven-metadata.xml"
)

// FromDefault is a simply wrapper that loads the default URL.
func FromDefault() (Response, error) {
	fetched, err := maven.Get(
		DefaultURL,
	)

	if err != nil {
		return Response{}, err
	}

	versions := Versions{}

	for _, version := range fetched.Versions {
		versions = append(
			versions,
			Version{
				Value: version,
			},
		)
	}

	return Response{
		Versions: versions,
	}, nil
}

// Response simply defines the result with all versions.
type Response struct {
	Versions Versions
}

// Version represents a single version of Fabric.
type Version struct {
	Value string
}

// Versions is a simple slice of available versions.
type Versions []Version

// ByVersion sorts a list of versions by ID.
type ByVersion Versions

// Sort simply sorts the versions list.
func (b ByVersion) Sort() {
	sort.Sort(b)
}

// Len is part of the sorting algorithm.
func (b ByVersion) Len() int {
	return len(b)
}

// Swap is part of the sorting algorithm.
func (b ByVersion) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// Less is part of the sorting algorithm.
func (b ByVersion) Less(i, j int) bool {
	cmp := version.CompareSimple(
		version.Normalize(b[i].Value),
		version.Normalize(b[j].Value),
	)

	if cmp == 0 {
		return b[i].Value < b[j].Value
	}

	return cmp < 0
}
