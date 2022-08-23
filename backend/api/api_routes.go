package api

import (
	"fmt"
	"net/http"
)

// GET /routes
func (s *server) routes(w http.ResponseWriter, r *http.Request) {
	// TODO
	w.Header().Set("content-type", "application/json")
	fmt.Fprintf(w, `{"key": "value"}`)
}
