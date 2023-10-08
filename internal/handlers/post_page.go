package handlers

import (
	"forum/structs"
	"net/http"
	"text/template"
)

func (h *Handler) PostPage(w http.ResponseWriter, r *http.Request) {
	//  check url path
	if r.URL.Path != "/submit-post" {
		return
	}
	tmp, err := template.ParseFiles("./ui/templates/postpage.html")
	if err != nil {
		return
	}
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			return
		}
		cookie, err := r.Cookie("Token")
		if err != nil {
			return
		}
		var post structs.Post
		post.Title = r.Form.Get("postTitle")
		post.Topic = r.Form.Get("postTopic")
		post.Content = r.Form.Get("postContent")

		err = h.Service.PostRedact.CreatePost(&post, cookie.Value)
		if err != nil {
			return
		}

	} else if r.Method == http.MethodGet {
		tmp.Execute(w, nil)
	} else {
		w.Write([]byte("internal Server Error"))
	}

	tmp.Execute(w, nil)
}
