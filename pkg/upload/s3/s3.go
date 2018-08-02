package s3

import (
	"net/http"
	"net/url"

	"github.com/kleister/kleister-api/pkg/upload"
)

type s3 struct {
	dsn *url.URL
}

// Close simply closes the upload handler.
func (u *s3) Close() error {
	return nil
}

// Handler implements an HTTP handler for asset uploads.
func (u *s3) Handler() http.Handler {
	// root.Handle("/storage/*", files(logger))

	// func files(logger log.Logger) http.Handler {
	// 	return http.StripPrefix(
	// 		fmt.Sprintf(
	// 			"%sstorage",
	// 			config.Server.Root,
	// 		),
	// 		http.FileServer(
	// 			http.Dir(config.Server.Storage),
	// 		),
	// 	)
	// }

	return nil
}

// New initializes a new S3 handler.
func New(dsn *url.URL) (upload.Upload, error) {
	return &s3{
		dsn: dsn,
	}, nil
}

// Must simply calls New and panics on an error.
func Must(dsn *url.URL) upload.Upload {
	db, err := New(dsn)

	if err != nil {
		panic(err)
	}

	return db
}
