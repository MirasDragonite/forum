package handlers

import (
	"net/http"
	"time"
)

func (h *Handler) authorized(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Token")
		if err != nil {
			h.errorLog(err.Error())
			http.Redirect(w, r, "/register", http.StatusSeeOther)
			return

		}

		if !cookie.Expires.Before(time.Now()) {
			h.infoLog("Token time expired")
			err := h.Service.Authorization.DeleteToken(cookie)
			http.Redirect(w, r, "/register", http.StatusSeeOther)
			if err != nil {
				h.errorLog(err.Error())
				return
			}
			return
		}
		h.infoLog("Token is available")

		next.ServeHTTP(w, r)
	})
}
