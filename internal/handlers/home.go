package handlers

import (
	"net/http"
	"text/template"
)

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	ts, err := template.ParseFiles("./ui/templates/home_page.html")
	if err != nil {
		h.logError(w, r, err, http.StatusInternalServerError)
		return
	}

	posts, err := h.Service.PostRedact.GetAllPosts()

	result := map[string]interface{}{
		"Posts": posts,
	}

	ts.Execute(w, result)
}
