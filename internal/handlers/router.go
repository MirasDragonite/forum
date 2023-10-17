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
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	h.Mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	h.Mux.HandleFunc("/", h.home)
	// h.Mux.Handle("/signin", h.isNotauthorized(h.signin))
	h.Mux.HandleFunc("/signin", h.signin)
	h.Mux.Handle("/register", h.isNotauthorized(h.signup))
	h.Mux.Handle("/profile", h.authorized(h.profile))
	h.Mux.Handle("/logout", h.authorized(h.logOut))
	h.Mux.Handle("/submit-post", h.authorized(h.PostPageCreate))
	h.Mux.Handle("/like-post/", h.authorized(h.likePost))
	h.Mux.Handle("/like-comment/", h.authorized(h.likeComment))
	h.Mux.HandleFunc("/post/", h.PostPage)
}
