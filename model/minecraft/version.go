package minecraft

type Version struct {
	ID   string
	Type string
	URL  string
}

func (s *Version) Invalid() bool {
	return s.Type != "release" && s.Type != "snapshot"
}
