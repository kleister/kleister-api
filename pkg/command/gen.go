package command

import (
	"github.com/spf13/cobra"
)

var (
	genCmd = &cobra.Command{
		Use:   "gen",
		Short: "Generate required stuff",
	}

	genCertCmd = &cobra.Command{
		Use:   "cert",
		Short: "Generate SSL certificates",
		Run:   genCertAction,
		Args:  cobra.NoArgs,
	}
)

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.AddCommand(genCertCmd)

	// serverCmd.PersistentFlags().String("metrics-addr", defaultMetricsAddr, "Address to bind the metrics")
	// viper.SetDefault("metrics.addr", defaultMetricsAddr)
	// viper.BindPFlag("metrics.addr", serverCmd.PersistentFlags().Lookup("metrics-addr"))

	// serverCmd.PersistentFlags().String("metrics-token", "", "Token to make metrics secure")
	// viper.SetDefault("metrics.token", "")
	// viper.BindPFlag("metrics.token", serverCmd.PersistentFlags().Lookup("metrics-token"))

	// serverCmd.PersistentFlags().String("server-addr", defaultServerAddr, "Address to bind the server")
	// viper.SetDefault("server.addr", defaultServerAddr)
	// viper.BindPFlag("server.addr", serverCmd.PersistentFlags().Lookup("server-addr"))

	// serverCmd.PersistentFlags().Bool("server-pprof", defaultServerPprof, "Enable pprof debugging")
	// viper.SetDefault("server.pprof", defaultServerPprof)
	// viper.BindPFlag("server.pprof", serverCmd.PersistentFlags().Lookup("server-pprof"))

	// serverCmd.PersistentFlags().String("server-root", defaultServerRoot, "Root path of the server")
	// viper.SetDefault("server.root", defaultServerRoot)
	// viper.BindPFlag("server.root", serverCmd.PersistentFlags().Lookup("server-root"))

	// serverCmd.PersistentFlags().String("server-host", defaultServerHost, "External access to server")
	// viper.SetDefault("server.host", defaultServerHost)
	// viper.BindPFlag("server.host", serverCmd.PersistentFlags().Lookup("server-host"))

	// serverCmd.PersistentFlags().String("server-cert", defaultServerCert, "Path to cert for SSL encryption")
	// viper.SetDefault("server.cert", defaultServerCert)
	// viper.BindPFlag("server.cert", serverCmd.PersistentFlags().Lookup("server-cert"))

	// serverCmd.PersistentFlags().String("server-key", defaultServerKey, "Path to key for SSL encryption")
	// viper.SetDefault("server.key", defaultServerKey)
	// viper.BindPFlag("server.key", serverCmd.PersistentFlags().Lookup("server-key"))

	// serverCmd.PersistentFlags().Bool("strict-curves", defaultServerStrictCurves, "Use strict SSL curves")
	// viper.SetDefault("server.strict_curves", defaultServerStrictCurves)
	// viper.BindPFlag("server.strict_curves", serverCmd.PersistentFlags().Lookup("strict-curves"))

	// serverCmd.PersistentFlags().Bool("strict-ciphers", defaultServerStrictCiphers, "Use strict SSL ciphers")
	// viper.SetDefault("server.strict_ciphers", defaultServerStrictCiphers)
	// viper.BindPFlag("server.strict_ciphers", serverCmd.PersistentFlags().Lookup("strict-ciphers"))

	// serverCmd.PersistentFlags().String("templates-path", defaultServerTemplates, "Path for overriding templates")
	// viper.SetDefault("server.templates", defaultServerTemplates)
	// viper.BindPFlag("server.templates", serverCmd.PersistentFlags().Lookup("templates-path"))

	// serverCmd.PersistentFlags().String("errors-path", defaultServerErrors, "Path for overriding errors")
	// viper.SetDefault("server.errors", defaultServerErrors)
	// viper.BindPFlag("server.errors", serverCmd.PersistentFlags().Lookup("errors-path"))
}

// func GenCertFlags(_ *config.Config) []cli.Flag {
// 	return []cli.Flag{
// 		&cli.StringSliceFlag{
// 			Name:    "cert-host",
// 			Value:   cli.NewStringSlice("kleister-api"),
// 			Usage:   "List of cert hosts",
// 			EnvVars: []string{"KLEISTER_API_CERT_HOSTS"},
// 		},
// 		&cli.StringFlag{
// 			Name:    "cert-org",
// 			Value:   "Kleister",
// 			Usage:   "Org for certificate",
// 			EnvVars: []string{"KLEISTER_API_ECDSA_CURVE"},
// 		},
// 		&cli.StringFlag{
// 			Name:    "cert-name",
// 			Value:   "API",
// 			Usage:   "Name for certificate",
// 			EnvVars: []string{"KLEISTER_API_ECDSA_CURVE"},
// 		},
// 		&cli.StringFlag{
// 			Name:    "ecdsa-curve",
// 			Value:   "",
// 			Usage:   "ECDSA curve to use",
// 			EnvVars: []string{"KLEISTER_API_ECDSA_CURVE"},
// 		},
// 		&cli.IntFlag{
// 			Name:    "rsa-bits",
// 			Value:   4096,
// 			Usage:   "Size of RSA to gen",
// 			EnvVars: []string{"KLEISTER_API_ECDSA_CURVE"},
// 		},
// 		&cli.DurationFlag{
// 			Name:    "valid-for",
// 			Value:   365 * 24 * time.Hour,
// 			Usage:   "Duration for the cert",
// 			EnvVars: []string{"KLEISTER_API_VALID_FOR"},
// 		},
// 		&cli.StringFlag{
// 			Name:    "server-cert",
// 			Value:   "server.crt",
// 			Usage:   "Path to SSL cert",
// 			EnvVars: []string{"KLEISTER_API_SERVER_CERT"},
// 		},
// 		&cli.StringFlag{
// 			Name:    "server-key",
// 			Value:   "server.key",
// 			Usage:   "Path to SSL key",
// 			EnvVars: []string{"KLEISTER_API_SERVER_KEY"},
// 		},
// 	}
// }

