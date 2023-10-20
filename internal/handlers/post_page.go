package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"forum/structs"
	"net/http"
	"strings"
	"text/template"
)

func (h *Handler) PostPage(w http.ResponseWriter, r *http.Request) {
	post_id := r.URL.Path[6:]
	if strings.TrimSpace(post_id) == "" || post_id[0] == '0' {
		h.errorHandler(w, r, http.StatusBadRequest)
		return
	}
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
			fmt.Println("POST METHOD")
			fmt.Println(err.Error)
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
		user, err := h.Service.Authorization.GetUserByToken(cookie.Value)
		comment.CommentAuthorName = user.Username
		post.Comments = append(post.Comments, *comment)
		err = h.Service.CommentRedact.CreateComment(comment)
		if err != nil {
			h.logError(w, r, err, http.StatusBadRequest)
			return
		}

		res := &structs.Data{
			Status: int(comment.CommentID),
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(&res)
		if err != nil {
			return
		}
		tmp.Execute(w, post)

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
			h.logError(w, r, err, http.StatusBadRequest)
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

		post.Comments = comments

		result := map[string]interface{}{
			"Post":     post,
			"Likes":    likes,
			"Dislikes": dislikes,
		}
		tmp.Execute(w, result)
	} else {
		h.logError(w, r, errors.New("Wrong Method"), http.StatusMethodNotAllowed)
		return
	}
}
