package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) deleteComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.logError(w, r, errors.New("Wrong Method"), http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()

	cookie, err := r.Cookie("Token")
	if err != nil {
		h.logError(w, r, err, http.StatusUnauthorized)
		return
	}
	comment_id_string := r.URL.Path[16:]
	fmt.Println(comment_id_string)
	comment_id, err := strconv.Atoi(comment_id_string)
	if err != nil {
		h.logError(w, r, err, http.StatusBadRequest)
		return
	}
	user, err := h.Service.Authorization.GetUserByToken(cookie.Value)
	if err != nil {
		h.logError(w, r, errors.New("User not authorized"), http.StatusUnauthorized)
		return
	}
	comment, err := h.Service.CommentRedact.GetCommentByID(int64(comment_id))
	if err != nil {
		h.logError(w, r, err, http.StatusBadRequest)
		return

	}
	postId := comment.PostID
	if user.Id == comment.CommentAuthorID {
		err = h.Service.CommentRedact.DeleteComment(&comment)
		if err != nil {
			h.logError(w, r, err, http.StatusBadRequest)
			return
		}
	}
	link := fmt.Sprintf("/post/%v", postId)
	http.Redirect(w, r, link, http.StatusSeeOther)
	return
}
