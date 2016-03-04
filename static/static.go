package static

import (
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
)

//go:generate go-bindata -ignore "\\.go" -pkg static -prefix dist -o bindata.go ./dist/...
//go:generate go fmt bindata.go

func Load() http.FileSystem {
	return &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    "",
	}
}
