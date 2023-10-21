package handlers

import (
	"errors"
	"net/http"
	"text/template"

	"forum/structs"
)

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.logError(w, r, errors.New("Wrong Method"), http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		h.errorHandler(w, r, http.StatusNotFound)
		return
	}

	ts, err := template.ParseFiles("./ui/templates/home_page.html")
	if err != nil {
		h.logError(w, r, err, http.StatusInternalServerError)
		return
	}
	var user *structs.User
	cookie, err := r.Cookie("Token")
	if err != nil {
	} else {
		user, err = h.Service.Authorization.GetUserByToken(cookie.Value)
		if err != nil {
			user = nil
		}
	}

	logged := false

	if user != nil {
		logged = true
	}

	posts, err := h.Service.PostRedact.GetAllPosts()

	result := map[string]interface{}{
		"Posts":  posts,
		"User":   user,
		"Logged": logged,
	}

	ts.Execute(w, result)
}
