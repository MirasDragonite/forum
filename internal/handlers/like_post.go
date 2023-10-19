package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"forum/structs"
	"net/http"
	"strconv"
)

var input struct {
	Reaction int64 `json:"reaction"`
}

func (h *Handler) likePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	fmt.Println("pOST LIKING")
	// err := json.NewDecoder(r.Body).Decode(&input)
	// if err != nil {
	// 	h.logError(w, r, err, http.StatusBadRequest)
	// 	return
	// }
	// fmt.Println("Data from input", input.Reaction)
	var button dataFromButton
	err = json.NewDecoder(r.Body).Decode(&button)

	fmt.Println("button:", button)
	switch button.Reaction {
	case "like":
		input.Reaction = 1
	case "dislike":
		input.Reaction = -1
	default:
		fmt.Println("he")
		h.logError(w, r, errors.New("No like"), http.StatusBadRequest)
		return
	}
	cookie, err := r.Cookie("Token")
	if err != nil {
		fmt.Println("Here3")
		h.logError(w, r, err, http.StatusNonAuthoritativeInfo)
		return
	}
	fmt.Println("url:", r.URL.Path)
	post_id_string := r.URL.Path[11:]
	fmt.Println("removed:", post_id_string)
	post_id, err := strconv.Atoi(post_id_string)
	if err != nil {
		fmt.Println("Here2")
		h.logError(w, r, err, http.StatusBadRequest)
		return
	}
	user, err := h.Service.Authorization.GetUserByToken(cookie.Value)

	err = h.Service.Reaction.ReactPost(int64(post_id), user.Id, input.Reaction)
	if err != nil {
		fmt.Println("Here1")
		h.logError(w, r, err, http.StatusBadRequest)
		return
	}
	fmt.Println("Here")
	likes, dislikes, err := h.Service.AllPostReactions(int64(post_id))
	if err != nil {
		return
	}
	reaction, err := h.Service.Reaction.GetPostReaction(user.Id, int64(post_id))
	res := structs.ResponseReaction{
		Likes:    likes,
		Dislikes: dislikes,
		Reaction: reaction,
	}
	fmt.Println(res)
	// link := fmt.Sprintf("/post/%v", post_id)
	// http.Redirect(w, r, link, http.StatusSeeOther)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(&res)
}
