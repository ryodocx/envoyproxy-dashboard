package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"net/url"
	"os"

	"github.com/ryodocx/envoyproxy-dashboard/backend/api"
	"github.com/ryodocx/envoyproxy-dashboard/backend/client/envoy"
)

var (
	listenAddr = "127.0.0.1:8080"

	// TODO: change dir
	//go:embed .tmp/dist/*
	assets embed.FS
)

func main() {
	if e := os.Getenv("ENVOY_DASHBOARD_LISTENADDR"); e != "" {
		// TODO: validation
		listenAddr = e
	}

	// setup static resource
	a, err := fs.Sub(assets, ".tmp/dist") // TODO: change dir
	if err != nil {
		panic(err)
	}

	// setup envoy client
	u, err := url.Parse(os.Getenv("ENVOY_ADDR"))
	if err != nil {
		panic(err)
	}
	c, err := envoy.New(envoy.Config{
		EnvoyURL: u,
	})

	// setup server instance
	s, err := api.New(
		api.Config{
			Assets: a,
			Client: c,
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Start server: %s\n", listenAddr)
	// TODO: gracefull shutdown
	if err := http.ListenAndServe(listenAddr, s); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
