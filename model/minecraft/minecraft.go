package minecraft

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	MINECRAFT_VERSIONS = "https://launchermeta.mojang.com/mc/game/version_manifest.json"
)

func Load() (*Remote, error) {
	res, err := http.Get(MINECRAFT_VERSIONS)

	if err != nil {
		return nil, fmt.Errorf("Failed to fetch minecraft versions. %s", err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("Failed to read minecraft versions. %s", err)
	}

	remote := &Remote{}

	if err := json.Unmarshal(body, &remote); err != nil {
		return nil, fmt.Errorf("Failed to parse minecraft versions. %s", err)
	}

	return remote, nil
}
