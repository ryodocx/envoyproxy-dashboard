package api

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/ryodocx/envoyproxy-dashboard/backend/db"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/mattn/go-sqlite3"
)

// interface guard
var (
	_ http.Handler = (*server)(nil)
)

type Config struct {
	DB     *sql.DB
	Assets fs.FS
}

type server struct {
	db  *db.Client
	mux *http.ServeMux
}

func New(c Config) (http.Handler, error) {
	// check DB connection
	if err := c.DB.Ping(); err != nil {
		return nil, fmt.Errorf("error at ping to DB: %s", err.Error())
	}

	// service instrance
	s := &server{
		db:  db.NewClient(db.Driver(entsql.OpenDB("sqlite3", c.DB))),
		mux: &http.ServeMux{},
	}

	// setup endpoints
	endpoints := []struct {
		path    string
		handler func(w http.ResponseWriter, r *http.Request)
		config  *middlewareConfig
	}{
		{
			path:    "/sample",
			handler: s.sampleAPI,
			config:  nil,
		},
		{
			path:    "/routes",
			handler: s.routes,
			config:  nil,
		},
	}
	for _, v := range endpoints {
		s.mux.HandleFunc(middleware(v.path, v.handler, v.config))
	}

	// setup static resource
	s.mux.Handle("/", http.FileServer(http.FS(c.Assets)))

	// migration
	if err := s.db.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("migration failed: %s", err.Error())
	}

	return s, nil
}

type middlewareConfig struct{}

func middleware(path string, handler http.HandlerFunc, c *middlewareConfig) (string, http.HandlerFunc) {
	return path, func(w http.ResponseWriter, r *http.Request) {
		// TODO: 共通処理
		handler(w, r)
	}
}

// implementation of http.Handler interface
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
	defer s.db.Close()
}
