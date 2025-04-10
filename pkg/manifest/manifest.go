package manifest

import (
	"encoding/json"
	"fmt"

	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/frontend"
)

// Document defines a single file within the frontend manifest.
type Document struct {
	File        string   `json:"file"`
	Name        string   `json:"name"`
	Source      string   `json:"src"`
	Entry       bool     `json:"isEntry"`
	Stylehseets []string `json:"css"`
}

// Manifest defines the map of documents within the frontend manifest.
type Manifest map[string]Document

// Index is just a wrapper to always return the index.html document.
func (m Manifest) Index() Document {
	if val, ok := m["index.html"]; ok {
		return val
	}

	return Document{}
}

// Read simply reads or initializes the frontend manifest.
func Read(cfg *config.Config) (Manifest, error) {
	file, err := frontend.Load(cfg).Open("manifest.json")

	if err != nil {
		return nil, fmt.Errorf("failed to read manifest: %w", err)
	}

	defer func() { _ = file.Close() }()
	result := Manifest{}

	if err := json.NewDecoder(file).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to parse manifest: %w", err)
	}

	return result, nil
}
