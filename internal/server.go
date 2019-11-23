package apiserver

import (
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	Mux          *mux.Router
	Config       *Config
}

func NewServer(config *Config) (*Server, error) {
	s := &Server{
		Mux:          mux.NewRouter(),
		Config:       config,
	}
	return s, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Mux.ServeHTTP(w, r)
}



func (s *Server) ConfigureServer(db *sql.DB) {
//	TODO confugure server
}


