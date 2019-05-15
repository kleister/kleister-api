package s3

import (
	"net/http"
	"net/url"

	"github.com/kleister/kleister-api/pkg/upload"
)

type s3 struct {
	dsn *url.URL
}

// Info prepares some informational message about the handler.
func (u *s3) Info() string {
	return ""
}

// Prepare simply prepares the upload handler.
func (u *s3) Prepare() (upload.Upload, error) {
	return u, nil
}

// Close simply closes the upload handler.
func (u *s3) Close() error {
	return nil
}

// Handler implements an HTTP handler for asset uploads.
func (u *s3) Handler(root string) http.Handler {
	return nil
}

// New initializes a new S3 handler.
func New(dsn *url.URL) (upload.Upload, error) {
	f := &s3{
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
