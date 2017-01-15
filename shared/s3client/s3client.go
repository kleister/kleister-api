package s3client

import (
	"bytes"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/kleister/kleister-api/config"
)

// S3Client is a simple wrapper around a real S3 client.
type S3Client struct {
	client *s3.S3
}

// List retrieves a list of available objects in the bucket.
func (u *S3Client) List() (*s3.ListObjectsOutput, error) {
	params := &s3.ListObjectsInput{
		Bucket: aws.String(config.S3.Bucket),
	}

	return u.client.ListObjects(
		params,
	)
}

// Upload stores an attachment within the defined S3 bucket.
func (u *S3Client) Upload(path string, ctype string, content []byte) (*s3.PutObjectOutput, error) {
	params := &s3.PutObjectInput{
		ACL:         aws.String("public-read"),
		Bucket:      aws.String(config.S3.Bucket),
		Key:         aws.String(path),
		ContentType: aws.String(ctype),
		Body:        bytes.NewReader(content),
	}

	return u.client.PutObject(
		params,
	)
}

// Delete removes an attachment from the defined S3 bucket.
func (u *S3Client) Delete(path string) (*s3.DeleteObjectOutput, error) {
	params := &s3.DeleteObjectInput{
		Bucket: aws.String(config.S3.Bucket),
		Key:    aws.String(path),
	}

	return u.client.DeleteObject(
		params,
	)
}

// New initializes a new S3 client connection based on config.
func New() *S3Client {
	var (
		cfg *aws.Config
	)

	if config.S3.Endpoint != "" {
		cfg = &aws.Config{
			Endpoint:         aws.String(config.S3.Endpoint),
			DisableSSL:       aws.Bool(strings.HasPrefix(config.S3.Endpoint, "http://")),
			Region:           aws.String(config.S3.Region),
			S3ForcePathStyle: aws.Bool(config.S3.PathStyle),
		}
	} else {
		cfg = &aws.Config{
			Region:           aws.String(config.S3.Region),
			S3ForcePathStyle: aws.Bool(config.S3.PathStyle),
		}
	}

	if config.S3.Access != "" && config.S3.Secret != "" {
		cfg.Credentials = credentials.NewStaticCredentials(
			config.S3.Access,
			config.S3.Secret,
			"",
		)
	}

	sess, _ := session.NewSession()

	return &S3Client{
		client: s3.New(
			sess,
			cfg,
		),
	}
}
