package api

import (
	"context"
	"database/sql"
	"io/fs"
	"net/http"

	"envoyproxy-dashboard/backend/db"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

type Config struct {
	DB     *sql.DB
	Assets fs.FS
}

type server struct {
	db  *db.Client
	mux *http.ServeMux
}

func NewServer(c Config) (http.Handler, error) {

	// check DB connection
	if err := c.DB.Ping(); err != nil {
		return nil, errors.Wrap(err, "error at ping to DB")
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
	if err := s.migration(); err != nil {
		return nil, errors.Wrap(err, "migration failed")
	}

	return s, nil
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *server) Debug() {

	s.db.Route.Create().
		SetDomain("domain").
		SetPath("/")
}

func (s *server) migration() error {
	if err := s.db.Schema.Create(context.Background()); err != nil {
		return errors.Wrap(err, "failed creating schema resources")
	}
	return nil
}
