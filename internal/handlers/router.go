package handlers

import (
	"forum/internal/service"
	"net/http"
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
	h.Mux.Handle("/signin", h.isNotauthorized(h.signin))
	h.Mux.Handle("/register", h.isNotauthorized(h.signup))
	h.Mux.Handle("/logout", h.authorized(h.logOut))
	h.Mux.Handle("/submit-post", h.authorized(h.PostPageCreate))
	h.Mux.Handle("/like-post/", h.authorized(h.likePost))
	h.Mux.Handle("/like-comment/", h.authorized(h.likeComment))
	h.Mux.Handle("/post/", h.tokenAvilableChecker(h.PostPage))
	h.Mux.Handle("/liked-posts", h.authorized(h.likedPosts))
	h.Mux.Handle("/created-posts", h.authorized(h.createdPosts))
	h.Mux.Handle("/notify", h.authorized(h.notify))
	h.Mux.Handle("/activities", h.authorized(h.activities))
	h.Mux.Handle("/delete-post/", h.authorized(h.deletePost))
	h.Mux.Handle("/delete-comment/", h.authorized(h.deleteComment))
	h.Mux.HandleFunc("/github/auth", h.githubLogin)
	h.Mux.HandleFunc("/github/callback", h.githubLoginCallBack)
	h.Mux.HandleFunc("/google/auth", h.googleLogin)
	h.Mux.HandleFunc("/google/callback", h.googleLoginCallBack)
}
