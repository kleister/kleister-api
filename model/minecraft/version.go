package minecraft

import (
	"github.com/mcuadros/go-version"
)

type Version struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

func (s *Version) Prepare() {

}

func (s *Version) Invalid() bool {
	return (s.Type != "release" && s.Type != "snapshot") || (s.Type == "release" && version.Compare(s.ID, "1.6.4", "<"))
}
