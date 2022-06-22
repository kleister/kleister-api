package s3

import (
	"bytes"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	s3client "github.com/aws/aws-sdk-go/service/s3"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/upload"
	"github.com/pkg/errors"
)

type s3 struct {
	dsn    *url.URL
	client *s3client.S3
}

// Info prepares some informational message about the handler.
func (u *s3) Info() map[string]interface{} {
	result := make(map[string]interface{})
	result["driver"] = "s3"
	result["bucket"] = u.bucket()
	result["region"] = u.region()
	result["pathstyle"] = u.pathstyle()
	result["endpoint"] = u.endpoint()

	return result
}

// Prepare simply prepares the upload handler.
func (u *s3) Prepare() (upload.Upload, error) {
	cfg := aws.NewConfig().
		WithRegion(u.region()).
		WithS3ForcePathStyle(u.pathstyle())

	if u.endpoint() != "" {
		cfg = cfg.WithEndpoint(u.endpoint())
	}

	if u.accesskey() != "" && u.secretkey() != "" {
		cfg = cfg.WithCredentials(credentials.NewStaticCredentials(
			u.accesskey(),
			u.secretkey(),
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
func (u *s3) Close() error {
	return nil
}

// Upload stores an attachment within the defined S3 bucket.
func (u *s3) Upload(path, ctype string, content []byte) error {
	params := &s3client.PutObjectInput{
		ACL:         aws.String("public-read"),
		Bucket:      aws.String(u.bucket()),
		Key:         aws.String(path),
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
func (u *s3) Delete(path string) error {
	params := &s3client.DeleteObjectInput{
		Bucket: aws.String(u.bucket()),
		Key:    aws.String(path),
	}

	if _, err := u.client.DeleteObject(
		params,
	); err != nil {
		return err
	}

	return nil
}

// Handler implements an HTTP handler for asset uploads.
func (u *s3) Handler(root string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, _ := u.client.GetObjectRequest(&s3client.GetObjectInput{
			Bucket: aws.String(u.bucket()),
			Key:    aws.String(strings.TrimPrefix(r.URL.Path, root)),
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

func (u *s3) accesskey() string {
	if val := u.dsn.Query().Get("access_key"); val != "" {
		return val
	}

	return ""
}

func (u *s3) secretkey() string {
	if val := u.dsn.Query().Get("secret_key"); val != "" {
		return val
	}

	return ""
}

func (u *s3) bucket() string {
	if val := u.dsn.Query().Get("bucket"); val != "" {
		return val
	}

	return ""
}

func (u *s3) region() string {
	if val := u.dsn.Query().Get("region"); val != "" {
		return val
	}

	return ""
}

func (u *s3) pathstyle() bool {
	if val := u.dsn.Query().Get("path_style"); val != "" {
		u, err := strconv.ParseBool(val)

		if err != nil {
			return false
		}

		return u
	}

	return false
}

func (u *s3) endpoint() string {
	return u.dsn.Host
}

// New initializes a new S3 handler.
func New(cfg config.Upload) (upload.Upload, error) {
	parsed, err := url.Parse(cfg.DSN)

	if err != nil {
		return nil, errors.Wrap(err, "failed to parse dsn")
	}

	f := &s3{
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
