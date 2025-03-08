package config

// AuthEndpoints defines the endpoints for external authentication.
type AuthEndpoints struct {
	Issuer  string `mapstructure:"issuer"`
	Auth    string `mapstructure:"auth"`
	Token   string `mapstructure:"token"`
	Profile string `mapstructure:"profile"`
	Email   string `mapstructure:"email"`
}

// AuthMappings defines the mappings for external authentication.
type AuthMappings struct {
	Login string `mapstructure:"login"`
	Name  string `mapstructure:"name"`
	Email string `mapstructure:"email"`
	Role  string `mapstructure:"role"`
}

// AuthAdmins defines the mappings for administrative users.
type AuthAdmins struct {
	Users  []string `mapstructure:"users"`
	Emails []string `mapstructure:"emails"`
	Roles  []string `mapstructure:"roles"`
}

// AuthProvider defines a single provider auth source.
type AuthProvider struct {
	Driver       string        `mapstructure:"driver"`
	Name         string        `mapstructure:"name"`
	Display      string        `mapstructure:"display"`
	Icon         string        `mapstructure:"icon"`
	Callback     string        `mapstructure:"callback"`
	ClientID     string        `mapstructure:"client_id"`
	ClientSecret string        `mapstructure:"client_secret"`
	Verifier     string        `mapstructure:"verifier"`
	Tenant       string        `mapstructure:"tenant"`
	Scopes       []string      `mapstructure:"scopes"`
	Endpoints    AuthEndpoints `mapstructure:"endpoints"`
	Mappings     AuthMappings  `mapstructure:"mappings"`
	Admins       AuthAdmins    `mapstructure:"admins"`
}

// AuthConfig defines the configuration for auth sources.
type AuthConfig struct {
	Providers []AuthProvider `mapstructure:"providers"`
}
