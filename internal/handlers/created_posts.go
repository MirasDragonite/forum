package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

func (h *Handler) createdPosts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/created-posts" {
		return
	}

	ts, err := template.ParseFiles("./ui/templates/created_posts.html")
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
	fmt.Println("USERID:", user_id)
	posts, err := h.Service.PostRedact.GetAllUserPosts(user_id)
	if err != nil {
		fmt.Println("HERE")
		fmt.Println(err.Error())
		return
	}
	fmt.Println("POsts:", posts)
	ts.Execute(w, posts)
}
