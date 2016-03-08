package config

// Config provides all application config values.
type Config struct {
	Version string
	Debug   bool
	Develop bool
	Server  struct {
		Addr string
		Cert string
		Key  string
		Root string
	}
	Database struct {
		Driver   string
		Username string
		Password string
		Name     string
		Host     string
	}
}
