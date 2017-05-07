package config

import (
	"time"
)

type database struct {
	Driver   string
	Username string
	Password string
	Name     string
	Host     string
}

type server struct {
	Host          string
	Addr          string
	Cert          string
	Key           string
	Root          string
	Storage       string
	Assets        string
	LetsEncrypt   bool
	StrictCurves  bool
	StrictCiphers bool
	Pprof         bool
}

type admin struct {
	Users  []string
	Create bool
}

type s3 struct {
	Enabled   bool
	Endpoint  string
	Bucket    string
	PathStyle bool
	Region    string
	Access    string
	Secret    string
}

type session struct {
	Expire time.Duration
}

var (
	// Debug represents the flag to enable or disable debug logging.
	Debug bool

	// Database represents the current database connection details.
	Database = &database{}

	// Server represents the informations about the server bindings.
	Server = &server{}

	// Admin represents the informations about the admin config.
	Admin = &admin{}

	// S3 represents the informations about s3 storage connections.
	S3 = &s3{}

	// Session represents the informations about the session handling.
	Session = &session{}
)
