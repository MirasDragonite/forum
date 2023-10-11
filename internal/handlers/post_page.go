package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"forum/structs"
)

func (h *Handler) PostPage(w http.ResponseWriter, r *http.Request) {
	post_id := r.URL.Path[6:]
	tmp, err := template.ParseFiles("./ui/templates/post_page.html")
	// h.logError(w, r, err, http.StatusInternalServerError)
	if err != nil {
		fmt.Println(err.Error())
		h.errorHandler(w, r, 500)
		return
	}
	if r.Method == http.MethodPost {

		cookie, err := r.Cookie("Token")
		if err != nil {
			// DONT DELETE THIS CODE LINES:
			// http.Redirect(w, r, "/register", http.StatusSeeOther)
			return
		}

		post, err := h.Service.PostRedact.GetPostBy("id", post_id)
		if err != nil {
			h.logError(w, r, err, http.StatusBadRequest)
			return
		}

		user_id, err := h.Service.PostRedact.GetUSerID(cookie.Value)
		if err != nil {
			h.logError(w, r, err, http.StatusBadRequest)
			return
		}
		var comment *structs.Comment
		err = json.NewDecoder(r.Body).Decode(&comment)
		if err != nil {
			h.logError(w, r, err, http.StatusBadRequest)
			return
		}
		comment.Dislike = 0
		comment.Like = 0
		comment.CommentAuthorID = user_id
		comment.PostID = post.Id

		post.Comments = append(post.Comments, *comment)
		err = h.Service.CommentRedact.CreateComment(comment)
		if err != nil {
			h.logError(w, r, err, http.StatusBadRequest)
			return
		}
		tmp.Execute(w, post)

	} else if r.Method == http.MethodGet {

		post, err := h.Service.PostRedact.GetPostBy("id", post_id)
		fmt.Println("handlers: ", post)
		if err != nil {
			fmt.Println("handlers: ", post)
			h.logError(w, r, err, http.StatusAccepted)
			return
		}
		// h.logError(w, r, err, http.StatusNotFound)

		tmp.Execute(w, post)
	} else {
		fmt.Println("else here")
		w.Write([]byte("internal Server Error"))
	}
}
