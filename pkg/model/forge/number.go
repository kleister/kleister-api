package forge

import (
	"github.com/mcuadros/go-version"
)

// Number represents a row within the remote number info.
type Number struct {
	ID        string `json:"version"`
	Minecraft string `json:"mcversion"`
}

// Invalid validates the current row of the remote info.
func (s *Number) Invalid() bool {
	return version.Compare(s.Minecraft, "1.6.4", "<")
}
