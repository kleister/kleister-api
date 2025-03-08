package command

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	genCmd = &cobra.Command{
		Use:   "gen",
		Short: "Generate required stuff",
		Args:  cobra.NoArgs,
	}

	genCertCmd = &cobra.Command{
		Use:   "cert",
		Short: "Generate SSL certificates",
		Run:   genCertAction,
		Args:  cobra.NoArgs,
	}

	defaultCertGenCertHosts  = []string{"kleister"}
	defaultCertGenCertOrg    = "Kleister"
	defaultCertGenCertName   = "API"
	defaultCertGenEcdsaCurve = ""
	defaultCertGenRSABits    = 4096
	defaultCertGenValidFor   = 365 * 24 * time.Hour
	defaultCertGenOutputCert = "server.crt"
	defaultCertGenOutputKey  = "server.key"
)

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.AddCommand(genCertCmd)

	genCertCmd.Flags().StringSlice("cert-hosts", defaultCertGenCertHosts, "List of cert hosts")
	viper.SetDefault("cert.hosts", defaultCertGenCertHosts)
	_ = viper.BindPFlag("cert.hosts", genCertCmd.Flags().Lookup("cert-hosts"))

	genCertCmd.Flags().String("cert-org", defaultCertGenCertOrg, "Org for certificate")
	viper.SetDefault("cert.org", defaultCertGenCertOrg)
	_ = viper.BindPFlag("cert.org", genCertCmd.Flags().Lookup("cert-org"))

	genCertCmd.Flags().String("cert-name", defaultCertGenCertName, "Name for certificate")
	viper.SetDefault("cert.name", defaultCertGenCertName)
	_ = viper.BindPFlag("cert.name", genCertCmd.Flags().Lookup("cert-name"))

	genCertCmd.Flags().String("ecdsa-curve", defaultCertGenEcdsaCurve, "ECDSA curve to use")
	viper.SetDefault("ecdsa.curve", defaultCertGenEcdsaCurve)
	_ = viper.BindPFlag("ecdsa.curve", genCertCmd.Flags().Lookup("ecdsa-curve"))

	genCertCmd.Flags().Int("rsa-bits", defaultCertGenRSABits, "Size of RSA to gen")
	viper.SetDefault("rsa.bits", defaultCertGenRSABits)
	_ = viper.BindPFlag("rsa.bits", genCertCmd.Flags().Lookup("rsa-bits"))

	genCertCmd.Flags().Duration("valid-for", defaultCertGenValidFor, "Duration for the cert")
	viper.SetDefault("valid.for", defaultCertGenValidFor)
	_ = viper.BindPFlag("valid.for", genCertCmd.Flags().Lookup("valid-for"))

	genCertCmd.Flags().String("output-cert", defaultCertGenOutputCert, "Path to SSL cert")
	viper.SetDefault("output.cert", defaultCertGenOutputCert)
	_ = viper.BindPFlag("output.cert", genCertCmd.Flags().Lookup("output-cert"))

	genCertCmd.Flags().String("output-key", defaultCertGenOutputKey, "Path to SSL key")
	viper.SetDefault("output.key", defaultCertGenOutputKey)
	_ = viper.BindPFlag("output.key", genCertCmd.Flags().Lookup("output-key"))
}

func genCertAction(_ *cobra.Command, _ []string) {
	priv, err := parseEcdsaCurve()

	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to gen private key")

		os.Exit(1)
	}

	notBefore := time.Now()
	notAfter := notBefore.Add(viper.GetDuration("valid.for"))

	serialNumber, err := buildSerialNumber()

	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to gen serial number")

		os.Exit(1)
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{viper.GetString("cert.org")},
			CommonName:   viper.GetString("cert.name"),
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	for _, host := range viper.GetStringSlice("cert.host") {
		if ip := net.ParseIP(host); ip != nil {
			template.IPAddresses = append(
				template.IPAddresses,
				ip,
			)
		} else {
			template.DNSNames = append(
				template.DNSNames,
				host,
			)
		}
	}

	der, err := x509.CreateCertificate(
		rand.Reader,
		&template,
		&template,
		extractPublicKey(priv),
		priv,
	)

	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to create certificate")

		os.Exit(1)
	}

	crt, err := os.OpenFile(
		viper.GetString("output.cert"),
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		0o600,
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("cert", viper.GetString("output.cert")).
			Msg("Failed to open cert file")

		os.Exit(1)
	}

	if err := pem.Encode(
		crt,
		publicEncodeBlock(der),
	); err != nil {
		log.Error().
			Err(err).
			Msg("Failed to encode cert")

		os.Exit(1)
	}

	if err := crt.Close(); err != nil {
		log.Error().
			Err(err).
			Str("cert", viper.GetString("output.cert")).
			Msg("Failed to close cert file")

		os.Exit(1)
	}

	key, err := os.OpenFile(
		viper.GetString("output.key"),
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		0o600,
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("key", viper.GetString("output.key")).
			Msg("Failed to open key file")

		os.Exit(1)
	}

	if err := pem.Encode(
		key,
		privateEncodeBlock(priv),
	); err != nil {
		log.Error().
			Err(err).
			Msg("Failed to encode key")

		os.Exit(1)
	}

	if err := key.Close(); err != nil {
		log.Error().
			Err(err).
			Str("key", viper.GetString("output.key")).
			Msg("Failed to close key file")

		os.Exit(1)
	}

	log.Info().
		Str("cert", viper.GetString("output.cert")).
		Str("key", viper.GetString("output.key")).
		Msg("Successfully generated")
}

func parseEcdsaCurve() (interface{}, error) {
	switch viper.GetString("ecdsa.curve") {
	case "":
		return rsa.GenerateKey(rand.Reader, viper.GetInt("rsa.bits"))
	case "P224":
		return ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	case "P256":
		return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case "P384":
		return ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case "P521":
		return ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	default:
		return nil, fmt.Errorf("unrecognized elliptic curve: %q", viper.GetString("ecdsa.curve"))
	}
}

func buildSerialNumber() (*big.Int, error) {
	return rand.Int(
		rand.Reader,
		new(
			big.Int,
		).Lsh(
			big.NewInt(1),
			128,
		),
	)
}

func publicEncodeBlock(der []byte) *pem.Block {
	return &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: der,
	}
}

func privateEncodeBlock(priv interface{}) *pem.Block {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(k),
		}
	case *ecdsa.PrivateKey:
		b, err := x509.MarshalECPrivateKey(k)

		if err != nil {
			log.Error().
				Err(err).
				Msg("unable to marshal ECDSA key")
		}

		return &pem.Block{
			Type:  "EC PRIVATE KEY",
			Bytes: b,
		}
	default:
		return nil
	}
}

func extractPublicKey(priv interface{}) interface{} {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &k.PublicKey
	case *ecdsa.PrivateKey:
		return &k.PublicKey
	default:
		return nil
	}
}
