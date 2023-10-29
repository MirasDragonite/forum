package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"text/template"
)

func (h *Handler) createdPosts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/created-posts" {
		return
	}
	if r.Method != http.MethodGet {
		h.logError(w, r, errors.New("Wrong Method"), http.StatusMethodNotAllowed)
		return
	}
	ts, err := template.ParseFiles("./ui/templates/created_posts.html")
	if err != nil {
		h.logError(w, r, err, http.StatusInternalServerError)
		return
	}
	cookie, err := r.Cookie("Token")
	if err != nil {
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	user_id, err := h.Service.PostRedact.GetUSerID(cookie.Value)
	if err != nil {
		h.logError(w, r, err, http.StatusBadRequest)
		return
	}
	user, err := h.Service.Authorization.GetUserByToken(cookie.Value)
	if err != nil {
		h.logError(w, r, err, http.StatusBadRequest)
		return
	}
	posts, err := h.Service.PostRedact.GetAllUserPosts(user_id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	logged := false

	if user != nil {
		logged = true
	}

	result := map[string]interface{}{
		"User":   user,
		"Logged": logged,
		"Post":   posts,
	}
	ts.Execute(w, result)
}
