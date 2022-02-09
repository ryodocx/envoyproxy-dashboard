package main

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"

	"envoyproxy-dashboard/pkg/service"

	_ "github.com/mattn/go-sqlite3"
)

// TODO: change dir
//go:embed .tmp/dist/*
var assets embed.FS

var listenAddr = "0.0.0.0:8080"

func init() {
	if e := os.Getenv("ENVOY_DASHBOARD_LISTENADDR"); e != "" {
		// TODO: validation
		listenAddr = e
	}
}

func main() {
	// setup DB
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	// setup service instance
	c := service.Config{
		DB: db,
	}
	s, err := service.New(c)
	if err != nil {
		panic(err)
	}

	// setup static resource
	webRoot, err := fs.Sub(assets, ".tmp/dist") // TODO: change dir
	if err != nil {
		panic(err)
	}
	s.HttpServeMux.Handle("/", http.FileServer(http.FS(webRoot)))

	fmt.Printf("Start server: %s\n", listenAddr)
	// TODO: gracefull shutdown
	if err := http.ListenAndServe(listenAddr, s.HttpServeMux); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
