package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/invopop/jsonschema"
	"github.com/kleister/kleister-api/pkg/config"
)

var schemas = map[string]any{
	"config.schema.json": config.Config{},
	"auth.schema.json":   config.AuthConfig{},
}

func main() {
	out := flag.String("out", "dist/schema", "Output directory for schema files")
	flag.Parse()

	if err := os.MkdirAll(*out, 0o755); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for name, target := range schemas {
		reflector := &jsonschema.Reflector{
			FieldNameTag:               "mapstructure",
			RequiredFromJSONSchemaTags: true,
			DoNotReference:             true,
		}

		if err := reflector.AddGoComments(
			"github.com/kleister/kleister-api",
			"./pkg/config",
		); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		schema := reflector.Reflect(target)

		content, err := json.MarshalIndent(
			schema,
			"",
			"  ",
		)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		dest := filepath.Join(
			*out,
			name,
		)

		if err := os.WriteFile(
			dest,
			append(content, '\n'),
			0o644,
		); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Println("Generated:", dest)
	}
}
