package api

import (
	"fmt"
	"net/http"
)

// /sample
func (s *server) sampleAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fmt.Fprintf(w, `{"key": "value"}`)
}