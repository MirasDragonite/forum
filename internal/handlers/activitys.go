package handlers

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
)

func (h *Handler) activities(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/activities" {
		h.errorHandler(w, r, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		h.logError(w, r, errors.New("Wrong Method"), http.StatusMethodNotAllowed)
		return
	}

	ts, err := template.ParseFiles("./ui/templates/activity_page.html")
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
	likedPosts, err := h.Service.PostRedact.GetAllLikedPosts(user_id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	dislikedPosts, err := h.Service.PostRedact.GetAllDislikedPosts(user_id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	createdPosts, err := h.Service.PostRedact.GetAllUserPosts(user_id)
	userComments, err := h.Service.CommentRedact.GetAllUserComments(user_id)
	logged := false

	if user != nil {
		logged = true
	}

	result := map[string]interface{}{
		"User":          user,
		"Logged":        logged,
		"LikedPosts":    likedPosts,
		"DislikedPosts": dislikedPosts,
		"CreatedPosts":  createdPosts,
		"Comments":      userComments,
	}
	ts.Execute(w, result)
}
