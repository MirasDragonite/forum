package handlers

import (
	"encoding/json"
	"net/http"
	"text/template"

	"forum/structs"
)

func (h *Handler) PostPageCreate(w http.ResponseWriter, r *http.Request) {
	//  check url path
	if r.URL.Path != "/submit-post" {
		return
	}
	tmp, err := template.ParseFiles("./ui/templates/createpostpage.html")
	if err != nil {
		h.logError(w, r, err, http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {

		cookie, err := r.Cookie("Token")
		h.logError(w, r, err, http.StatusInternalServerError)
		var post *structs.Post
		// post.Title = r.Form.Get("postTitle")
		// post.Topic = r.Form.Get("postTopic")
		// post.Content = r.Form.Get("postContent")
		err = json.NewDecoder(r.Body).Decode(&post)
		post.PostAuthorName, err = h.Service.PostRedact.GetUserName(cookie.Value)
		if err != nil {
			h.logError(w, r, err, http.StatusInternalServerError)
			return
		}

		err = h.Service.PostRedact.CreatePost(post, cookie.Value)
		h.logError(w, r, err, http.StatusBadRequest)

	} else if r.Method == http.MethodGet {
		cookie, err := r.Cookie("Token")
		h.logError(w, r, err, http.StatusInternalServerError)
		user, err := h.Service.Authorization.GetUserByToken(cookie.Value)
		if err != nil {
			h.logError(w, r, err, http.StatusInternalServerError)
			return
		}

		tmp.Execute(w, user)
	} else {
		w.Write([]byte("internal Server Error"))
	}
}
