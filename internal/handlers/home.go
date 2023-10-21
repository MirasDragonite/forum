package handlers

import (
	"errors"
	"fmt"
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
	fmt.Println("POSTS:", posts)
	result := map[string]interface{}{
		"Posts":  posts,
		"User":   user,
		"Logged": logged,
	}
	fmt.Println(result)
	ts.Execute(w, result)
}
