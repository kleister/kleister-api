package file

import (
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"

	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/upload"
	"github.com/pkg/errors"
)

type file struct {
	dsn *url.URL
}

// Info prepares some informational message about the handler.
func (u *file) Info() map[string]interface{} {
	result := make(map[string]interface{})
	result["driver"] = "file"
	result["path"] = u.path()

	return result
}

// Prepare simply prepares the upload handler.
func (u *file) Prepare() (upload.Upload, error) {
	if _, err := os.Stat(u.path()); os.IsNotExist(err) {
		if err := os.MkdirAll(u.path(), u.perms()); err != nil {
			return nil, err
		}
	}

	return u, nil
}

// Close simply closes the upload handler.
func (u *file) Close() error {
	return nil
}

// Upload stores an attachment within the defined S3 bucket.
func (u *file) Upload(path, ctype string, content []byte) error {
	return nil
}

// Delete removes an attachment from the defined S3 bucket.
func (u *file) Delete(path string) error {
	return nil
}

// Handler implements an HTTP handler for asset uploads.
func (u *file) Handler(root string) http.Handler {
	return http.StripPrefix(
		root+"/",
		http.FileServer(
			http.Dir(u.path()),
		),
	)
}

// perms retrieves the dir perms from dsn or fallback.
func (u *file) perms() os.FileMode {
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
func (u *file) path() string {
	return path.Join(
		u.dsn.Host,
		u.dsn.EscapedPath(),
	)
}

// New initializes a new file handler.
func New(cfg config.Upload) (upload.Upload, error) {
	parsed, err := url.Parse(cfg.DSN)

	if err != nil {
		return nil, errors.Wrap(err, "failed to parse dsn")
	}

	f := &file{
		dsn: parsed,
	}

	return f.Prepare()
}

// Must simply calls New and panics on an error.
func Must(cfg config.Upload) upload.Upload {
	db, err := New(cfg)

	if err != nil {
		panic(err)
	}

	return db
}
