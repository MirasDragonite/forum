package handlers

import (
	"net/http"
	"text/template"
)

func (h *Handler) profile(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/templates/profilepage.html")
	h.logError(w, r, err, http.StatusInternalServerError)

	ts.Execute(w, nil)
}
