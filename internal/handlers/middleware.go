package handlers

import (
	"fmt"
	"log"
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
		// user, err := h.Service.Authorization.GetUserByToken(cookie.Value)
		// if user.Id <= 0 {
		// 	http.Redirect(w, r, "/logout", http.StatusSeeOther)
		// 	return
		// }
		h.infoLog("Token is available")

		next.ServeHTTP(w, r)
	})
}

func (h *Handler) isNotauthorized(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println()
		cookie, err := r.Cookie("Token")
		if err != nil {
			fmt.Println("Here3")
			h.errorLog(err.Error())
			if err == http.ErrNoCookie {
				fmt.Println("NoCookie")
				next.ServeHTTP(w, r)
				fmt.Println("After serve")
				return

			} else {
				fmt.Println("Here4")
				http.Redirect(w, r, "/register", http.StatusSeeOther)
				return
			}

		}

		if cookie.Expires.Before(time.Now()) {
			h.infoLog("Token time not expired")
			// err := h.Service.Authorization.DeleteToken(cookie)
			http.Redirect(w, r, "/profile", http.StatusSeeOther)
			return
		}
		h.infoLog("Token is available")

		next.ServeHTTP(w, r)
	})
}
