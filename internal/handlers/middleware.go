package handlers

import (
	"net/http"
	"time"
)

func (h *Handler) authorized(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Token")
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		user, err := h.Service.GetUserByToken(cookie.Value)
		if user.ExpairedData < time.Now().Format("2006-01-02 15:04:05") {

			err := h.Service.Authorization.DeleteToken(cookie)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			if err != nil {
				return
			}

		}

		next.ServeHTTP(w, r)
	})
}
