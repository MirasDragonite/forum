package handlers

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func (h *Handler) editPost(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	cookie, err := r.Cookie("Token")
	if err != nil {
		h.logError(w, r, err, http.StatusUnauthorized)
		return
	}

	user, err := h.Service.Authorization.GetUserByToken(cookie.Value)
	if err != nil {
		h.logError(w, r, err, http.StatusBadRequest)
		return
	}

	post, err := h.Service.PostRedact.GetPostBy("id", id, user.Id)
	if err != nil {
		h.logError(w, r, err, http.StatusBadRequest)
		return
	}
	if user.Id != post.PostAuthorID {
		h.logError(w, r, errors.New("Not your post"), http.StatusBadRequest)
		return
	}

	result := map[string]interface{}{
		"Post":         post,
		"EmptyTitle":   false,
		"EmptyContent": false,
		"User":         user,
	}

	tmp, err := template.ParseFiles("./ui/templates/edit_post_page.html")
	if err != nil {
		h.logError(w, r, err, http.StatusInternalServerError)
		return
	}
	err = r.ParseForm()
	if err != nil {
		h.logError(w, r, err, http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodGet {
		tmp.Execute(w, result)
	} else if r.Method == http.MethodPost {
		post.Title = strings.TrimSpace(r.Form.Get("postTitle"))
		post.Content = strings.TrimSpace(r.Form.Get("postContent"))

		if len(post.Title) < 5 || len(post.Title) > 50 {
			w.WriteHeader(http.StatusBadRequest)
			result["EmptyTitle"] = true
			tmp.Execute(w, result)
			return
		}
		if len(post.Content) < 15 || len(post.Content) > 250 {
			w.WriteHeader(http.StatusBadRequest)
			result["EmptyContent"] = true
			tmp.Execute(w, result)
			return
		}
		fmt.Println("Content:", post.Content)
		fmt.Println("Title", post.Title)
		fmt.Println("ID:", post.Id)
		err = h.Service.PostRedact.RedactContentPost(post)
		if err != nil {
			h.logError(w, r, err, http.StatusBadRequest)
			return
		}
		link := fmt.Sprintf("/post/%v", id)
		http.Redirect(w, r, link, http.StatusSeeOther)
		return
	} else {
		h.logError(w, r, errors.New("Wrong Method"), http.StatusMethodNotAllowed)
		return
	}
}
