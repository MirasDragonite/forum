package main

import (
	"flag"
	"log"

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
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handlers.NewHandler(service)
	handler.Router()
	srv := new(server.Server)
	if err := srv.Run(port, handler.Mux); err != nil {
		log.Fatal(err)
	}
}
