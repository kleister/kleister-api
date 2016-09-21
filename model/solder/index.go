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
						result.Modpacks[record.Slug] = getPackValue(record, client, key, include)
						break
					}
				}
			case key != nil:
				result.Modpacks[record.Slug] = getPackValue(record, client, key, include)
			}
		} else {
			result.Modpacks[record.Slug] = getPackValue(record, client, key, include)
		}
	}

	return result
}

func getPackValue(record *model.Pack, client *model.Client, key *model.Key, include string) interface{} {
	switch include {
	case "full":
		return NewPackFromModel(record, client, key, include)
	default:
		return record.Name
	}
}
