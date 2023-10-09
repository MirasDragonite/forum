package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"forum/structs"
)

func (h *Handler) PostPage(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path[0:6] != "/post/" {
	// 	return
	// }

	post_id := r.URL.Path[6:]
	tmp, err := template.ParseFiles("./ui/templates/post_page.html")
	// h.logError(w, r, err, http.StatusInternalServerError)
	if err != nil {
		h.errorHandler(w, r, 500)
		return
	}
	if r.Method == http.MethodPost {

		cookie, err := r.Cookie("Token")
		if err != nil {
			if err != nil {
				http.Redirect(w, r, "/register", http.StatusSeeOther)
				return
			}
		}

		post, err := h.Service.PostRedact.GetPostBy("id", post_id)
		h.logError(w, r, err, http.StatusBadRequest)
		err = r.ParseForm()
		h.logError(w, r, err, http.StatusInternalServerError)

		user_id, err := h.Service.PostRedact.GetUSerID(cookie.Value)
		h.logError(w, r, err, http.StatusBadRequest)
		var comment structs.Comment
		comment.Dislike = 0
		comment.Like = 0
		comment.CommentAuthorID = user_id
		comment.PostID = post.Id
		comment.Content = r.FormValue("commentContent")
		post.Comments = append(post.Comments, comment)
		err = h.Service.CommentRedact.CreateComment(&comment)
		h.logError(w, r, err, http.StatusBadRequest)
		tmp.Execute(w, post)

	} else if r.Method == http.MethodGet {

		post, err := h.Service.PostRedact.GetPostBy("id", post_id)
		if err != nil {
			h.logError(w, r, err, http.StatusAccepted)
			return
		}
		// h.logError(w, r, err, http.StatusNotFound)

		tmp.Execute(w, post)
	} else {
		fmt.Println("mb here")
		w.Write([]byte("internal Server Error"))
	}
}
