package api

import (
	"fmt"
	"net/http"
)

// GET /sample
func (s *server) sampleAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fmt.Fprintf(w, `{"hello": "world"}`)
}
