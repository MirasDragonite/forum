package handlers

import (
	"fmt"
	"net/http"
)

var input struct {
	Reaction int64 `json:"reaction"`
}

func (h *Handler) likePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	fmt.Println("gfd")
	// err := json.NewDecoder(r.Body).Decode(&input)
	// if err != nil {
	// 	h.logError(w, r, err, http.StatusBadRequest)
	// 	return
	// }
	// fmt.Println("Data from input", input.Reaction)
	button := r.Form.Get("button")

	fmt.Println("button:", button)
	switch button {
	case "like":
		input.Reaction = 1
	case "dislike":
		input.Reaction = -1
	default:
		fmt.Println("he")
		input.Reaction = 1
	}
	cookie, err := r.Cookie("Token")
	if err != nil {
		fmt.Println("Here3")
		h.logError(w, r, err, http.StatusNonAuthoritativeInfo)
		return
	}
	// post_id_string := r.URL.Path[6:]
	// post_id, err := strconv.Atoi(post_id_string)
	if err != nil {
		fmt.Println("Here2")
		h.logError(w, r, err, http.StatusBadRequest)
		return
	}
	user, err := h.Service.Authorization.GetUserByToken(cookie.Value)

	err = h.Service.Reaction.ReactPost(1, user.Id, input.Reaction)
	if err != nil {
		fmt.Println("Here1")
		h.logError(w, r, err, http.StatusBadRequest)
		return
	}
	fmt.Println("Here")
	http.Redirect(w, r, "/post/1", http.StatusSeeOther)
}
