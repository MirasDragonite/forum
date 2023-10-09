package handlers

import (
	"net/http"
	"text/template"

	"forum/structs"
)

func (h *Handler) PostPageCreate(w http.ResponseWriter, r *http.Request) {
	//  check url path
	if r.URL.Path != "/submit-post" {
		return
	}
	tmp, err := template.ParseFiles("./ui/templates/post_page_creation.html")
	h.logError(w, r, err, http.StatusInternalServerError)
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		h.logError(w, r, err, http.StatusInternalServerError)
		cookie, err := r.Cookie("Token")
		h.logError(w, r, err, http.StatusInternalServerError)
		var post structs.Post
		post.Title = r.Form.Get("postTitle")
		post.Topic = r.Form.Get("postTopic")
		post.Content = r.Form.Get("postContent")

		err = h.Service.PostRedact.CreatePost(&post, cookie.Value)
		h.logError(w, r, err, http.StatusBadRequest)

	} else if r.Method == http.MethodGet {
		tmp.Execute(w, nil)
	} else {
		w.Write([]byte("internal Server Error"))
	}

	tmp.Execute(w, nil)
}
