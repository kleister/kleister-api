package forge

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	FORGE_VERSIONS = "http://files.minecraftforge.net/maven/net/minecraftforge/forge/json"
)

func Load() (*Remote, error) {
	res, err := http.Get(FORGE_VERSIONS)

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

	fmt.Println(remote)

	return remote, nil
}
