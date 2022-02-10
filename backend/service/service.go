package service

import (
	"context"
	"database/sql"
	"net/http"

	"envoyproxy-dashboard/backend/db"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

type Config struct {
	DB *sql.DB
}

type Service struct {
	db           *db.Client
	HttpServeMux *http.ServeMux
}

func New(c Config) (*Service, error) {

	// check DB connection
	if err := c.DB.Ping(); err != nil {
		return nil, errors.Wrap(err, "error at ping to DB")
	}

	// service instrance
	s := &Service{
		db:           db.NewClient(db.Driver(entsql.OpenDB("sqlite3", c.DB))),
		HttpServeMux: &http.ServeMux{},
	}

	// setup endpoints
	s.HttpServeMux.HandleFunc("/sample", s.sampleAPI)

	// migration
	if err := s.migration(); err != nil {
		return nil, errors.Wrap(err, "migration failed")
	}

	return s, nil
}

func (s *Service) Close() {
	s.db.Close()
}

func (s *Service) Debug() {

	s.db.Route.Create().
		SetDomain("domain").
		SetPath("/")
}

func (s *Service) migration() error {
	if err := s.db.Schema.Create(context.Background()); err != nil {
		return errors.Wrap(err, "failed creating schema resources")
	}
	return nil
}
