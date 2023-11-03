package handlers

import (
	"errors"
	"fmt"
	"net/http"
)

func (h *Handler) deletePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.logError(w, r, errors.New("Wrong Method"), http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()

	cookie, err := r.Cookie("Token")
	if err != nil {
		h.logError(w, r, err, http.StatusNonAuthoritativeInfo)
		return
	}
	post_id_string := r.URL.Path[13:]
	fmt.Println(post_id_string)
	user, err := h.Service.Authorization.GetUserByToken(cookie.Value)

	post, err := h.Service.PostRedact.GetPostBy("id", post_id_string, user.Id)
	if err != nil {
		h.logError(w, r, err, http.StatusInternalServerError)
		return
	}
	if post.PostAuthorID != user.Id {
		h.logError(w, r, errors.New("You can't delete this post"), http.StatusBadRequest)
		return
	}

	err = h.Service.PostRedact.DeletePost(post)
	if err != nil {
		h.logError(w, r, err, http.StatusBadRequest)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}
