package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"forum/structs"
)

type dataFromButton struct {
	Reaction string `json:"reaction"`
}

func (h *Handler) likeComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.logError(w, r, errors.New("Wrong Method"), http.StatusMethodNotAllowed)
		return
	}

	var button *dataFromButton
	err := json.NewDecoder(r.Body).Decode(&button)
	if err != nil {
		h.logError(w, r, errors.New("Decode error"), http.StatusBadRequest)
		return
	}

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
		h.logError(w, r, err, http.StatusUnauthorized)
		return
	}
	comment_id_string := r.URL.Path[14:]
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

	err = h.Service.Reaction.ReactComment(int64(comment_id), user.Id, input.Reaction)
	if err != nil {
		h.logError(w, r, err, http.StatusBadRequest)
		return
	}

	likes, dislikes, err := h.Service.Reaction.AllCommentReactions(int64(comment_id))
	if err != nil {
		h.logError(w, r, errors.New("Can't get all comment reactions"), http.StatusInternalServerError)
		return
	}
	reaction, err := h.Service.Reaction.GetCommentReaction(user.Id, int64(comment_id))

	res := structs.ResponseReaction{
		Likes:    likes,
		Dislikes: dislikes,
		Reaction: reaction,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(&res)
	if err != nil {
		h.logError(w, r, errors.New("Encode error"), http.StatusInternalServerError)
		return
	}
}
