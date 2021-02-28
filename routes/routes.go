package routes

import (
	"net/http"
	"path/filepath"

	"egreb.net/todos/generated"
	"github.com/pacedotdev/oto/otohttp"
)

func (s *Server) routes() {
	g := GreeterService{}
	server := otohttp.NewServer()
	server.Basepath = "/api/"

	generated.RegisterGreeterService(server, g)

	s.Mux.Handle("/api/", server)
	s.Mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	s.Mux.HandleFunc("/", s.handleIndex())
}

func (s *Server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("public", "index.html"))
	}
}
