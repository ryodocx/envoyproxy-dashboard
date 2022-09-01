package api

import (
	"io/fs"
	"net/http"

	"github.com/ryodocx/envoyproxy-dashboard/backend/client/envoy"
)

// interface guard
var (
	_ http.Handler = (*server)(nil)
)

type Config struct {
	// DB     *sql.DB
	Client *envoy.Client
	Assets fs.FS
}

type server struct {
	// db  *db.Client
	mux  *http.ServeMux
	conf *Config
}

func New(c Config) (http.Handler, error) {

	s := &server{
		// db:  db.NewClient(db.Driver(entsql.OpenDB("sqlite3", c.DB))),
		mux:  &http.ServeMux{},
		conf: &c,
	}

	// setup endpoints
	endpoints := []struct {
		path    string
		handler func(w http.ResponseWriter, r *http.Request)
		config  *middlewareConfig
	}{
		{
			path:    "/", // TODO: remove
			handler: s.routes,
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
	// if err := s.db.Schema.Create(context.Background()); err != nil {
	// 	return nil, fmt.Errorf("migration failed: %s", err.Error())
	// }

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
	// defer s.db.Close()
}
