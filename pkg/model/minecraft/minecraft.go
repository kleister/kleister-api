package minecraft

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// MinecraftVersions represents the URL to fetch all available Minecraft versions
	MinecraftVersions = "https://launchermeta.mojang.com/mc/game/version_manifest.json"
)

// Load initializes and fetches the Minecraft versions from the remote service.
func Load() (*Remote, error) {
	res, err := http.Get(MinecraftVersions)

	if err != nil {
		return nil, fmt.Errorf("Failed to fetch Minecraft versions. %s", err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("Failed to read Minecraft versions. %s", err)
	}

	remote := &Remote{}

	if err := json.Unmarshal(body, &remote); err != nil {
		return nil, fmt.Errorf("Failed to parse Minecraft versions. %s", err)
	}

	return remote, nil
}
