package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
)

// TODO: change dir
//go:embed .tmp/dist/*
var assets embed.FS

var listenAddr = "0.0.0.0:8080"

func init() {
	if e := os.Getenv("ENVOY_DASHBOARD_LISTENADDR"); e != "" {
		listenAddr = e
	}
}

func main() {
	fmt.Printf("Start server: %s\n", listenAddr)

	webRoot, err := fs.Sub(assets, ".tmp/dist") // TODO: change dir
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/envoydatas", envoyDataEndpoint)
	http.Handle("/", http.FileServer(http.FS(webRoot)))

	// TODO: gracefull shutdown
	if err := http.ListenAndServe(listenAddr, http.DefaultServeMux); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func envoyDataEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fmt.Fprintf(w, `{"key": "value"}`)
}
