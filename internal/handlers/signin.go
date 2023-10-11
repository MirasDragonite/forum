package handlers

import (
	"net/http"
	"text/template"
)

// logger
func (h *Handler) signin(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/signin" {
		return
	}
	ts, err := template.ParseFiles("./ui/templates/signin.html")
	err = r.ParseForm()
	h.logError(w, r, err, http.StatusInternalServerError)
	if r.Method == http.MethodPost {
		email := r.Form.Get("email")
		password := r.Form.Get("password")

		cookie, err := h.Service.Authorization.GetUser(email, password)
		h.logError(w, r, err, http.StatusBadRequest)

		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/submit-post", http.StatusSeeOther)

	} else if r.Method == http.MethodGet {
		ts.Execute(w, "")
	} else {
	}
}
