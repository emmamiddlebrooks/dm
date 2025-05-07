package security

import (
	"crypto/tls"
	"testing"
)

func TestLoadCertificates(t *testing.T) {
	t.Run("invalid path", func(t *testing.T) {
		domains := map[string]string{
			"example.com": "/invalid/path",
		}
		_, err := LoadCertificates(domains)
		if err == nil {
			t.Fatal("expected error when loading from invalid path, got nil")
		}
	})
}

func TestCreateTLSConfig(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		cert := tls.Certificate{}
		certMap := map[string]tls.Certificate{
			"example.com": cert,
		}
		tlsConfig := CreateTLSConfig(certMap)
		if tlsConfig.GetCertificate == nil {
			t.Fatal("expected GetCertificate to be defined")
		}

		hello := &tls.ClientHelloInfo{ServerName: "example.com"}
		result, err := tlsConfig.GetCertificate(hello)
		if err != nil || result == nil {
			t.Fatalf("expected certificate to be returned, got: %v, err: %v", result, err)
		}
	})
}
