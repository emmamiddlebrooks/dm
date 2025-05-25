package server

import (
	"crypto/tls"
	"gfi/internal"
	"gfi/security"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func StartHTTPRedirectServer() {
	go func() {
		if err := http.ListenAndServe(":80", http.HandlerFunc(internal.Redirect)); err != nil {
			log.Fatalf("HTTP redirect server failed: %v", err)
		}
	}()
}

func StartHTTPSServer(certMap map[string]tls.Certificate) {
	tlsConfig := security.CreateTLSConfig(certMap)
	server := &http.Server{
		Addr:      ":443",
		TLSConfig: tlsConfig,
		Handler:   nil,
	}
	log.Fatal(server.ListenAndServeTLS("", ""))
}

func DefaultFileServer(rw http.ResponseWriter, req *http.Request) {
	host := req.Host
	subdir := internal.GetSubdir(host)
	root := http.Dir("./static/" + subdir)

	path := filepath.Join("static", subdir, req.URL.Path)
	fi, err := os.Stat(path)

	if err != nil || fi.IsDir() {
		if req.URL.Path == "/" {
			http.ServeFile(rw, req, filepath.Join("static", subdir, "index.html"))
		} else {
			http.Redirect(rw, req, "/", http.StatusFound)
		}
		return
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setNoCacheHeaders(w)
		http.FileServer(root).ServeHTTP(w, r)
	})
	handler.ServeHTTP(rw, req)
}

func setNoCacheHeaders(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Surrogate-Control", "no-store")
}
