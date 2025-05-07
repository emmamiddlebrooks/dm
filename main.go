//go:build !test

package main

import (
	"context"
	"gfi/db"
	"gfi/guest"
	"gfi/security"
	"gfi/server"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	// connect to db
	client, err := db.ConnectToMongo(context.Background(), "mongodb://localhost:27017")
	if err != nil {
		logger.Error("mongo connection failed", slog.String("error", err.Error()))
		return
	}
	logger.Info("Successfully connected to MongoDB")

	http.HandleFunc("/", server.DefaultFileServer)
	http.Handle("/submit", guest.NewGuestHandler(context.Background(), logger, client))

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
