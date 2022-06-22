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
	Info() map[string]interface{}
	Prepare() (Upload, error)
	Close() error
	Upload(string, string, []byte) error
	Delete(string) error
	Handler(string) http.Handler
}
