package handlers

import (
	"fmt"
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
	if err != nil {
		return
	}

	if r.Method == http.MethodPost {

		input := structs.CreateUser(r.Form.Get("username"), r.Form.Get("email"), r.Form.Get("password"))
		id, err := h.Service.Authorization.CreateUser(input)
		if err != nil {
			fmt.Println("Cannot create user")
			errorHandler(w, 405)
			return
		}

		input.Id = id
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
	} else if r.Method == http.MethodGet {
		ts.Execute(w, "")
	} else {
		w.Write([]byte("Internal server error"))
	}
}
