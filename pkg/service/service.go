package service

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
)

type Config struct {
	DB *sql.DB
}

type Service struct {
	db           *sql.DB
	HttpServeMux *http.ServeMux
}

func New(c Config) (*Service, error) {

	// check DB connection
	if err := c.DB.Ping(); err != nil {
		return nil, errors.Wrap(err, "error at ping to DB")
	}

	// setup endpoints
	serveMux := &http.ServeMux{}
	serveMux.HandleFunc("/sample", sampleAPI)

	return &Service{
		db:           c.DB,
		HttpServeMux: serveMux,
	}, nil
}
