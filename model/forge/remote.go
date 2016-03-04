package forge

type Remote struct {
	Webpath string             `json:"webpath"`
	Numbers map[string]*Number `json:"number"`
}
