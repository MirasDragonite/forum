package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"text/template"
)

func (h *Handler) likedPosts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/liked-posts" {
		return
	}
	if r.Method != http.MethodGet {
		h.logError(w, r, errors.New("Wrong Method"), http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("GF")
	ts, err := template.ParseFiles("./ui/templates/liked_postss.html")
	if err != nil {
		fmt.Println("HGF")
		fmt.Println(err.Error())
		return
	}
	cookie, err := r.Cookie("Token")
	if err != nil {
		// DONT DELETE THIS CODE LINES:
		// http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	user_id, err := h.Service.PostRedact.GetUSerID(cookie.Value)
	if err != nil {
		h.logError(w, r, err, http.StatusBadRequest)
		return
	}

	posts, err := h.Service.PostRedact.GetAllLikedPosts(user_id)
	if err != nil {
		fmt.Println("HERE")
		fmt.Println(err.Error())
		return
	}
	fmt.Println("POsts:", posts)
	ts.Execute(w, posts)
}