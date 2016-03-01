package template

import (
	"html/template"

	"github.com/solderapp/solder/config"
)

//go:generate esc -o bindata.go -pkg template -prefix files files/

func Load(cfg *config.Config) *template.Template {
	file := FSMustString(false, "/index.html")

	return template.Must(
		template.New(
			"index.html",
		).Parse(
			string(file),
		),
	)
}
