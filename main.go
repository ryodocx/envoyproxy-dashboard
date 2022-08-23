package main

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"

	"github.com/ryodocx/envoyproxy-dashboard/backend/api"
	_ "modernc.org/sqlite"
)

var (
	listenAddr = "0.0.0.0:8080"

	// TODO: change dir
	//go:embed .tmp/dist/*
	assets embed.FS
)

func main() {
	if e := os.Getenv("ENVOY_DASHBOARD_LISTENADDR"); e != "" {
		// TODO: validation
		listenAddr = e
	}

	// setup DB
	db, err := sql.Open("sqlite3", ":memory:?_fk=1")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// setup static resource
	a, err := fs.Sub(assets, ".tmp/dist") // TODO: change dir
	if err != nil {
		panic(err)
	}

	// setup server instance
	s, err := api.New(
		api.Config{
			DB:     db,
			Assets: a,
		},
	)
	if err != nil {
		panic(err)
	}
	defer s.Close()

	fmt.Printf("Start server: %s\n", listenAddr)
	// TODO: gracefull shutdown
	if err := http.ListenAndServe(listenAddr, s); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
