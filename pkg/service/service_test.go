package service_test

import (
	"database/sql"
	"envoyproxy-dashboard/pkg/service"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestNew(t *testing.T) {

	inMemDB, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal("can't open in-memory DB", err)
	}

	c := service.Config{
		DB: inMemDB,
	}
	s, err := service.New(c)
	if err != nil {
		t.Fatal("can't init new Service instance", err)
	}

	_ = s // TODO: check behavior
}
