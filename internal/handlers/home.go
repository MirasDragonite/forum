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
	h.logError(w, r, err, http.StatusInternalServerError)
	ts.Execute(w, "")
}
