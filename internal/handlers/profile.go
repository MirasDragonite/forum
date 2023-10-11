package handlers

import (
	"net/http"
	"text/template"
)

func (h *Handler) profile(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/templates/profile_page.html")
	if err != nil {
		h.logError(w, r, err, http.StatusInternalServerError)
		return
	}
	ts.Execute(w, nil)
}
