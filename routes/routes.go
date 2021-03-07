package routes

import (
	"net/http"
	"path/filepath"

	"egreb.net/todos/generated"
	"github.com/go-pg/pg/v10"
	"github.com/pacedotdev/oto/otohttp"
)

// Routes HTTP
func (s *Server) Routes(db *pg.DB) {

	g := GreeterService{}
	todoHandler := TodoService{DB: db}

	server := otohttp.NewServer()
	server.Basepath = "/api/"

	generated.RegisterGreeterService(server, g)
	generated.RegisterTodoService(server, todoHandler)

	s.Mux.Handle("/api/", server)
	s.Mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	s.Mux.HandleFunc("/", s.handleIndex())
}

func (s *Server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("public", "index.html"))
	}
}
