package main

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path"
	"strings"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/router"
	"github.com/kleister/kleister-api/pkg/s3client"
	"github.com/kleister/kleister-api/pkg/storage"
	"github.com/oklog/oklog/pkg/group"
	"golang.org/x/crypto/acme/autocert"
	"gopkg.in/urfave/cli.v2"
)

var (
	defaultAddr = "0.0.0.0:8080"
)

// Server provides the sub-command to start the server.
func Server() *cli.Command {
	return &cli.Command{
		Name:  "server",
		Usage: "start integrated server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "server-host",
				Value:       "http://localhost:8080",
				Usage:       "external access to server",
				EnvVars:     []string{"KLEISTER_SERVER_HOST"},
				Destination: &config.Server.Host,
			},
			&cli.StringFlag{
				Name:        "server-addr",
				Value:       defaultAddr,
				Usage:       "address to bind the server",
				EnvVars:     []string{"KLEISTER_SERVER_ADDR"},
				Destination: &config.Server.Addr,
			},
			&cli.StringFlag{
				Name:        "server-root",
				Value:       "/",
				Usage:       "root folder of the app",
				EnvVars:     []string{"KLEISTER_SERVER_ROOT"},
				Destination: &config.Server.Root,
			},
			&cli.StringFlag{
				Name:        "storage-path",
				Value:       "storage/",
				Usage:       "folder for storing uploads",
				EnvVars:     []string{"KLEISTER_SERVER_STORAGE"},
				Destination: &config.Server.Storage,
			},
			&cli.StringFlag{
				Name:        "assets-path",
				Value:       "",
				Usage:       "path to custom assets and templates",
				EnvVars:     []string{"KLEISTER_SERVER_ASSETS"},
				Destination: &config.Server.Assets,
			},
			&cli.BoolFlag{
				Name:        "enable-pprof",
				Value:       false,
				Usage:       "enable pprof debugging server",
				EnvVars:     []string{"KLEISTER_SERVER_PPROF"},
				Destination: &config.Server.Pprof,
			},
			&cli.BoolFlag{
				Name:        "enable-prometheus",
				Value:       false,
				Usage:       "enable prometheus exporter",
				EnvVars:     []string{"KLEISTER_SERVER_PROMETHEUS"},
				Destination: &config.Server.Prometheus,
			},
			&cli.StringFlag{
				Name:        "server-cert",
				Value:       "",
				Usage:       "path to ssl cert",
				EnvVars:     []string{"KLEISTER_SERVER_CERT"},
				Destination: &config.Server.Cert,
			},
			&cli.StringFlag{
				Name:        "server-key",
				Value:       "",
				Usage:       "path to ssl key",
				EnvVars:     []string{"KLEISTER_SERVER_KEY"},
				Destination: &config.Server.Key,
			},
			&cli.BoolFlag{
				Name:        "enable-letsencrypt",
				Value:       false,
				Usage:       "enable let's encrypt ssl",
				EnvVars:     []string{"KLEISTER_SERVER_LETSENCRYPT"},
				Destination: &config.Server.LetsEncrypt,
			},
			&cli.BoolFlag{
				Name:        "strict-curves",
				Value:       false,
				Usage:       "use strict ssl curves",
				EnvVars:     []string{"KLEISTER_STRICT_CURVES"},
				Destination: &config.Server.StrictCurves,
			},
			&cli.BoolFlag{
				Name:        "strict-ciphers",
				Value:       false,
				Usage:       "use strict ssl ciphers",
				EnvVars:     []string{"KLEISTER_STRICT_CIPHERS"},
				Destination: &config.Server.StrictCiphers,
			},
			&cli.DurationFlag{
				Name:        "session-expire",
				Value:       time.Hour * 24,
				Usage:       "session expire duration",
				EnvVars:     []string{"KLEISTER_SESSION_EXPIRE"},
				Destination: &config.Session.Expire,
			},
			&cli.StringSliceFlag{
				Name:    "admin-user",
				Value:   &cli.StringSlice{},
				Usage:   "enforce user as an admin",
				EnvVars: []string{"KLEISTER_ADMIN_USERS"},
			},
			&cli.BoolFlag{
				Name:        "admin-create",
				Value:       true,
				Usage:       "create an initial admin user",
				EnvVars:     []string{"KLEISTER_ADMIN_CREATE"},
				Destination: &config.Admin.Create,
			},
			&cli.BoolFlag{
				Name:        "s3-enabled",
				Value:       false,
				Usage:       "enable s3 uploads",
				EnvVars:     []string{"KLEISTER_S3_ENABLED"},
				Destination: &config.S3.Enabled,
			},
			&cli.StringFlag{
				Name:        "s3-endpoint",
				Value:       "",
				Usage:       "s3 api endpoint",
				EnvVars:     []string{"KLEISTER_S3_ENDPOINT"},
				Destination: &config.S3.Endpoint,
			},
			&cli.StringFlag{
				Name:        "s3-bucket",
				Value:       "kleister",
				Usage:       "s3 bucket name",
				EnvVars:     []string{"KLEISTER_S3_BUCKET"},
				Destination: &config.S3.Bucket,
			},
			&cli.StringFlag{
				Name:        "s3-region",
				Value:       "us-east-1",
				Usage:       "s3 region name",
				EnvVars:     []string{"KLEISTER_S3_REGION"},
				Destination: &config.S3.Region,
			},
			&cli.StringFlag{
				Name:        "s3-access",
				Value:       "",
				Usage:       "s3 public key",
				EnvVars:     []string{"KLEISTER_S3_ACCESS_KEY"},
				Destination: &config.S3.Access,
			},
			&cli.StringFlag{
				Name:        "s3-secret",
				Value:       "",
				Usage:       "s3 secret key",
				EnvVars:     []string{"KLEISTER_S3_SECRET_KEY"},
				Destination: &config.S3.Secret,
			},
			&cli.BoolFlag{
				Name:        "s3-pathstyle",
				Value:       false,
				Usage:       "s3 path style",
				EnvVars:     []string{"KLEISTER_S3_PATH_STYLE"},
				Destination: &config.S3.PathStyle,
			},
		},
		Before: func(c *cli.Context) error {
			if len(c.StringSlice("admin-user")) > 0 {
				// StringSliceFlag doesn't support Destination
				config.Admin.Users = c.StringSlice("admin-user")
			}

			return nil
		},
		Action: func(c *cli.Context) error {
			logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))

			switch strings.ToLower(config.LogLevel) {
			case "debug":
				logger = level.NewFilter(logger, level.AllowDebug())
			case "warn":
				logger = level.NewFilter(logger, level.AllowWarn())
			case "error":
				logger = level.NewFilter(logger, level.AllowError())
			default:
				logger = level.NewFilter(logger, level.AllowInfo())
			}

			logger = log.WithPrefix(logger,
				"app", c.App.Name,
				"ts", log.DefaultTimestampUTC,
			)

			if config.S3.Enabled {
				if err := s3client.New().Ping(); err != nil {
					level.Error(logger).Log(
						"msg", "failed to connect to s3",
						"err", err,
					)

					return err
				}
			}

			store, err := storage.Load(logger)

			if err != nil {
				level.Error(logger).Log(
					"msg", "failed to initialize database",
					"err", err,
				)

				return err
			}

			var gr group.Group

			if config.Server.LetsEncrypt || (config.Server.Cert != "" && config.Server.Key != "") {
				cfg := &tls.Config{
					PreferServerCipherSuites: true,
					MinVersion:               tls.VersionTLS12,
				}

				if config.Server.StrictCurves {
					cfg.CurvePreferences = []tls.CurveID{
						tls.CurveP521,
						tls.CurveP384,
						tls.CurveP256,
					}
				}

				if config.Server.StrictCiphers {
					cfg.CipherSuites = []uint16{
						tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
						tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
						tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
						tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
					}
				}

				if config.Server.LetsEncrypt {
					if config.Server.Addr != defaultAddr {
						level.Info(logger).Log(
							"msg", "enabled let's encrypt, overwriting the port",
						)
					}

					parsed, err := url.Parse(config.Server.Host)

					if err != nil {
						level.Error(logger).Log(
							"msg", "failed to parse host",
							"err", err,
						)

						return err
					}

					certManager := autocert.Manager{
						Prompt:     autocert.AcceptTOS,
						HostPolicy: autocert.HostWhitelist(parsed.Host),
						Cache:      autocert.DirCache(path.Join(config.Server.Storage, "certs")),
					}

					cfg.GetCertificate = certManager.GetCertificate

					splitAddr := strings.SplitN(
						config.Server.Addr,
						":",
						2,
					)

					{
						addr := net.JoinHostPort(splitAddr[0], "80")

						server := &http.Server{
							Addr:         addr,
							Handler:      redirect(logger),
							ReadTimeout:  5 * time.Second,
							WriteTimeout: 10 * time.Second,
						}

						gr.Add(func() error {
							level.Info(logger).Log(
								"msg", "starting http server",
								"addr", addr,
							)

							return server.ListenAndServe()
						}, func(reason error) {
							ctx, cancel := context.WithTimeout(context.Background(), time.Second)
							defer cancel()

							if err := server.Shutdown(ctx); err != nil {
								level.Error(logger).Log(
									"msg", "failed to shutdown http server gracefully",
									"err", err,
								)

								return
							}

							level.Info(logger).Log(
								"msg", "http server shutdown gracefully",
								"reason", reason,
							)
						})
					}

					{
						addr := net.JoinHostPort(splitAddr[0], "443")

						server := &http.Server{
							Addr:         addr,
							Handler:      router.Load(store, logger),
							ReadTimeout:  5 * time.Second,
							WriteTimeout: 10 * time.Second,
							TLSConfig:    cfg,
						}

						gr.Add(func() error {
							level.Info(logger).Log(
								"msg", "starting https server",
								"addr", addr,
							)

							return server.ListenAndServeTLS("", "")
						}, func(reason error) {
							ctx, cancel := context.WithTimeout(context.Background(), time.Second)
							defer cancel()

							if err := server.Shutdown(ctx); err != nil {
								level.Error(logger).Log(
									"msg", "failed to shutdown https server gracefully",
									"err", err,
								)

								return
							}

							level.Info(logger).Log(
								"msg", "https server shutdown gracefully",
								"reason", reason,
							)
						})
					}
				} else {
					cert, err := tls.LoadX509KeyPair(
						config.Server.Cert,
						config.Server.Key,
					)

					if err != nil {
						level.Error(logger).Log(
							"msg", "failed to load certificates",
							"err", err,
						)

						return err
					}

					cfg.Certificates = []tls.Certificate{
						cert,
					}

					{
						server := &http.Server{
							Addr:         config.Server.Addr,
							Handler:      router.Load(store, logger),
							ReadTimeout:  5 * time.Second,
							WriteTimeout: 10 * time.Second,
							TLSConfig:    cfg,
						}

						gr.Add(func() error {
							level.Info(logger).Log(
								"msg", "starting https server",
								"addr", config.Server.Addr,
							)

							return server.ListenAndServeTLS("", "")
						}, func(reason error) {
							ctx, cancel := context.WithTimeout(context.Background(), time.Second)
							defer cancel()

							if err := server.Shutdown(ctx); err != nil {
								level.Error(logger).Log(
									"msg", "failed to shutdown https server gracefully",
									"err", err,
								)

								return
							}

							level.Info(logger).Log(
								"msg", "https server shutdown gracefully",
								"reason", reason,
							)
						})
					}
				}
			} else {
				{
					server := &http.Server{
						Addr:         config.Server.Addr,
						Handler:      router.Load(store, logger),
						ReadTimeout:  5 * time.Second,
						WriteTimeout: 10 * time.Second,
					}

					gr.Add(func() error {
						level.Info(logger).Log(
							"msg", "starting http server",
							"addr", config.Server.Addr,
						)

						return server.ListenAndServe()
					}, func(reason error) {
						ctx, cancel := context.WithTimeout(context.Background(), time.Second)
						defer cancel()

						if err := server.Shutdown(ctx); err != nil {
							level.Error(logger).Log(
								"msg", "failed to shutdown http server gracefully",
								"err", err,
							)

							return
						}

						level.Info(logger).Log(
							"msg", "http server shutdown gracefully",
							"reason", reason,
						)
					})
				}
			}

			{
				gr.Add(func() error {
					stop := make(chan os.Signal, 1)
					signal.Notify(stop, os.Interrupt)

					<-stop

					return nil
				}, func(err error) {

				})
			}

			return gr.Run()
		},
	}
}

func redirect(logger log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		target := strings.Join(
			[]string{
				"https://",
				r.Host,
				r.URL.Path,
			},
			"",
		)

		if len(r.URL.RawQuery) > 0 {
			target += "?" + r.URL.RawQuery
		}

		level.Debug(logger).Log(
			"msg", "redirecting to https",
			"target", target,
		)

		http.Redirect(w, r, target, http.StatusPermanentRedirect)
	})
}
