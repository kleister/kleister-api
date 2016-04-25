package config

type database struct {
	Driver   string
	Username string
	Password string
	Name     string
	Host     string
}

type server struct {
	Addr    string
	Cert    string
	Key     string
	Root    string
	Storage string
}

var (
	// Update represents the flag to enable or disable auto updates.
	Update bool

	// Debug represents the flag to enable or disable debug logging.
	Debug bool

	// Database represents the current database connection details.
	Database = &database{}

	// Server represents the informations about the server bindings.
	Server = &server{}
)
