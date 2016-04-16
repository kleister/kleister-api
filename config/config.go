package config

// Config provides all application config values.
type Config struct {
	Debug  bool
	Server struct {
		Addr    string
		Cert    string
		Key     string
		Root    string
		Storage string
	}
	Database struct {
		Driver   string
		Username string
		Password string
		Name     string
		Host     string
	}
}

var (
	Values = &Config{}
)
