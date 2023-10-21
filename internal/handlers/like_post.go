package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

var input struct {
	Reaction int64 `json:"reaction"`
}

func (h *Handler) likePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.logError(w, r, errors.New("Wrong Method"), http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()

	var button dataFromButton
	// err := json.NewDecoder(r.Body).Decode(&button)
	button.Reaction = r.Form.Get("button")
	switch button.Reaction {
	case "like":
		input.Reaction = 1
	case "dislike":
		input.Reaction = -1
	default:
		h.logError(w, r, errors.New("No like"), http.StatusBadRequest)
		return
	}
	cookie, err := r.Cookie("Token")
	if err != nil {
		h.logError(w, r, err, http.StatusNonAuthoritativeInfo)
		return
	}
	post_id_string := r.URL.Path[11:]
	post_id, err := strconv.Atoi(post_id_string)
	if err != nil {
		h.logError(w, r, err, http.StatusBadRequest)
		return
	}

	user, err := h.Service.Authorization.GetUserByToken(cookie.Value)
	err = h.Service.Reaction.ReactPost(int64(post_id), user.Id, input.Reaction)
	if err != nil {
		h.logError(w, r, err, http.StatusBadRequest)
		return
	}

	link := fmt.Sprintf("/post/%v", post_id)
	http.Redirect(w, r, link, http.StatusSeeOther)
	return
	// likes, dislikes, err := h.Service.AllPostReactions(int64(post_id))
	// if err != nil {
	// 	return
	// }

	// reaction, err := h.Service.Reaction.GetPostReaction(user.Id, int64(post_id))
	// res := structs.ResponseReaction{
	// 	Likes:    likes,
	// 	Dislikes: dislikes,
	// 	Reaction: reaction,
	// }
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)

	// err = json.NewEncoder(w).Encode(&res)
}
