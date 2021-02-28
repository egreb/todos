package routes

import (
	"net/http"
)

// Server http
type Server struct {
	Mux *http.ServeMux
}

// NewServer http server
func NewServer() (*Server, error) {
	srv := &Server{
		Mux: http.NewServeMux(),
	}

	srv.routes()

	return srv, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Mux.ServeHTTP(w, r)
}
