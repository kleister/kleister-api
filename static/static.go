package static

import (
	"net/http"

	"github.com/solderapp/solder/config"
)

//go:generate esc -o bindata.go -pkg static -prefix files files/

func Load(cfg *config.Config) http.FileSystem {
	return FS(false)
}
