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
	Debug    bool
	Database = &database{}
	Server   = &server{}
)