func genCertAction(_ *cobra.Command, _ []string) {
	// priv, err := parseEcdsaCurve(c)

	// if err != nil {
	// 	log.Error().
	// 		Err(err).
	// 		Msg("failed to gen private key")

	// 	return err
	// }

	// notBefore := time.Now()
	// notAfter := notBefore.Add(c.Duration("valid-for"))

	// serialNumber, err := buildSerialNumber()

	// if err != nil {
	// 	log.Error().
	// 		Err(err).
	// 		Msg("failed to gen serial number")

	// 	return err
	// }

	// template := x509.Certificate{
	// 	SerialNumber: serialNumber,
	// 	Subject: pkix.Name{
	// 		Organization: []string{c.String("cert-org")},
	// 		CommonName:   c.String("cert-name"),
	// 	},
	// 	NotBefore:             notBefore,
	// 	NotAfter:              notAfter,
	// 	KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	// 	ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	// 	BasicConstraintsValid: true,
	// 	IsCA:                  true,
	// }

	// for _, host := range c.StringSlice("cert-host") {
	// 	if ip := net.ParseIP(host); ip != nil {
	// 		template.IPAddresses = append(
	// 			template.IPAddresses,
	// 			ip,
	// 		)
	// 	} else {
	// 		template.DNSNames = append(
	// 			template.DNSNames,
	// 			host,
	// 		)
	// 	}
	// }

	// der, err := x509.CreateCertificate(
	// 	rand.Reader,
	// 	&template,
	// 	&template,
	// 	extractPublicKey(priv),
	// 	priv,
	// )

	// if err != nil {
	// 	log.Error().
	// 		Err(err).
	// 		Msg("failed to create certificate")

	// 	return err
	// }

	// crt, err := os.OpenFile(
	// 	c.String("server-cert"),
	// 	os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
	// 	0o600,
	// )

	// if err != nil {
	// 	log.Error().
	// 		Err(err).
	// 		Str("cert", c.String("server-cert")).
	// 		Msg("failed to open cert file")

	// 	return err
	// }

	// if err := pem.Encode(
	// 	crt,
	// 	publicEncodeBlock(der),
	// ); err != nil {
	// 	log.Error().
	// 		Err(err).
	// 		Msg("failed to encode cert")

	// 	return err
	// }

	// if err := crt.Close(); err != nil {
	// 	log.Error().
	// 		Err(err).
	// 		Str("cert", c.String("server-cert")).
	// 		Msg("failed to close cert file")

	// 	return err
	// }

	// key, err := os.OpenFile(
	// 	c.String("server-key"),
	// 	os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
	// 	0o600,
	// )

	// if err != nil {
	// 	log.Error().
	// 		Err(err).
	// 		Str("key", c.String("server-key")).
	// 		Msg("failed to open key file")

	// 	return err
	// }

	// if err := pem.Encode(
	// 	key,
	// 	privateEncodeBlock(priv),
	// ); err != nil {
	// 	log.Error().
	// 		Err(err).
	// 		Msg("failed to encode key")

	// 	return err
	// }

	// if err := key.Close(); err != nil {
	// 	log.Error().
	// 		Err(err).
	// 		Str("key", c.String("server-key")).
	// 		Msg("failed to close key file")

	// 	return err
	// }

	// log.Info().
	// 	Str("cert", c.String("server-cert")).
	// 	Str("key", c.String("server-key")).
	// 	Msg("successfully generated")

}

// func parseEcdsaCurve(c *cli.Context) (interface{}, error) {
// 	switch c.String("ecdsa-curve") {
// 	case "":
// 		return rsa.GenerateKey(rand.Reader, c.Int("rsa-bits"))
// 	case "P224":
// 		return ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
// 	case "P256":
// 		return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
// 	case "P384":
// 		return ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
// 	case "P521":
// 		return ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
// 	default:
// 		return nil, fmt.Errorf("unrecognized elliptic curve: %q", c.String("ecdsa-curve"))
// 	}
// }

// func buildSerialNumber() (*big.Int, error) {
// 	return rand.Int(
// 		rand.Reader,
// 		new(
// 			big.Int,
// 		).Lsh(
// 			big.NewInt(1),
// 			128,
// 		),
// 	)
// }

// func publicEncodeBlock(der []byte) *pem.Block {
// 	return &pem.Block{
// 		Type:  "CERTIFICATE",
// 		Bytes: der,
// 	}
// }

// func privateEncodeBlock(priv interface{}) *pem.Block {
// 	switch k := priv.(type) {
// 	case *rsa.PrivateKey:
// 		return &pem.Block{
// 			Type:  "RSA PRIVATE KEY",
// 			Bytes: x509.MarshalPKCS1PrivateKey(k),
// 		}
// 	case *ecdsa.PrivateKey:
// 		b, err := x509.MarshalECPrivateKey(k)

// 		if err != nil {
// 			log.Error().
// 				Err(err).
// 				Msg("unable to marshal ECDSA key")
// 		}

// 		return &pem.Block{
// 			Type:  "EC PRIVATE KEY",
// 			Bytes: b,
// 		}
// 	default:
// 		return nil
// 	}
// }

// func extractPublicKey(priv interface{}) interface{} {
// 	switch k := priv.(type) {
// 	case *rsa.PrivateKey:
// 		return &k.PublicKey
// 	case *ecdsa.PrivateKey:
// 		return &k.PublicKey
// 	default:
// 		return nil
// 	}
// }
