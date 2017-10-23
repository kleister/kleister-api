package template

import (
	"html/template"
	"io/ioutil"
	"os"
	"path"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/kleister/kleister-api/pkg/config"
)

//go:generate retool -tool-dir ../../_tools do fileb0x ab0x.yaml

// Load initializes the template files.
func Load(logger log.Logger) *template.Template {
	file, err := ReadFile("index.html")

	if err != nil {
		level.Warn(logger).Log(
			"msg", "failed to read builtin template",
			"err", err,
		)
	}

	tpls, err := template.New(
		"index.html",
	).Parse(
		string(file),
	)

	if err != nil {
		level.Warn(logger).Log(
			"msg", "failed to parse builtin template",
			"err", err,
		)
	}

	if config.Server.Assets != "" {
		if stat, err := os.Stat(config.Server.Assets); err == nil && stat.IsDir() {
			customIndex := path.Join(
				config.Server.Assets,
				"index.html",
			)

			if _, err := os.Stat(customIndex); !os.IsNotExist(err) {
				content, err := ioutil.ReadFile(customIndex)

				if err != nil {
					level.Warn(logger).Log(
						"msg", "failed to parse custom template",
						"err", err,
					)
				} else {
					tpls.New(
						"index.html",
					).Parse(
						string(content),
					)
				}
			}
		} else {
			level.Warn(logger).Log(
				"msg", "custom assets directory doesn't exist",
			)
		}
	}

	return tpls
}
