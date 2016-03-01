package config

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
		Driver string
		Config string
	}
}
