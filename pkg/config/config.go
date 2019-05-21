package config

// Database defines the database configuration.
type Database struct {
	DSN string
}

// Upload defines the asset upload configuration.
type Upload struct {
	DSN string
}

// Server defines the webserver configuration.
type Server struct {
	Host  string
	Root  string
	Addr  string
	Pprof bool
	Docs  bool
}

// Metrics defines the metrics server configuration.
type Metrics struct {
	Addr  string
	Token string
}

// Admin defines the initial admin user configuration.
type Admin struct {
	Create   bool
	Username string
	Password string
	Email    string
}

// Logs defines the level and color for log configuration.
type Logs struct {
	Level  string
	Pretty bool
	Color  bool
}

// Tracing defines the tracing client configuration.
type Tracing struct {
	Enabled  bool
	Endpoint string
}

// Config is a combination of all available configurations.
type Config struct {
	Database Database
	Upload   Upload
	Server   Server
	Metrics  Metrics
	Admin    Admin
	Logs     Logs
	Tracing  Tracing
}

// Load initializes a default configuration struct.
func Load() *Config {
	return &Config{}
}
