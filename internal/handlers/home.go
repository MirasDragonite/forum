package handlers

import (
	"net/http"
	"text/template"
)

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	ts, err := template.ParseFiles("./ui/templates/index.html")
	if err != nil {
		return
	}
	ts.Execute(w, "")
}
