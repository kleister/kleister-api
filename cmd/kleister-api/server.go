package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/kleister/kleister-api/config"
	"github.com/kleister/kleister-api/router"
	"github.com/kleister/kleister-api/shared/s3client"
	"github.com/urfave/cli"
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sync/errgroup"
)

var (
	defaultAddr = ":8080"
)

// Server provides the sub-command to start the API server.
func Server() cli.Command {
	return cli.Command{
		Name:  "server",
		Usage: "Start the Kleister API",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "db-driver",
				Value:       "mysql",
				Usage:       "Database driver selection",
				EnvVar:      "KLEISTER_DB_DRIVER",
				Destination: &config.Database.Driver,
			},
			cli.StringFlag{
				Name:        "db-name",
				Value:       "kleister",
				Usage:       "Name for database connection",
				EnvVar:      "KLEISTER_DB_NAME",
				Destination: &config.Database.Name,
			},
			cli.StringFlag{
				Name:        "db-username",
				Value:       "root",
				Usage:       "Username for database connection",
				EnvVar:      "KLEISTER_DB_USERNAME",
				Destination: &config.Database.Username,
			},
			cli.StringFlag{
				Name:        "db-password",
				Value:       "root",
				Usage:       "Password for database connection",
				EnvVar:      "KLEISTER_DB_PASSWORD",
				Destination: &config.Database.Password,
			},
			cli.StringFlag{
				Name:        "db-host",
				Value:       "localhost:3306",
				Usage:       "Host for database connection",
				EnvVar:      "KLEISTER_DB_HOST",
				Destination: &config.Database.Host,
			},
			cli.StringFlag{
				Name:        "host",
				Value:       "http://localhost:8080",
				Usage:       "External access to server",
				EnvVar:      "KLEISTER_SERVER_HOST",
				Destination: &config.Server.Host,
			},
			cli.StringFlag{
				Name:        "addr",
				Value:       defaultAddr,
				Usage:       "Address to bind the server",
				EnvVar:      "KLEISTER_SERVER_ADDR",
				Destination: &config.Server.Addr,
			},
			cli.StringFlag{
				Name:        "root",
				Value:       "/",
				Usage:       "Root folder of the app",
				EnvVar:      "KLEISTER_SERVER_ROOT",
				Destination: &config.Server.Root,
			},
			cli.StringFlag{
				Name:        "storage",
				Value:       "storage/",
				Usage:       "Folder for storing uploads",
				EnvVar:      "KLEISTER_SERVER_STORAGE",
				Destination: &config.Server.Storage,
			},
			cli.BoolFlag{
				Name:        "pprof",
				Usage:       "Enable pprof debugging server",
				EnvVar:      "KLEISTER_SERVER_PPROF",
				Destination: &config.Server.Pprof,
			},
			cli.StringFlag{
				Name:        "cert",
				Value:       "",
				Usage:       "Path to SSL cert",
				EnvVar:      "KLEISTER_SERVER_CERT",
				Destination: &config.Server.Cert,
			},
			cli.StringFlag{
				Name:        "key",
				Value:       "",
				Usage:       "Path to SSL key",
				EnvVar:      "KLEISTER_SERVER_KEY",
				Destination: &config.Server.Key,
			},
			cli.BoolFlag{
				Name:        "letsencrypt",
				Usage:       "Enable Let's Encrypt SSL",
				EnvVar:      "KLEISTER_SERVER_LETSENCRYPT",
				Destination: &config.Server.LetsEncrypt,
			},
			cli.BoolFlag{
				Name:   "strict-curves",
				Usage:  "Use strict SSL curves",
				EnvVar: "KLEISTER_STRICT_CURVES",
			},
			cli.BoolFlag{
				Name:   "strict-ciphers",
				Usage:  "Use strict SSL ciphers",
				EnvVar: "KLEISTER_STRICT_CIPHERS",
			},
			cli.DurationFlag{
				Name:        "expire",
				Value:       time.Hour * 24,
				Usage:       "Session expire duration",
				EnvVar:      "KLEISTER_SESSION_EXPIRE",
				Destination: &config.Session.Expire,
			},
			cli.StringSliceFlag{
				Name:   "admin-user",
				Value:  &cli.StringSlice{},
				Usage:  "Enforce user as an admin",
				EnvVar: "KLEISTER_ADMIN_USERS",
			},
			cli.BoolTFlag{
				Name:        "admin-create",
				Usage:       "Create an initial admin user",
				EnvVar:      "KLEISTER_ADMIN_CREATE",
				Destination: &config.Admin.Create,
			},
			cli.BoolFlag{
				Name:        "s3-enabled",
				Usage:       "Enable S3 uploads",
				EnvVar:      "KLEISTER_S3_ENABLED",
				Destination: &config.S3.Enabled,
			},
			cli.StringFlag{
				Name:        "s3-endpoint",
				Value:       "",
				Usage:       "S3 API endpoint",
				EnvVar:      "KLEISTER_S3_ENDPOINT",
				Destination: &config.S3.Endpoint,
			},
			cli.StringFlag{
				Name:        "s3-bucket",
				Value:       "kleister",
				Usage:       "S3 bucket name",
				EnvVar:      "KLEISTER_S3_BUCKET",
				Destination: &config.S3.Bucket,
			},
			cli.StringFlag{
				Name:        "s3-region",
				Value:       "us-east-1",
				Usage:       "S3 region name",
				EnvVar:      "KLEISTER_S3_REGION",
				Destination: &config.S3.Region,
			},
			cli.StringFlag{
				Name:        "s3-access",
				Value:       "",
				Usage:       "S3 public key",
				EnvVar:      "KLEISTER_S3_ACCESS_KEY",
				Destination: &config.S3.Access,
			},
			cli.StringFlag{
				Name:        "s3-secret",
				Value:       "",
				Usage:       "S3 secret key",
				EnvVar:      "KLEISTER_S3_SECRET_KEY",
				Destination: &config.S3.Secret,
			},
			cli.BoolFlag{
				Name:        "s3-path-style",
				Usage:       "S3 path style",
				EnvVar:      "KLEISTER_S3_PATH_STYLE",
				Destination: &config.S3.PathStyle,
			},
		},
		Before: func(c *cli.Context) error {
			if len(c.StringSlice("admin-user")) > 0 {
				// StringSliceFlag doesn't support Destination
				config.Admin.Users = c.StringSlice("admin-user")
			}

			if config.S3.Enabled {
				_, err := s3client.New().List()

				if err != nil {
					return fmt.Errorf("Failed to connect to S3. %s", err)
				}
			}

			return nil
		},
		Action: func(c *cli.Context) {
			logrus.Infof("Starting API on %s", config.Server.Addr)

			if config.Server.LetsEncrypt || (config.Server.Cert != "" && config.Server.Key != "") {
				cfg := &tls.Config{
					PreferServerCipherSuites: true,
					MinVersion:               tls.VersionTLS12,
				}

				if c.Bool("strict-curves") {
					cfg.CurvePreferences = []tls.CurveID{
						tls.CurveP521,
						tls.CurveP384,
						tls.CurveP256,
					}
				}

				if c.Bool("strict-ciphers") {
					cfg.CipherSuites = []uint16{
						tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
						tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
						tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
						tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
					}
				}

				if config.Server.LetsEncrypt {
					if config.Server.Addr != defaultAddr {
						logrus.Infof("With Let's Encrypt bind port have been overwritten!")
					}

					parsed, err := url.Parse(config.Server.Host)

					if err != nil {
						logrus.Fatal("Failed to parse host name. %s", err)
					}

					certManager := autocert.Manager{
						Prompt:     autocert.AcceptTOS,
						HostPolicy: autocert.HostWhitelist(parsed.Host),
						Cache:      autocert.DirCache(path.Join(config.Server.Storage, "certs")),
					}

					cfg.GetCertificate = certManager.GetCertificate

					var (
						g errgroup.Group
					)

					splitAddr := strings.SplitN(config.Server.Addr, ":", 2)
					logrus.Infof("Starting on %s:80 and %s:443", splitAddr[0], splitAddr[0])

					g.Go(func() error {
						return http.ListenAndServe(
							fmt.Sprintf("%s:80", splitAddr[0]),
							http.HandlerFunc(redirect),
						)
					})

					g.Go(func() error {
						return startServer(&http.Server{
							Addr:         fmt.Sprintf("%s:443", splitAddr[0]),
							Handler:      router.Load(),
							ReadTimeout:  5 * time.Second,
							WriteTimeout: 10 * time.Second,
							TLSConfig:    cfg,
						})
					})

					if err := g.Wait(); err != nil {
						logrus.Fatal(err)
					}
				} else {
					logrus.Infof("Starting on %s", config.Server.Addr)

					cert, err := tls.LoadX509KeyPair(
						config.Server.Cert,
						config.Server.Key,
					)

					if err != nil {
						logrus.Fatal("Failed to load SSL certificates. %s", err)
					}

					cfg.Certificates = []tls.Certificate{
						cert,
					}

					server := &http.Server{
						Addr:         config.Server.Addr,
						Handler:      router.Load(),
						ReadTimeout:  5 * time.Second,
						WriteTimeout: 10 * time.Second,
						TLSConfig:    cfg,
					}

					if err := startServer(server); err != nil {
						logrus.Fatal(err)
					}
				}
			} else {
				logrus.Infof("Starting on %s", config.Server.Addr)

				server := &http.Server{
					Addr:         config.Server.Addr,
					Handler:      router.Load(),
					ReadTimeout:  5 * time.Second,
					WriteTimeout: 10 * time.Second,
				}

				if err := startServer(server); err != nil {
					logrus.Fatal(err)
				}
			}
		},
	}
}

func redirect(w http.ResponseWriter, req *http.Request) {
	target := "https://" + req.Host + req.URL.Path

	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}

	logrus.Debugf("Redirecting to %s", target)
	http.Redirect(w, req, target, http.StatusTemporaryRedirect)
}
