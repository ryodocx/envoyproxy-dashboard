package api

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"net/http"

	"envoyproxy-dashboard/backend/db"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	DB     *sql.DB
	Assets fs.FS
}

type server struct {
	db  *db.Client
	mux *http.ServeMux
}

func NewServer(c Config) (*server, error) {
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
	s.mux.HandleFunc("/sample", s.sampleAPI)

	// setup static resource
	s.mux.Handle("/", http.FileServer(http.FS(c.Assets)))

	// migration
	if err := s.db.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("migration failed: %s", err.Error())
	}

	return s, nil
}

// implementation of http.Handler interface
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
