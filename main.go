package main

import (
	"context"
	"crypto/tls"
	"gfi/guest"
	"gfi/internal"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func defaultFileServer(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	subdir := internal.GetSubdir(host)
	root := http.Dir("./static/" + subdir)

	path := filepath.Join("static", subdir, r.URL.Path)
	fi, err := os.Stat(path)

	if err != nil || fi.IsDir() {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, filepath.Join("static", subdir, "index.html"))
		} else {
			http.Redirect(w, r, "/", http.StatusFound)
		}
		return
	}

	http.FileServer(root).ServeHTTP(w, r)
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	// connect to db
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		logger.Error(err.Error())
		return
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	logger.Info("Successfully connected to MongoDB")

	http.HandleFunc("/", defaultFileServer)
	guestHandler := guest.NewGuestHandler(context.Background(), logger, client)
	http.Handle("/submit", guestHandler)

	// HTTP
	go http.ListenAndServe(":80", http.HandlerFunc(internal.Redirect))

	// HTTPS
	certMap := map[string]tls.Certificate{}
	load := func(domain, path string) {
		cert, err := tls.LoadX509KeyPair(path+"/fullchain.pem", path+"/privkey.pem")
		if err != nil {
			log.Fatalf("Failed to load cert for %s: %v", domain, err)
		}
		certMap[domain] = cert
	}
	load("dynamicmultimediaga.com", "/etc/letsencrypt/live/dynamicmultimediaga.com")
	load("wildleap.dynamicmultimediaga.com", "/etc/letsencrypt/live/wildleap.dynamicmultimediaga.com")
	tlsConfig := &tls.Config{
		GetCertificate: func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
			name := strings.ToLower(hello.ServerName)
			if cert, ok := certMap[name]; ok {
				return &cert, nil
			}
			return nil, nil
		},
	}
	server := &http.Server{
		Addr:      ":443",
		TLSConfig: tlsConfig,
		Handler:   nil,
	}
	log.Fatal(server.ListenAndServeTLS("", ""))
}
