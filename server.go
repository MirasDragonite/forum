package server

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler, tlsConfig *tls.Config) error {
	s.httpServer = &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		TLSConfig:    tlsConfig,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  5 * time.Second,
	}

	fmt.Println(s.httpServer.Addr)
	fmt.Printf("Server running on https://127.0.0.1:%s\n", port)

	return s.httpServer.ListenAndServeTLS("tls/cert.pem", "tls/key.pem")
}
