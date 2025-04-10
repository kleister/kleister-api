package maven

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
)

// Get simply fetches and parses a maven metadata file.
func Get(from string) (*Maven, error) {
	result := &Maven{}

	parsedFrom, err := url.Parse(
		from,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to parse url: %w", err)
	}

	resp, err := http.Get(
		parsedFrom.String(),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch content: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	decoder := xml.NewDecoder(
		resp.Body,
	)

	if err := decoder.Decode(
		result,
	); err != nil {
		return nil, fmt.Errorf("failed to parse content: %w", err)
	}

	return result, nil
}

// Maven defines the metadata file from a maven server.
type Maven struct {
	Group    string   `xml:"groupId"`
	Artifact string   `xml:"artifactId"`
	Latest   string   `xml:"versioning>latest"`
	Release  string   `xml:"versioning>release"`
	Versions []string `xml:"versioning>versions>version"`
}
