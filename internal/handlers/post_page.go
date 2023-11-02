package handlers

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"forum/structs"
)

func (h *Handler) PostPage(w http.ResponseWriter, r *http.Request) {
	post_id := r.URL.Path[6:]
	if strings.TrimSpace(post_id) == "" || post_id[0] == '0' {
		h.errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmp, err := template.ParseFiles("./ui/templates/post_page.html")
	if err != nil {
		h.logError(w, r, err, http.StatusInternalServerError)

		return
	}

	err = r.ParseForm()
	if err != nil {
		h.logError(w, r, err, http.StatusInternalServerError)
		return
	}
	result := map[string]interface{}{
		"Post":     nil,
		"Likes":    nil,
		"Dislikes": nil,
		"Empty":    nil,
		"Logged":   nil,
	}
	if r.Method == http.MethodPost {
		cookie, err := r.Cookie("Token")
		if err != nil {
			http.Redirect(w, r, "/register", http.StatusSeeOther)
			h.logError(w, r, errors.New("Wrong Method"), http.StatusUnauthorized)
			return

		}
		user_id, err := h.Service.PostRedact.GetUSerID(cookie.Value)
		if err != nil {
			h.logError(w, r, err, http.StatusInternalServerError)
			return
		}

		post, err := h.Service.PostRedact.GetPostBy("id", post_id, user_id)
		if err != nil {
			h.logError(w, r, err, http.StatusNotFound)
			return
		}

		likes, dislikes, err := h.Service.Reaction.AllPostReactions(post.Id)
		if err != nil {
			h.logError(w, r, err, http.StatusInternalServerError)
			return
		}
		if user_id != 0 {
			result["Logged"] = true
		}
		comment := &structs.Comment{}

		comment.Content = strings.TrimSpace(r.Form.Get("commentContent"))
		result["Post"] = post
		result["Likes"] = likes
		result["Dislikes"] = dislikes
		if len(comment.Content) < 2 || len(comment.Content) > 100 {
			w.WriteHeader(http.StatusBadRequest)
			result["Empty"] = true
			tmp.Execute(w, result)
			return
		}

		comment.Dislike = 0
		comment.Like = 0
		comment.CommentAuthorID = user_id
		comment.PostID = post.Id
		user, err := h.Service.Authorization.GetUserByToken(cookie.Value)
		comment.CommentAuthorName = user.Username
		post.Comments = append(post.Comments, *comment)
		err = h.Service.CommentRedact.CreateComment(comment, post.PostAuthorID)
		if err != nil {
			if err.Error() == "Can't be empty" {
			} else {
				h.logError(w, r, err, http.StatusBadRequest)
				return
			}
		}

		link := fmt.Sprintf("/post/%v", post.Id)
		http.Redirect(w, r, link, http.StatusSeeOther)
		tmp.Execute(w, post)
		return

	} else if r.Method == http.MethodGet {
		var user_id int64
		cookie, err := r.Cookie("Token")

		if err == nil {

			user_id, err = h.Service.PostRedact.GetUSerID(cookie.Value)
			if err != nil {
				user_id = 0
			}
		} else {
			user_id = 0
		}

		post, err := h.Service.PostRedact.GetPostBy("id", post_id, user_id)
		if err != nil {
			h.logError(w, r, err, http.StatusNotFound)
			return
		}

		likes, dislikes, err := h.Service.Reaction.AllPostReactions(post.Id)
		if err != nil {
			h.logError(w, r, err, http.StatusInternalServerError)
			return
		}

		comments, err := h.Service.CommentRedact.GetAllComments(post.Id, user_id)
		if err != nil {
			h.logError(w, r, err, http.StatusInternalServerError)
			return
		}
		if user_id != 0 {
			result["Logged"] = true
		}
		post.Comments = comments
		result["Post"] = post
		result["Likes"] = likes
		result["Dislikes"] = dislikes
		tmp.Execute(w, result)
	} else {
		h.logError(w, r, errors.New("Wrong Method"), http.StatusMethodNotAllowed)
		return
	}
}
