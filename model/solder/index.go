package solder

import (
	"github.com/kleister/kleister-api/model"
)

// Packs is simply a collection of pack structs.
type Packs struct {
	Modpacks map[string]interface{} `json:"modpacks"`
}

// NewPacksFromList generates a solder model from our used models.
func NewPacksFromList(records *model.Packs, client *model.Client, key *model.Key, include string) *Packs {
	result := &Packs{
		Modpacks: make(map[string]interface{}, 0),
	}

	for _, record := range *records {
		if record.Private || record.Hidden {
			switch {
			case client != nil:
				for _, pack := range client.Packs {
					if pack.ID == record.ID {
						result.Modpacks[record.Slug] = getPackValue(include, record)
						break
					}
				}
			case key != nil:
				result.Modpacks[record.Slug] = getPackValue(include, record)
			}
		} else {
			result.Modpacks[record.Slug] = getPackValue(include, record)
		}
	}

	return result
}

func getPackValue(include string, record *model.Pack) interface{} {
	switch include {
	case "full":
		return record
	default:
		return record.Name
	}
}
