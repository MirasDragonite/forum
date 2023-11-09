package handlers

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) editComment(w http.ResponseWriter, r *http.Request) {
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
	intID, err := strconv.Atoi(id)
	if err != nil {
		h.logError(w, r, err, http.StatusBadRequest)
		return
	}
	comment, err := h.Service.CommentRedact.GetCommentByID(int64(intID))
	if err != nil {
		h.logError(w, r, err, http.StatusBadRequest)
		return
	}

	if user.Id != comment.CommentAuthorID {
		h.logError(w, r, errors.New("Not your comment"), http.StatusBadRequest)
		return
	}

	result := map[string]interface{}{
		"Comment":      comment,
		"EmptyContent": false,
		"User":         user,
	}

	tmp, err := template.ParseFiles("./ui/templates/edit_comment_page.html")
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

		comment.Content = strings.TrimSpace(r.Form.Get("commentContent"))

		if len(comment.Content) < 15 || len(comment.Content) > 150 {
			w.WriteHeader(http.StatusBadRequest)
			result["EmptyContent"] = true
			tmp.Execute(w, result)
			return
		}

		err = h.Service.CommentRedact.UpdateComment(comment)
		if err != nil {
			h.logError(w, r, err, http.StatusBadRequest)
			return
		}
		link := fmt.Sprintf("/post/%v", comment.PostID)
		http.Redirect(w, r, link, http.StatusSeeOther)
		return
	} else {
		h.logError(w, r, errors.New("Wrong Method"), http.StatusMethodNotAllowed)
		return
	}
}
