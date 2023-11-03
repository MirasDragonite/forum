package handlers

// func (h *Handler) editPost(w http.ResponseWriter, r *http.Request) {
// 	tmp, err := template.ParseFiles("./ui/templates/edit_post_page.html")
// 	if err != nil {
// 		h.logError(w, r, err, http.StatusInternalServerError)
// 		return
// 	}
// 	err = r.ParseForm()
// 	if err != nil {
// 		h.logError(w, r, err, http.StatusBadRequest)
// 		return
// 	}

// 	title := r.Form.Get("postTitle")
// 	content := r.Form.Get("postContent")

// 	h.Service.PostRedact.
// 		tmp.Execute(w, nil)
// }
