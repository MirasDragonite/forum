package handlers

import (
	"errors"
	"fmt"
	"forum/structs"
	"html/template"
	"net/http"
	"strings"
)

func (h *Handler) PostPageCreate(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/submit-post" {
		h.errorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		h.logError(w, r, err, http.StatusInternalServerError)
		return
	}
	tmp, err := template.ParseFiles("./ui/templates/create_post_page.html")
	if err != nil {
		h.logError(w, r, err, http.StatusInternalServerError)
		return
	}
	result := map[string]interface{}{
		"User":         nil,
		"Logged":       false,
		"EmptyTitle":   false,
		"EmptyContent": false,
	}

	cookie, err := r.Cookie("Token")
	if err != nil {
		if err != http.ErrNoCookie {
			h.logError(w, r, err, http.StatusInternalServerError)
			return
		}
	}

	user, err := h.Service.Authorization.GetUserByToken(cookie.Value)
	if err != nil {
		h.logError(w, r, err, http.StatusInternalServerError)
		return
	}
	logged := false

	if user != nil {
		logged = true
	}
	result["Logged"] = logged
	result["User"] = user
	if r.Method == http.MethodPost {

		post := &structs.Post{}

		post.Title = strings.TrimSpace(r.Form.Get("postTitle"))
		if r.Form.Get("postTopicJava") == "Java" {
			post.Topic = append(post.Topic, r.Form.Get("postTopicJava"))
		}
		if r.Form.Get("postTopicKotlin") == "Kotlin" {
			post.Topic = append(post.Topic, r.Form.Get("postTopicKotlin"))
		}
		if r.Form.Get("postTopicPython") == "Python" {
			post.Topic = append(post.Topic, r.Form.Get("postTopicPython"))
		}
		if r.Form.Get("postTopicInput") != "" {
			post.Topic = append(post.Topic, r.Form.Get("postTopicInput"))
		}
		post.Content = strings.TrimSpace(r.Form.Get("postContent"))

		if len(post.Title) < 5 || len(post.Title) > 50 {
			w.WriteHeader(http.StatusBadRequest)
			result["EmptyTitle"] = true
			tmp.Execute(w, result)
			return
		}
		if len(post.Content) < 15 || len(post.Content) > 250 {
			w.WriteHeader(http.StatusBadRequest)
			result["EmptyContent"] = true
			tmp.Execute(w, result)
			return
		}

		topicStr := ""
		for _, topic := range post.Topic {
			topicStr += topic + "|"
		}
		if len(topicStr) == 0 {
			post.TopicString = ""
		} else {
			topicStr = topicStr[:len(topicStr)-1]
			post.TopicString = topicStr
		}

		post.PostAuthorName, err = h.Service.PostRedact.GetUserName(cookie.Value)
		if err != nil {
			h.logError(w, r, err, http.StatusInternalServerError)
			return
		}

		err = h.Service.PostRedact.CreatePost(post, cookie.Value)
		if err != nil {
			h.logError(w, r, errors.New("Wrong Method"), http.StatusInternalServerError)
			return
		}
		link := fmt.Sprintf("/post/%v", post.Id)
		http.Redirect(w, r, link, http.StatusSeeOther)
		return

	} else if r.Method == http.MethodGet {
		tmp.Execute(w, result)
	} else {
		h.logError(w, r, errors.New("Wrong Method"), http.StatusMethodNotAllowed)
		return
	}
}
