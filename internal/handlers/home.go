package handlers

import (
	"errors"
	"net/http"
	"text/template"
)

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.logError(w, r, errors.New("Wrong Method"), http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {

		h.errorHandler(w, r, http.StatusNotFound)
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
