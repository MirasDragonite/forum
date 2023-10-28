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
			h.errorLog("Don't have any Token")
			http.Redirect(w, r, "/register", http.StatusSeeOther)
			return

		}
		_, err = h.Service.Authorization.GetUserByToken(cookie.Value)
		if err != nil {
			fmt.Println(err.Error())
			cookie.Name = "Token"
			cookie.Value = ""
			cookie.Path = "/"
			cookie.MaxAge = -1
			cookie.HttpOnly = false
			http.SetCookie(w, cookie)
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

func (h *Handler) isNotauthorized(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println()
		cookie, err := r.Cookie("Token")
		if err != nil {

			h.errorLog(err.Error())
			if err == http.ErrNoCookie {
				next.ServeHTTP(w, r)
				return

			} else {
				http.Redirect(w, r, "/register", http.StatusSeeOther)
				return
			}

		}

		user, err := h.Service.Authorization.GetUserByToken(cookie.Value)
		fmt.Println("USER IN FKCING :", user)
		if err != nil {
			fmt.Println(err.Error())
			cookie.Name = "Token"
			cookie.Value = ""
			cookie.Path = "/"
			cookie.MaxAge = -1
			cookie.HttpOnly = false
			http.SetCookie(w, cookie)
			next.ServeHTTP(w, r)
			return
		}

		if cookie.Expires.Before(time.Now()) {
			h.infoLog("Token time not expired")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		h.infoLog("Token is available")

		next.ServeHTTP(w, r)
	})
}

func (h *Handler) tokenAvilableChecker(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Token")
		if err != nil {
			fmt.Println("He's her")
			h.errorLog(err.Error())
			if err == http.ErrNoCookie {
				next.ServeHTTP(w, r)
				return

			} else {
				http.Redirect(w, r, "/register", http.StatusSeeOther)
				return
			}

		}
		_, err = h.Service.Authorization.GetUserByToken(cookie.Value)
		if err != nil {
			fmt.Println("Deleted token from session")
			fmt.Println(err.Error())
			cookie.Name = "Token"
			cookie.Value = ""
			cookie.Path = "/"
			cookie.MaxAge = -1
			cookie.HttpOnly = false
			fmt.Println("From middleware:", cookie.Value)
			http.SetCookie(w, cookie)

		}

		next.ServeHTTP(w, r)
	})
}
