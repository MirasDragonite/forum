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
	fmt.Println("HERE")
	post_id := r.URL.Path[6:]
	tmp, err := template.ParseFiles("./ui/templates/post_page.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if r.Method == http.MethodPost {
		fmt.Println("Post here")
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
		user_id, err := h.Service.PostRedact.GetUSerID(cookie.Value)
		if err != nil {
			return
		}
		var comment structs.Comment
		comment.Dislike = 0
		comment.Like = 0
		comment.CommentAuthorID = user_id
		comment.PostID = post.Id
		comment.Content = r.FormValue("commentContent")
		post.Comments = append(post.Comments, comment)
		err = h.Service.CommentRedact.CreateComment(&comment)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		tmp.Execute(w, post)
	} else if r.Method == http.MethodGet {
		fmt.Println("GET HERE")
		post, err := h.Service.PostRedact.GetPostBy("id", post_id)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(post)
		tmp.Execute(w, post)
	} else {
		fmt.Println("mb here")
		w.Write([]byte("internal Server Error"))
	}
}
