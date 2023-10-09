package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}
	fmt.Printf("Server running on http://localhost:%s\n", port)
	return s.httpServer.ListenAndServe()
}
