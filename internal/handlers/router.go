package handlers

import (
	"net/http"

	"forum/internal/service"
	"forum/structs"
)

type Handler struct {
	Service *service.Service
	Mux     *http.ServeMux
	Cache   map[string]structs.User
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{Service: service, Mux: http.NewServeMux()}
}

func (h *Handler) Router() {
	h.Mux.HandleFunc("/", h.home)
	h.Mux.HandleFunc("/auth", h.signin)
	h.Mux.HandleFunc("/register", h.signup)
}
