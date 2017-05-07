package template

import (
	"html/template"
	"io/ioutil"
	"path"

	"github.com/Sirupsen/logrus"
	"github.com/Unknwon/com"
	"github.com/kleister/kleister-api/config"
)

//go:generate fileb0x ab0x.yaml

// Load initializes the template files.
func Load() *template.Template {
	tpls := template.New("")

	file, err := ReadFile("index.html")

	if err != nil {
		logrus.Warnf("Failed to read builtin index template. %s", err)
	} else {
		tpls.New(
			"index.html",
		).Parse(
			string(file),
		)
	}

	if config.Server.Assets != "" && com.IsDir(config.Server.Assets) {
		customIndex := path.Join(
			config.Server.Assets,
			"index.html",
		)

		if com.IsFile(customIndex) {
			content, err := ioutil.ReadFile(customIndex)

			if err != nil {
				logrus.Warnf("Failed to read custom index template. %s", err)
			} else {
				tpls.New(
					"index.html",
				).Parse(
					string(content),
				)
			}
		}
	}

	return tpls
}
