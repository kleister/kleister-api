package config

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"time"
)

// Server defines the webserver configuration.
type Server struct {
	Addr string `mapstructure:"addr"`
	Host string `mapstructure:"host"`
	Root string `mapstructure:"root"`
	Cert string `mapstructure:"cert"`
	Key  string `mapstructure:"key"`
	Docs bool   `mapstructure:"docs"`
}

// Metrics defines the metrics server configuration.
type Metrics struct {
	Addr  string `mapstructure:"addr"`
	Token string `mapstructure:"token"`
	Pprof bool   `mapstructure:"pprof"`
}

// Logs defines the level and color for log configuration.
type Logs struct {
	Level  string `mapstructure:"level"`
	Pretty bool   `mapstructure:"pretty"`
	Color  bool   `mapstructure:"color"`
}

// Auth defines the authentication configuration.
type Auth struct {
	Config string `mapstructure:"config"`
}

// Database defines the database configuration.
type Database struct {
	Driver   string            `mapstructure:"driver"`
	Address  string            `mapstructure:"address"`
	Port     string            `mapstructure:"port"`
	Username string            `mapstructure:"username"`
	Password string            `mapstructure:"password"`
	Name     string            `mapstructure:"name"`
	Options  map[string]string `mapstructure:"options"`
}

// Upload defines the asset upload configuration.
type Upload struct {
	Driver   string `mapstructure:"driver"`
	Endpoint string `mapstructure:"endpoint"`
	Path     string `mapstructure:"path"`
	Access   string `mapstructure:"access"`
	Secret   string `mapstructure:"secret"`
	Bucket   string `mapstructure:"bucket"`
	Region   string `mapstructure:"region"`
	Perms    string `mapstructure:"perms"`
}

// Session defines the session handle configuration.
type Session struct {
	Secret string        `mapstructure:"secret"`
	Expire time.Duration `mapstructure:"expire"`
	Secure bool          `mapstructure:"secure"`
}

// Admin defines the initial admin user configuration.
type Admin struct {
	Create   bool   `mapstructure:"create"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Email    string `mapstructure:"email"`
}

// Config is a combination of all available configurations.
type Config struct {
	Server   Server   `mapstructure:"server"`
	Metrics  Metrics  `mapstructure:"metrics"`
	Logs     Logs     `mapstructure:"log"`
	Auth     Auth     `mapstructure:"auth"`
	Database Database `mapstructure:"database"`
	Upload   Upload   `mapstructure:"upload"`
	Session  Session  `mapstructure:"session"`
	Admin    Admin    `mapstructure:"admin"`
}

// Load initializes a default configuration struct.
func Load() *Config {
	return &Config{}
}

// Value returns the config value based on a DSN.
func Value(val string) (string, error) {
	if strings.HasPrefix(val, "file://") {
		content, err := os.ReadFile(
			strings.TrimPrefix(val, "file://"),
		)

		if err != nil {
			return "", fmt.Errorf("failed to parse secret file: %w", err)
		}

		return string(content), nil
	}

	if strings.HasPrefix(val, "base64://") {
		content, err := base64.StdEncoding.DecodeString(
			strings.TrimPrefix(val, "base64://"),
		)

		if err != nil {
			return "", fmt.Errorf("failed to parse base64 value: %w", err)
		}

		return string(content), nil
	}

	return val, nil
}
