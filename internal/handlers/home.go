package handlers

import (
	"fmt"
	"forum/structs"
	"html/template"
	"net/http"
)

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		h.errorHandler(w, r, http.StatusNotFound)
		return
	}

	ts, err := template.ParseFiles("./ui/templates/home_page.html")
	if err != nil {
		h.logError(w, r, err, http.StatusInternalServerError)
		return
	}
	err = r.ParseForm()

	var user *structs.User
	cookie, err := r.Cookie("Token")
	fmt.Println("cookie in home", cookie)
	if err != nil {
		fmt.Println("Cant find token")
	} else {
		fmt.Println("here ")
		user, err = h.Service.Authorization.GetUserByToken(cookie.Value)
		if err != nil {
			user = nil
		}

	}
	logged := false

	if user != nil {
		logged = true
	}
	var posts []structs.Post
	if r.Method == http.MethodPost {
		posts = h.filter(w, r)
	} else if r.Method == http.MethodGet {
		posts, err = h.Service.PostRedact.GetAllPosts()
		if err != nil {

			h.logError(w, r, err, http.StatusInternalServerError)
			return
		}
	} else {
		h.logError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	result := map[string]interface{}{
		"Posts":  posts,
		"User":   user,
		"Logged": logged,
	}

	ts.Execute(w, result)
}
