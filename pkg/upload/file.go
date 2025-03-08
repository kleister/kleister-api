package upload

import (
	"bytes"
	"context"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/kleister/kleister-api/pkg/config"
)

// FileUpload implements the Upload interface.
type FileUpload struct {
	path  string
	perms fs.FileMode
	root  *os.Root
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

	{
		root, err := os.OpenRoot(u.path)

		if err != nil {
			return nil, err
		}

		u.root = root
	}

	return u, nil
}

// Close simply closes the upload handler.
func (u *FileUpload) Close() error {
	return u.root.Close()
}

// Upload stores an attachment within the defined S3 bucket.
func (u *FileUpload) Upload(_ context.Context, path string, content *bytes.Buffer) error {
	parent := filepath.Dir(
		path,
	)

	if _, err := u.root.Stat(parent); os.IsNotExist(err) {
		if err := os.MkdirAll(
			filepath.Join(
				u.root.Name(),
				parent,
			),
			u.perms,
		); err != nil {
			return err
		}
	}

	file, err := u.root.OpenFile(
		path,
		os.O_CREATE|os.O_TRUNC|os.O_RDWR,
		u.mode(),
	)

	if err != nil {
		return err
	}

	defer file.Close()

	if _, err = file.Write(
		content.Bytes(),
	); err != nil {
		return err
	}

	return nil
}

// Delete removes an attachment from the defined S3 bucket.
func (u *FileUpload) Delete(_ context.Context, path string, recursive bool) error {
	if recursive {
		fullPath := filepath.Join(
			u.root.Name(),
			path,
		)

		relPath, err := filepath.Rel(
			u.root.Name(),
			fullPath,
		)

		if err != nil || relPath == ".." || relPath[:3] == "../" {
			return fmt.Errorf("denied to delete unsafe path")
		}

		return os.RemoveAll(fullPath)
	}

	return u.root.Remove(path)
}

// Handler implements an HTTP handler for asset uploads.
func (u *FileUpload) Handler(root string) http.Handler {
	if !strings.HasSuffix(root, "/") {
		root = root + "/"
	}

	return http.StripPrefix(
		root,
		http.FileServer(
			http.FS(
				u.root.FS(),
			),
		),
	)
}

func (u *FileUpload) mode() os.FileMode {
	mode := u.perms &^ 0111
	mode &^= os.ModeDir
	return mode
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
