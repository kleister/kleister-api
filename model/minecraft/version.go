package minecraft

import (
	"github.com/mcuadros/go-version"
)

// Version represents a row within the remote version info.
type Version struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// Invalid validates the current row of the remote info.
func (s *Version) Invalid() bool {
	return (s.Type != "release" && s.Type != "snapshot") || (s.Type == "release" && version.Compare(s.ID, "1.6.4", "<"))
}
