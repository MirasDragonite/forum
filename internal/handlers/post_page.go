package handlers

import (
	"forum/structs"
	"net/http"
	"strconv"
	"text/template"
)

func (h *Handler) PostPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path[0:6] == "/post/" {
		return
	}
	post_id := r.URL.Path[6:]
	tmp, err := template.ParseFiles("./ui/templates/post_page.html")
	if err != nil {
		return
	}
	if r.Method == http.MethodPost {	
		post, err := h.Service.PostRedact.GetPostBy("id", post_id)
		
		if err != nil {
			return
		}
		err = r.ParseForm()
		if err != nil {
			return
		}
		cookie, err := r.Cookie("Token")
		if err != nil {
			return
		}
		user_id, err:= h.Service.PostRedact.GetUSerID(cookie.Value)
		if err != nil {
			return
		}
		var comment structs.Comment
		comment.Dislike =0
		comment.Like = 0
		comment.CommentAuthorID= user_id
		comment.PostID = post.Id
		comment.Content = r.FormValue("content")
		post.Comments = append(post.Comments, comment)
		err = h.Service.CommentRedact.CreateComment(&comment)
		if err != nil {
			return
		}
	} else if r.Method == http.MethodGet {
		post, err := h.Service.PostRedact.GetPostBy("id", post_id)
		
		if err != nil {
			return
		}
		tmp.Execute(w, post)
	} else {
		w.Write([]byte("internal Server Error"))
	}
}
