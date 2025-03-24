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
	Addr      string `mapstructure:"addr"`
	Host      string `mapstructure:"host"`
	Root      string `mapstructure:"root"`
	Cert      string `mapstructure:"cert"`
	Key       string `mapstructure:"key"`
	Templates string `mapstructure:"templates"`
	Frontend  string `mapstructure:"frontend"`
	Docs      bool   `mapstructure:"docs"`
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

// Cleanup defines the cleanup process configuration.
type Cleanup struct {
	Enabled  bool          `mapstructure:"enabled"`
	Interval time.Duration `mapstructure:"interval"`
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
	Driver    string `mapstructure:"driver"`
	Endpoint  string `mapstructure:"endpoint"`
	Pathstyle bool   `mapstructure:"pathstyle"`
	Path      string `mapstructure:"path"`
	Access    string `mapstructure:"access"`
	Secret    string `mapstructure:"secret"`
	Bucket    string `mapstructure:"bucket"`
	Region    string `mapstructure:"region"`
	Perms     string `mapstructure:"perms"`
	Proxy     bool   `mapstructure:"proxy"`
}

// Token defines the token handle configuration.
type Token struct {
	Secret string        `mapstructure:"secret"`
	Expire time.Duration `mapstructure:"expire"`
}

// Admin defines the initial admin user configuration.
type Admin struct {
	Create   bool   `mapstructure:"create"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Email    string `mapstructure:"email"`
}

// Scim defines the scim provisioning configuration.
type Scim struct {
	Enabled bool   `mapstructure:"enabled"`
	Token   string `mapstructure:"token"`
}

// Config is a combination of all available configurations.
type Config struct {
	Server   Server   `mapstructure:"server"`
	Metrics  Metrics  `mapstructure:"metrics"`
	Logs     Logs     `mapstructure:"log"`
	Cleanup  Cleanup  `mapstructure:"cleanup"`
	Auth     Auth     `mapstructure:"auth"`
	Database Database `mapstructure:"database"`
	Upload   Upload   `mapstructure:"upload"`
	Token    Token    `mapstructure:"token"`
	Scim     Scim     `mapstructure:"scim"`
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
