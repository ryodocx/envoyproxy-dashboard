package api_test

import (
	"database/sql"
	"fmt"
	"io/fs"
	"testing"

	"github.com/ryodocx/envoyproxy-dashboard/backend/api"

	_ "github.com/mattn/go-sqlite3"
)

type dummyFS struct{}

func (d *dummyFS) Open(name string) (fs.File, error) {
	return nil, fmt.Errorf("I am dummy file system")
}

func TestNewServer(t *testing.T) {

	inMemDB, err := sql.Open("sqlite3", ":memory:?_fk=1")
	if err != nil {
		t.Fatal("can't open in-memory DB", err)
	}

	c := api.Config{
		DB:     inMemDB,
		Assets: &dummyFS{},
	}
	s, err := api.New(c)
	if err != nil {
		t.Fatal("can't init new Server instance:", err)
	}

	_ = s // TODO: check behavior
	// https://budougumi0617.github.io/2020/05/29/go-testing-httptest/
}
