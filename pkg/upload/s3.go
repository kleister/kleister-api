package upload

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	transport "github.com/aws/smithy-go/endpoints"
	"github.com/gabriel-vasile/mimetype"
	"github.com/kleister/kleister-api/pkg/config"
)

// S3Upload implements the Upload interface.
type S3Upload struct {
	endpoint  string
	path      string
	access    string
	secret    string
	bucket    string
	region    string
	pathstyle bool
	proxy     bool

	client  *s3.Client
	presign *s3.PresignClient
}

// Info prepares some informational message about the handler.
func (u *S3Upload) Info() map[string]interface{} {
	result := make(map[string]interface{})
	result["driver"] = "s3"
	result["endpoint"] = u.endpoint
	result["path"] = u.path
	result["bucket"] = u.bucket
	result["region"] = u.region
	result["pathstyle"] = u.pathstyle
	result["proxy"] = u.proxy

	return result
}

// Prepare simply prepares the upload handler.
func (u *S3Upload) Prepare() (Upload, error) {
	globalOpts := []func(*awsconfig.LoadOptions) error{
		awsconfig.WithRegion(
			u.region,
		),
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

		globalOpts = append(
			globalOpts,
			awsconfig.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider(
					access,
					secret,
					"",
				),
			),
		)
	}

	cfg, err := awsconfig.LoadDefaultConfig(
		context.Background(),
		globalOpts...,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}

	serviceOpts := []func(*s3.Options){
		func(o *s3.Options) {
			o.UsePathStyle = u.pathstyle
		},
	}

	if u.endpoint != "" {
		endpoint, err := url.Parse(u.endpoint)

		if err != nil {
			return nil, err
		}

		serviceOpts = append(
			serviceOpts,
			s3.WithEndpointResolverV2(
				&CustomEndpointResolver{
					Endpoint:  endpoint,
					Bucket:    u.bucket,
					PathStyle: u.pathstyle,
				},
			),
		)

		if endpoint.Scheme == "http" {
			serviceOpts = append(
				serviceOpts,
				func(o *s3.Options) {
					o.HTTPClient = &http.Client{
						Transport: &http.Transport{
							TLSClientConfig: nil,
						},
					}
				},
			)
		}
	}

	u.client = s3.NewFromConfig(
		cfg,
		serviceOpts...,
	)

	u.presign = s3.NewPresignClient(
		u.client,
	)

	return u, nil
}

// Close simply closes the upload handler.
func (u *S3Upload) Close() error {
	return nil
}

// Upload stores an attachment within the defined S3 bucket.
func (u *S3Upload) Upload(ctx context.Context, key string, content *bytes.Buffer) error {
	reader := bytes.NewReader(
		content.Bytes(),
	)

	mtype, err := mimetype.DetectReader(
		reader,
	)

	if err != nil {
		return err
	}

	params := &s3.PutObjectInput{
		ACL:         types.ObjectCannedACLPublicRead,
		Bucket:      aws.String(u.bucket),
		Key:         aws.String(path.Join(u.path, key)),
		ContentType: aws.String(mtype.String()),
		Body:        reader,
	}

	if _, err := u.client.PutObject(
		ctx,
		params,
	); err != nil {
		return err
	}

	return nil
}

