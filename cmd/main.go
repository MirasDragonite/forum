package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	server "forum"

	"forum/internal/handlers"
	"forum/internal/repository"
	"forum/internal/service"

	"golang.org/x/crypto/acme/autocert"
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

	certManager := autocert.Manager{
		Prompt: autocert.AcceptTOS,
		Cache:  autocert.DirCache("cach"),
		Email:  "kabykenov.miras@yandex.kz",
	}
	tlsConfig := &tls.Config{
		Time:           time.Now,
		GetCertificate: certManager.GetCertificate,
		CipherSuites:   []uint16{tls.TLS_AES_128_GCM_SHA256, tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256},
		Certificates:   []tls.Certificate{},
	}
	go func() {
		fmt.Println("here")
		if err := http.ListenAndServe(":8080", certManager.HTTPHandler(nil)); err != nil {
			fmt.Println(err)
			fmt.Print("go routine")
			return
		}
	}()
	srv := new(server.Server)
	if err := srv.Run(port, handler.Mux, tlsConfig); err != nil {
		log.Fatal(err)
	}
}
