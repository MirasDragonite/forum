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
		input := structs.User{
			Username:       r.Form.Get("username"),
			Email:          r.Form.Get("email"),
			HashedPassword: r.Form.Get("password"),
		}

		id, err := h.Service.Authorization.CreateUser(input)
		if err != nil {
			fmt.Println("Cannot create user")
			return
		}

		ts.Execute(w, id)
	} else if r.Method == http.MethodGet {
		ts.Execute(w, "")
	} else {
		w.Write([]byte("Internal server error"))
	}
}
