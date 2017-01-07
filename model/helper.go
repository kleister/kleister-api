package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/kleister/kleister-api/assets"
	"github.com/kleister/kleister-api/config"
	"github.com/kleister/kleister-api/shared/s3client"
	"github.com/vincent-petithory/dataurl"
)

// RandomHash generates a random string for unique slugs.
func RandomHash() string {
	hash := md5.Sum(
		[]byte(
			strconv.FormatInt(time.Now().Unix(), 10),
		),
	)
	return hex.EncodeToString(hash[:])
}

// AttachmentDefault returns checksum and URL for default images.
func AttachmentDefault(kind string) (string, string) {
	content, err := assets.ReadFile(
		fmt.Sprintf(
			"images/default-%s.png",
			kind,
		),
	)

	if err != nil {
		return "", ""
	}

	check := md5.Sum(
		content,
	)

	hash := hex.EncodeToString(
		check[:],
	)

	url := fmt.Sprintf(
		"%s/assets/images/default-%s.png",
		config.Server.Host,
		kind,
	)

	return url, hash
}

// AttachmentURL generates an absolute URL to attachments.
func AttachmentURL(attachment string) string {
	if config.S3.Enabled {
		if config.S3.Endpoint == "" {
			return fmt.Sprintf(
				"https://s3-%s.amazonaws.com/%s/%s",
				config.S3.Region,
				config.S3.Bucket,
				attachment,
			)
		}

		return fmt.Sprintf(
			"%s/%s/%s",
			config.S3.Endpoint,
			config.S3.Bucket,
			attachment,
		)
	}

	return fmt.Sprintf(
		"%s/%s/%s",
		config.Server.Host,
		"storage",
		attachment,
	)
}

// AttachmentHash calculates checksums and media-types.
func AttachmentHash(upload *dataurl.DataURL) (string, string) {
	check := md5.Sum(
		upload.Data,
	)

	hash := hex.EncodeToString(
		check[:],
	)

	return hash, upload.MediaType.String()
}

// AttachmentDelete removes an attachment from the storage.
func AttachmentDelete(dest string) error {
	if config.S3.Enabled {
		_, err := s3client.New().Delete(
			dest,
		)

		return err
	}

	if err := os.Remove(dest); err != nil {
		return err
	}

	return nil
}

// AttachmentUpload stores an attachment to the storage.
func AttachmentUpload(dest string, contentType string, content []byte) error {
	if config.S3.Enabled {
		_, err := s3client.New().Upload(
			dest,
			contentType,
			content,
		)

		return err
	}

	dir := filepath.Dir(
		dest,
	)

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	if err := ioutil.WriteFile(dest, content, 0644); err != nil {
		return err
	}

	return nil
}
