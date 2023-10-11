package handlers

import (
	"net/http"
	"text/template"

	"forum/structs"
)

func (h *Handler) signup(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		return
	}
	ts, err := template.ParseFiles("./ui/templates/signup.html")
	err = r.ParseForm()
	h.logError(w, r, err, http.StatusInternalServerError)

	if r.Method == http.MethodPost {

		input := structs.CreateUser(r.Form.Get("username"), r.Form.Get("email"), r.Form.Get("password"))
		err := h.Service.Authorization.CreateUser(input)
		if err != nil {
			h.errorLog(err.Error())
			h.errorHandler(w, r, 405)
			return
		}

		http.Redirect(w, r, "/signin", http.StatusSeeOther)
	} else if r.Method == http.MethodGet {
		ts.Execute(w, "")
	} else {
		w.Write([]byte("Internal server error"))
	}
}
