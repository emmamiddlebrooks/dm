package security

import (
	"crypto/tls"
	"fmt"
	"path/filepath"
	"strings"
)

func LoadCertificates(domains map[string]string) (map[string]tls.Certificate, error) {
	certMap := map[string]tls.Certificate{}
	for domain, path := range domains {
		cert, err := tls.LoadX509KeyPair(filepath.Join(path, "fullchain.pem"), filepath.Join(path, "privkey.pem"))
		if err != nil {
			return nil, fmt.Errorf("failed to load cert for %s: %w", domain, err)
		}
		certMap[domain] = cert
	}
	return certMap, nil
}

func CreateTLSConfig(certMap map[string]tls.Certificate) *tls.Config {
	return &tls.Config{
		GetCertificate: func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
			name := strings.ToLower(hello.ServerName)
			if cert, ok := certMap[name]; ok {
				return &cert, nil
			}
			return nil, nil
		},
	}
}
