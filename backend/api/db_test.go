package api

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

type dummyFS struct{}

func (d *dummyFS) Open(name string) (fs.File, error) {
	return nil, fmt.Errorf("I am dummy file system")
}

func TestDB(t *testing.T) {

	inMemDB, err := sql.Open("sqlite3", ":memory:?_fk=1")
	if err != nil {
		t.Fatal("can't open in-memory DB", err)
	}

	c := Config{
		DB:     inMemDB,
		Assets: &dummyFS{},
	}
	s, err := New(c)
	if err != nil {
		t.Fatal("can't init new Server instance:", err)
	}

	_, err = s.db.Route.Create().
		SetDomain("domain").
		SetPath("/").
		Save(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	routes, err := s.db.Route.Query().All(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range routes {
		t.Log(v.String())
	}
}
