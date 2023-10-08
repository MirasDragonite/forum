package handlers

import (
	"net/http"

	"forum/internal/service"
)

type Handler struct {
	Service *service.Service
	Mux     *http.ServeMux
	Cache   map[string]int64
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{Service: service, Mux: http.NewServeMux(), Cache: make(map[string]int64)}
}

func (h *Handler) Router() {
	h.Mux.HandleFunc("/", h.home)
	h.Mux.HandleFunc("/signin", h.signin)
	h.Mux.HandleFunc("/register", h.signup)
	h.Mux.HandleFunc("/profile", h.profile)
	h.Mux.HandleFunc("/logout", h.logOut)
	h.Mux.HandleFunc("/submit-post", h.PostPageCreate)
	h.Mux.HandleFunc("/post/", h.PostPage)
}
