package main

import (
	"crypto/tls"
	"flag"
	"log"
	"time"

	server "forum"

	"forum/internal/handlers"
	"forum/internal/repository"
	"forum/internal/service"
)

var port = ""

func main() {
	flag.StringVar(&port, "port", "8000", "to set port")
	flag.Parse()
	db, err := repository.NewDB()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handlers.NewHandler(service)
	handler.Router()

	tlsConfig := &tls.Config{
		Time: time.Now,
		// GetCertificate: certManager.GetCertificate,
		CipherSuites: []uint16{tls.TLS_AES_128_GCM_SHA256, tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256},
	}

	srv := new(server.Server)
	if err := srv.Run(port, handler.Mux, tlsConfig); err != nil {
		log.Fatal(err)
	}
}
