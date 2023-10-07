package handlers

import (
	"net/http"
	"text/template"
)

func (h *Handler) profile(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/templates/profilepage.html")
	if err != nil {
		return
	}

	ts.Execute(w, nil)
}
