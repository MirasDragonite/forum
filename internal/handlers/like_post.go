package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) likePost(w http.ResponseWriter, r *http.Request) {
	var input struct {
		reaction int64 `json:"reaction"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		h.logError(w, r, err, http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("Token")
	if err != nil {
		h.logError(w, r, err, http.StatusNonAuthoritativeInfo)
		return
	}
	// post_id_string := r.URL.Path[6:]
	// post_id, err := strconv.Atoi(post_id_string)
	if err != nil {
		h.logError(w, r, err, http.StatusBadRequest)
		return
	}
	user, err := h.Service.Authorization.GetUserByToken(cookie.Value)

	err = h.Service.Reaction.ReactPost(1, user.Id, input.reaction)
	if err != nil {
		h.logError(w, r, err, http.StatusBadRequest)
		return
	}

	// http.Redirect(w, r, "/posts/1,", http.StatusSeeOther)
}
