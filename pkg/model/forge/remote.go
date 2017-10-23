package forge

// Remote represents the structure of the remote info.
type Remote struct {
	Numbers map[string]*Number `json:"number"`
}
