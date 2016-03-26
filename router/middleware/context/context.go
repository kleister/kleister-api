package context

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/config"
	"github.com/solderapp/solder-api/model"
)

// Config gets the config from the context.
func Config(c *gin.Context) config.Config {
	return c.MustGet("config").(config.Config)
}

// SetConfig injects the config into the context.
func SetConfig(config config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("config", config)
		c.Next()
	}
}

// Store gets the storage from the context.
func Store(c *gin.Context) model.Store {
	return c.MustGet("store").(model.Store)
}

// SetStore injects the storage into the context.
func SetStore(store model.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("store", store)
		c.Next()
	}
}

// Root gets the root URL from the context.
func Root(c *gin.Context) string {
	return c.MustGet("root").(string)
}

// SetRoot injects the root URL into the context.
func SetRoot() gin.HandlerFunc {
	return func(c *gin.Context) {
		config := Config(c)

		root := fmt.Sprintf(
			"%s://%s%s",
			resolveScheme(c.Request),
			resolveHost(c.Request),
			config.Server.Root,
		)

		if strings.HasSuffix(root, "/") {
			root = strings.TrimSuffix(root, "/")
		}

		c.Set("root", root)
		c.Next()
	}
}

// resolveScheme is a helper function that evaluates the http.Request and
// returns the scheme, HTTP or HTTPS. It is able to detect, using the
// X-Forwarded-Proto, if the original request was HTTPS and routed through
// a reverse proxy with SSL termination.
func resolveScheme(r *http.Request) string {
	switch {
	case r.URL.Scheme == "https":
		return "https"
	case r.TLS != nil:
		return "https"
	case strings.HasPrefix(r.Proto, "HTTPS"):
		return "https"
	case r.Header.Get("X-Forwarded-Proto") == "https":
		return "https"
	case len(r.Header.Get("Forwarded")) != 0 && len(parseHeader(r, "Forwarded", "proto")) != 0 && parseHeader(r, "Forwarded", "proto")[0] == "https":
		return "https"
	default:
		return "http"
	}
}

// resolveHost is a helper function that evaluates the http.Request and
// returns the hostname. It is able to detect, using the X-Forarded-For
// header, the original hostname when routed through a reverse proxy.
func resolveHost(r *http.Request) string {
	switch {
	case len(r.Host) != 0:
		return r.Host
	case len(r.URL.Host) != 0:
		return r.URL.Host
	case len(r.Header.Get("X-Forwarded-For")) != 0:
		return r.Header.Get("X-Forwarded-For")
	case len(r.Header.Get("Forwarded")) != 0 && len(parseHeader(r, "Forwarded", "for")) != 0:
		return parseHeader(r, "Forwarded", "for")[0]
	case len(r.Header.Get("X-Host")) != 0:
		return r.Header.Get("X-Host")
	case len(r.Header.Get("Forwarded")) != 0 && len(parseHeader(r, "Forwarded", "host")) != 0:
		return parseHeader(r, "Forwarded", "host")[0]
	case len(r.Header.Get("XFF")) != 0:
		return r.Header.Get("XFF")
	case len(r.Header.Get("X-Real-IP")) != 0:
		return r.Header.Get("X-Real-IP")
	default:
		return "localhost:8080"
	}
}

// parseHeader parses non unique headers value from a http.Request and
// return a slice of the values queried from the header
func parseHeader(r *http.Request, header string, token string) (val []string) {
	for _, v := range r.Header[header] {
		options := strings.Split(v, ";")

		for _, o := range options {
			keyvalue := strings.Split(o, "=")
			var key, value string

			if len(keyvalue) > 1 {
				key, value = strings.TrimSpace(keyvalue[0]), strings.TrimSpace(keyvalue[1])
			}

			key = strings.ToLower(key)

			if key == token {
				val = append(val, value)
			}
		}
	}

	return
}
