package upload

import (
	"bytes"
	"fmt"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	s3client "github.com/aws/aws-sdk-go/service/s3"
	"github.com/kleister/kleister-api/pkg/config"
)

// S3Upload implements the Upload interface.
type S3Upload struct {
	endpoint string
	path     string
	access   string
	secret   string
	bucket   string
	region   string

	client *s3client.S3
}

// Info prepares some informational message about the handler.
func (u *S3Upload) Info() map[string]interface{} {
	result := make(map[string]interface{})
	result["driver"] = "s3"
	result["endpoint"] = u.endpoint
	result["path"] = u.path
	result["bucket"] = u.bucket
	result["region"] = u.region

	return result
}

// Prepare simply prepares the upload handler.
func (u *S3Upload) Prepare() (Upload, error) {

	cfg := aws.NewConfig().
		WithRegion(u.region).
		WithS3ForcePathStyle(true)

	if u.endpoint != "" {
		cfg = cfg.WithEndpoint(u.endpoint)
	}

	if u.access != "" && u.secret != "" {
		access, err := config.Value(u.access)

		if err != nil {
			return nil, fmt.Errorf("failed to parse access key: %w", err)
		}

		secret, err := config.Value(u.secret)

		if err != nil {
			return nil, fmt.Errorf("failed to parse secret key: %w", err)
		}

		cfg = cfg.WithCredentials(credentials.NewStaticCredentials(
			access,
			secret,
			"",
		))
	}

	sess, err := session.NewSession()

	if err != nil {
		return u, err
	}

	u.client = s3client.New(
		sess,
		cfg,
	)

	return u, nil
}

// Close simply closes the upload handler.
func (u *S3Upload) Close() error {
	return nil
}

// Upload stores an attachment within the defined S3 bucket.
func (u *S3Upload) Upload(key, ctype string, content []byte) error {
	params := &s3client.PutObjectInput{
		ACL:         aws.String("public-read"),
		Bucket:      aws.String(u.bucket),
		Key:         aws.String(path.Join(u.path, key)),
		ContentType: aws.String(ctype),
		Body:        bytes.NewReader(content),
	}

	if _, err := u.client.PutObject(
		params,
	); err != nil {
		return err
	}

	return nil
}

// Delete removes an attachment from the defined S3 bucket.
func (u *S3Upload) Delete(key string) error {
	params := &s3client.DeleteObjectInput{
		Bucket: aws.String(u.bucket),
		Key:    aws.String(path.Join(u.path, key)),
	}

	if _, err := u.client.DeleteObject(
		params,
	); err != nil {
		return err
	}

	return nil
}

// Handler implements an HTTP handler for asset uploads.
func (u *S3Upload) Handler(root string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, _ := u.client.GetObjectRequest(&s3client.GetObjectInput{
			Bucket: aws.String(u.bucket),
			Key:    aws.String(path.Join(u.path, strings.TrimPrefix(r.URL.Path, root))),
		})

		url, err := req.Presign(5 * time.Minute)

		if err != nil {
			http.Error(
				w,
				http.StatusText(http.StatusNotFound),
				http.StatusNotFound,
			)
		}

		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	})
}

// NewS3Upload initializes a new S3 handler.
func NewS3Upload(cfg config.Upload) (Upload, error) {
	path := "/"

	if cfg.Path != "" {
		path = cfg.Path
	}

	f := &S3Upload{
		endpoint: cfg.Endpoint,
		path:     path,
		access:   cfg.Access,
		secret:   cfg.Secret,
		bucket:   cfg.Bucket,
		region:   cfg.Region,
	}

	return f.Prepare()
}

// MustS3Upload simply calls NewS3Upload and panics on an error.
func MustS3Upload(cfg config.Upload) Upload {
	db, err := NewS3Upload(cfg)

	if err != nil {
		panic(err)
	}

	return db
}
