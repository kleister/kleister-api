package assets

import (
	"net/http"
	"os"
	"path"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/kleister/kleister-api/pkg/config"
)

//go:generate retool -tool-dir ../../_tools do fileb0x ab0x.yaml

// Load initializes the static files.
func Load(logger log.Logger) http.FileSystem {
	return ChainedFS{
		logger: logger,
	}
}

// ChainedFS is a simple HTTP filesystem including custom path.
type ChainedFS struct {
	logger log.Logger
}

// Open just implements the HTTP filesystem interface.
func (c ChainedFS) Open(origPath string) (http.File, error) {
	if config.Server.Assets != "" {
		if stat, err := os.Stat(config.Server.Assets); err == nil && stat.IsDir() {
			customPath := path.Join(
				config.Server.Assets,
				origPath,
			)

			if _, err := os.Stat(customPath); !os.IsNotExist(err) {
				f, err := os.Open(customPath)

				if err != nil {
					return nil, err
				}

				return f, nil
			}
		} else {
			level.Warn(c.logger).Log(
				"msg", "custom assets directory doesn't exist",
			)
		}
	}

	f, err := FS.OpenFile(CTX, origPath, os.O_RDONLY, 0644)

	if err != nil {
		return nil, err
	}

	return f, nil
}