// Delete removes an attachment from the defined S3 bucket.
func (u *S3Upload) Delete(ctx context.Context, key string, recursive bool) error {
	if recursive {
		var (
			continuation *string
		)

		for {
			objects, err := u.client.ListObjectsV2(
				ctx, &s3.ListObjectsV2Input{
					Bucket:            aws.String(u.bucket),
					Prefix:            aws.String(path.Join(u.path, key)),
					ContinuationToken: continuation,
				},
			)

			if err != nil {
				return fmt.Errorf("failed to list objects: %w", err)
			}

			if len(objects.Contents) == 0 {
				break
			}

			deletable := make([]types.ObjectIdentifier, 0)

			for _, object := range objects.Contents {
				deletable = append(
					deletable,
					types.ObjectIdentifier{
						Key: object.Key,
					},
				)
			}

			if _, err = u.client.DeleteObjects(
				ctx,
				&s3.DeleteObjectsInput{
					Bucket: aws.String(u.bucket),
					Delete: &types.Delete{
						Objects: deletable,
						Quiet:   aws.Bool(true),
					},
				},
			); err != nil {
				return fmt.Errorf("failed to delete objects: %w", err)
			}

			if !aws.ToBool(objects.IsTruncated) {
				break
			}

			continuation = objects.NextContinuationToken
		}
	} else {
		params := &s3.DeleteObjectInput{
			Bucket: aws.String(u.bucket),
			Key:    aws.String(path.Join(u.path, key)),
		}

		if _, err := u.client.DeleteObject(
			ctx,
			params,
		); err != nil {
			return err
		}
	}

	return nil
}

// Handler implements an HTTP handler for asset uploads.
func (u *S3Upload) Handler(root string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if u.proxy {
			u.proxyHandler(root, w, r)
		} else {
			u.presignHandler(root, w, r)
		}
	})
}

func (u *S3Upload) proxyHandler(root string, w http.ResponseWriter, r *http.Request) {
	obj := strings.TrimPrefix(
		path.Join(
			u.path,
			strings.TrimPrefix(
				r.URL.Path,
				root,
			),
		),
		"/",
	)

	req, err := u.client.GetObject(
		r.Context(),
		&s3.GetObjectInput{
			Bucket: aws.String(u.bucket),
			Key:    aws.String(obj),
		},
	)

	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusNotFound),
			http.StatusNotFound,
		)

		return
	}

	defer func() { _ = req.Body.Close() }()

	modified := time.Now()

	if req.LastModified != nil {
		modified = *req.LastModified
	}

	buffer, err := io.ReadAll(req.Body)

	if err != nil {
		http.Error(
			w,
			"Failed to read buffer",
			http.StatusInternalServerError,
		)

		return
	}

	http.ServeContent(
		w,
		r,
		obj,
		modified,
		bytes.NewReader(buffer),
	)
}

func (u *S3Upload) presignHandler(root string, w http.ResponseWriter, r *http.Request) {
	obj := strings.TrimPrefix(
		path.Join(
			u.path,
			strings.TrimPrefix(
				r.URL.Path,
				root,
			),
		),
		"/",
	)

	req, err := u.presign.PresignGetObject(
		r.Context(),
		&s3.GetObjectInput{
			Bucket: aws.String(u.bucket),
			Key:    aws.String(obj),
		},
		func(opts *s3.PresignOptions) {
			opts.Expires = time.Duration(5 * time.Minute)
		},
	)

	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusNotFound),
			http.StatusNotFound,
		)

		return
	}

	http.Redirect(w, r, req.URL, http.StatusTemporaryRedirect)
}

// NewS3Upload initializes a new S3 handler.
func NewS3Upload(cfg config.Upload) (Upload, error) {
	f := &S3Upload{
		endpoint:  cfg.Endpoint,
		pathstyle: cfg.Pathstyle,
		path:      cfg.Path,
		access:    cfg.Access,
		secret:    cfg.Secret,
		bucket:    cfg.Bucket,
		region:    cfg.Region,
		proxy:     cfg.Proxy,
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

// CustomEndpointResolver is used for S3 compatible storage endpoints.
type CustomEndpointResolver struct {
	Endpoint  *url.URL
	Bucket    string
	PathStyle bool
}

// ResolveEndpoint resolves endpoints for a specific service and region
func (r *CustomEndpointResolver) ResolveEndpoint(_ context.Context, _ s3.EndpointParameters) (transport.Endpoint, error) {
	endpoint := r.Endpoint

	if r.PathStyle {
		endpoint = endpoint.JoinPath(
			r.Bucket,
		)
	}

	return transport.Endpoint{
		URI: *endpoint,
	}, nil
}
