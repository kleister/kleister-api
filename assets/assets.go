package assets

import (
	"net/http"
)

//go:generate fileb0x ab0x.yaml

// Load initializes the static files.
func Load() http.FileSystem {
	return HTTP
}
