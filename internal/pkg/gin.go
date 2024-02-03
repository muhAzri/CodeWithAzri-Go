package pkg

import (
	"net/http"
)

type Server struct {
	Mux *http.ServeMux
}

func NewServer() *Server {
	s := &Server{
		Mux: http.NewServeMux(),
	}
	return s
}
