package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type dataFromButton struct {
	Reaction string `json:"reaction"`
}

func (h *Handler) likeComment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	fmt.Println("COMMENT LIKING")
	// err := json.NewDecoder(r.Body).Decode(&input)
	// if err != nil {
	// 	h.logError(w, r, err, http.StatusBadRequest)
	// 	return
	// }
	// fmt.Println("Data from input", input.Reaction)
	var button *dataFromButton
	err = json.NewDecoder(r.Body).Decode(&button)

	fmt.Println("commButton:", button)
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
	comment_id_string := r.URL.Path[14:]
	fmt.Println("removed:", comment_id_string)
	comment_id, err := strconv.Atoi(comment_id_string)
	if err != nil {
		fmt.Println("Here2")
		h.logError(w, r, err, http.StatusBadRequest)
		return
	}
	user, err := h.Service.Authorization.GetUserByToken(cookie.Value)

	err = h.Service.Reaction.ReactComment(int64(comment_id), user.Id, input.Reaction)
	if err != nil {
		fmt.Println("Here1")
		h.logError(w, r, err, http.StatusBadRequest)
		return
	}
	fmt.Println("Here")
	comment, err := h.Service.GetCommentByID(int64(comment_id))
	link := fmt.Sprintf("/post/%v", comment.PostID)
	http.Redirect(w, r, link, http.StatusSeeOther)
}
