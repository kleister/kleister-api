package config

import (
	"time"
)

// Logs defines the level and color for log configuration.
type Logs struct {
	Level  string `mapstructure:"level"`
	Pretty bool   `mapstructure:"pretty"`
	Color  bool   `mapstructure:"color"`
}

// Database defines the database configuration.
type Database struct {
	DSN string `mapstructure:"dsn"`
}

// Upload defines the asset upload configuration.
type Upload struct {
	DSN string `mapstructure:"dsn"`
}

// Session defines the session handle configuration.
type Session struct {
	Expire time.Duration `mapstructure:"expire"`
	Secret string        `mapstructure:"secret"`
}

// Server defines the webserver configuration.
type Server struct {
	Host string `mapstructure:"host"`
	Root string `mapstructure:"root"`
	Addr string `mapstructure:"addr"`
	Cert string `mapstructure:"cert"`
	Key  string `mapstructure:"key"`
}

// Metrics defines the metrics server configuration.
type Metrics struct {
	Addr  string `mapstructure:"addr"`
	Token string `mapstructure:"token"`
	Pprof bool   `mapstructure:"pprof"`
}

// Admin defines the initial admin user configuration.
type Admin struct {
	Create   bool   `mapstructure:"create"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Email    string `mapstructure:"emails"`
}

// Config is a combination of all available configurations.
type Config struct {
	File     string   `mapstructure:"-"`
	Logs     Logs     `mapstructure:"logs"`
	Database Database `mapstructure:"database"`
	Upload   Upload   `mapstructure:"upload"`
	Session  Session  `mapstructure:"session"`
	Server   Server   `mapstructure:"server"`
	Metrics  Metrics  `mapstructure:"metrics"`
	Admin    Admin    `mapstructure:"admins"`
}

// Load initializes a default configuration struct.
func Load() *Config {
	return &Config{}
}
