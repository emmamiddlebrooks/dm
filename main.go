//go:build !test

package main

import (
	"context"
	"gfi/guest"
	"gfi/security"
	"gfi/server"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	http.HandleFunc("/", server.DefaultFileServer)
	http.Handle("/submit", guest.NewGuestHandler(context.Background(), logger))

	// HTTP
	server.StartHTTPRedirectServer()

	// HTTPS
	certs, err := security.LoadCertificates(map[string]string{
		"dynamicmultimediaga.com":          "/etc/letsencrypt/live/dynamicmultimediaga.com",
		"wildleap.dynamicmultimediaga.com": "/etc/letsencrypt/live/wildleap.dynamicmultimediaga.com",
	})
	if err != nil {
		logger.Error("Certificate loading failed", slog.String("error", err.Error()))
		return
	}
	server.StartHTTPSServer(certs)
}
