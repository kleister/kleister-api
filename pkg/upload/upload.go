package upload

import (
	"net/http"

	"github.com/pkg/errors"
)

var (
	// ErrUnknownDriver defines a named error for unknown upload drivers.
	ErrUnknownDriver = errors.New("unknown upload driver")
)

// Upload provides the interface for the upload implementations.
type Upload interface {
	Close() error
	Handler() http.Handler
}

// Load initializes the upload implementation based on supplied config.
func Load(dsn string) (Upload, error) {
	parsed, err := url.Parse(dsn)

	if err != nil {
		return nil, errors.Wrap(err, "failed to parse dsn")
	}

	switch parsed.Scheme {
	case "file":
		return file.New(parsed)
	case "s3":
		return s3.New(parsed)
	case "minio":
		return s3.New(parsed)
	}

	return nil, ErrUnknownDriver
}
