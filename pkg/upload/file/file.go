package file

import (
	"net/http"
	"net/url"
	"path"
	"os"
	"fmt"
	"strconv"

	"github.com/kleister/kleister-api/pkg/upload"
)

type file struct {
	dsn *url.URL
}

// Info prepares some informational message about the handler.
func (u *file) Info() string {
	return fmt.Sprintf("prepared file storage at %s", u.path())
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

// Handler implements an HTTP handler for asset uploads.
func (u *file) Handler(root string) http.Handler {
	return http.StripPrefix(
		root + "/",
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

	return 0755
}

// path cleans the dsn and returns a valid path.
func (u *file) path() string {
	return path.Join(
		u.dsn.Host,
		u.dsn.EscapedPath(),
	)
}

// New initializes a new file handler.
func New(dsn *url.URL) (upload.Upload, error) {
	f := &file{
		dsn: dsn,
	}

	return f.Prepare()
}

// Must simply calls New and panics on an error.
func Must(dsn *url.URL) upload.Upload {
	db, err := New(dsn)

	if err != nil {
		panic(err)
	}

	return db
}
