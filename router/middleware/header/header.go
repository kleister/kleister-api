package header

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kleister/kleister-api/pkg/version"
)

// SetCache writes required cache headers to all requests.
func SetCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
		c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
		c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))

		c.Next()
	}
}

// SetOptions writes required option headers to all requests.
func SetOptions() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "OPTIONS" {
			c.Next()
		} else {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
			c.Header("Allow", "HEAD, GET, POST, PUT, PATCH, DELETE, OPTIONS")
			c.Header("Content-Type", "application/json")

			c.AbortWithStatus(200)
		}
	}
}

// SetSecure writes required access headers to all requests.
func SetSecure() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-XSS-Protection", "1; mode=block")

		if c.Request.TLS != nil {
			c.Header("Strict-Transport-Security", "max-age=31536000")
		}

		c.Next()
	}
}

// SetVersion writes the current API version to the headers.
func SetVersion() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-KLEISTER-VERSION", version.Version.String())
		c.Next()
	}
}
