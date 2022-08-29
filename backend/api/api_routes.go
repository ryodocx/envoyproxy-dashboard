package api

import (
	"fmt"
	"net/http"
)

// GET /routes
func (s *server) routes(w http.ResponseWriter, r *http.Request) {
	// TODO
	routes, err := s.conf.Client.GetRouteConfigurations()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%v", err)

	}
	for i, v := range routes {
		fmt.Fprintf(w, "name[%d]: %s\n", i, v.Name)
	}
	// w.Header().Set("content-type", "application/json")
}
