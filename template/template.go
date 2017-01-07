package template

import (
	"html/template"
)

//go:generate fileb0x ab0x.yaml

// Load initializes the template files.
func Load() *template.Template {
	file, _ := ReadFile("index.html")

	return template.Must(
		template.New(
			"index.html",
		).Parse(
			string(file),
		),
	)
}
