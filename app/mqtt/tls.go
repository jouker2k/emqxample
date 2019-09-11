package mqtt

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

// NewTLSConfig returns default TLS config
func NewTLSConfig(capath string) *tls.Config {
	certpool := x509.NewCertPool()
	pemCerts, err := ioutil.ReadFile(capath)

	if err != nil {
		log.Fatalf("Cannot read root ca file: %s", capath)
	}

	certpool.AppendCertsFromPEM(pemCerts)

	return &tls.Config{
		// RootCAs = certs used to verify server cert.
		RootCAs: certpool,
		// ClientAuth = whether to request cert from server.
		// Since the server is set up for SSL, this happens
		// anyways.
		ClientAuth: tls.NoClientCert,
		// ClientCAs = certs used to validate client cert.
		ClientCAs: nil,
		// InsecureSkipVerify = verify that cert contents
		// match server. IP matches what is in cert etc.
		InsecureSkipVerify: true,
		// Certificates = list of certs client sends to server.
		Certificates: []tls.Certificate{},
	}
}
