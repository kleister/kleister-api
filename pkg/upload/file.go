package upload

import (
	"io/fs"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/kleister/kleister-api/pkg/config"
	"github.com/rs/zerolog/log"
)

// FileUpload implements the Upload interface.
type FileUpload struct {
	path  string
	perms fs.FileMode
}

// Info prepares some informational message about the handler.
func (u *FileUpload) Info() map[string]interface{} {
	result := make(map[string]interface{})
	result["driver"] = "file"
	result["path"] = u.path

	return result
}

// Prepare simply prepares the upload handler.
func (u *FileUpload) Prepare() (Upload, error) {
	if _, err := os.Stat(u.path); os.IsNotExist(err) {
		if err := os.MkdirAll(u.path, u.perms); err != nil {
			return nil, err
		}
	}

	return u, nil
}

// Close simply closes the upload handler.
func (u *FileUpload) Close() error {
	return nil
}

// Upload stores an attachment within the defined S3 bucket.
func (u *FileUpload) Upload(path, ctype string, content []byte) error {
	log.Debug().
		Str("path", path).
		Str("ctype", ctype).
		Bytes("content", content).
		Msg("Upload")

	return nil
}

// Delete removes an attachment from the defined S3 bucket.
func (u *FileUpload) Delete(path string) error {
	log.Debug().
		Str("path", path).
		Msg("Delete")

	return nil
}

// Handler implements an HTTP handler for asset uploads.
func (u *FileUpload) Handler(root string) http.Handler {
	if !strings.HasSuffix(root, "/") {
		root = root + "/"
	}

	return http.StripPrefix(
		root,
		http.FileServer(
			http.Dir(u.path),
		),
	)
}

// NewFileUpload initializes a new file handler.
func NewFileUpload(cfg config.Upload) (Upload, error) {
	perms := os.FileMode(0755)

	if cfg.Perms != "" {
		res, err := strconv.ParseUint(
			cfg.Perms,
			8,
			32,
		)

		if err == nil {
			perms = os.FileMode(res)
		}
	}

	f := &FileUpload{
		path:  cfg.Path,
		perms: perms,
	}

	return f.Prepare()
}

// MustFileUpload simply calls NewFileUpload and panics on an error.
func MustFileUpload(cfg config.Upload) Upload {
	db, err := NewFileUpload(cfg)

	if err != nil {
		panic(err)
	}

	return db
}
