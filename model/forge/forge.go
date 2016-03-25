package forge

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// ForgeVersions represents the URL to fetch all available Forge versions
	ForgeVersions = "http://files.minecraftforge.net/maven/net/minecraftforge/forge/json"
)

// Load initializes and fetches the Forge versions from the remote service.
func Load() (*Remote, error) {
	res, err := http.Get(ForgeVersions)

	if err != nil {
		return nil, fmt.Errorf("Failed to fetch Forge versions. %s", err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("Failed to read Forge versions. %s", err)
	}

	remote := &Remote{}

	if err := json.Unmarshal(body, &remote); err != nil {
		return nil, fmt.Errorf("Failed to parse Forge versions. %s", err)
	}

	return remote, nil
}
