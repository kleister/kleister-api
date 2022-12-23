package upload

import (
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"

	"github.com/kleister/kleister-api/pkg/config"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// FileUpload implements the Upload interface.
type FileUpload struct {
	dsn *url.URL
}

// Info prepares some informational message about the handler.
func (u *FileUpload) Info() map[string]interface{} {
	result := make(map[string]interface{})
	result["driver"] = "file"
	result["path"] = u.path()

	return result
}

// Prepare simply prepares the upload handler.
func (u *FileUpload) Prepare() (Upload, error) {
	if _, err := os.Stat(u.path()); os.IsNotExist(err) {
		if err := os.MkdirAll(u.path(), u.perms()); err != nil {
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
	return http.StripPrefix(
		root+"/",
		http.FileServer(
			http.Dir(u.path()),
		),
	)
}

// perms retrieves the dir perms from dsn or fallback.
func (u *FileUpload) perms() os.FileMode {
	if val := u.dsn.Query().Get("perms"); val != "" {
		u, err := strconv.ParseUint(val, 8, 32)

		if err != nil {
			return 0755
		}

		return os.FileMode(u)
	}

	return os.FileMode(0755)
}

// path cleans the dsn and returns a valid path.
func (u *FileUpload) path() string {
	return path.Join(
		u.dsn.Host,
		u.dsn.EscapedPath(),
	)
}

// NewFileUpload initializes a new file handler.
func NewFileUpload(cfg config.Upload) (Upload, error) {
	parsed, err := url.Parse(cfg.DSN)

	if err != nil {
		return nil, errors.Wrap(err, "failed to parse dsn")
	}

	f := &FileUpload{
		dsn: parsed,
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
