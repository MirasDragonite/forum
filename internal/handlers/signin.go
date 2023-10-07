package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

func (h *Handler) signin(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/signin" {
		return
	}
	ts, err := template.ParseFiles("./ui/templates/signin.html")
	err = r.ParseForm()
	if err != nil {
		return
	}
	if r.Method == http.MethodPost {
		email := r.Form.Get("email")
		password := r.Form.Get("password")

		cookie, id, hashedToken, err := h.Service.Authorization.GetUser(email, password)
		if err != nil {
			return
		}
		h.Cache[hashedToken] = id
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		fmt.Println("REdirec..")

	} else if r.Method == http.MethodGet {
		ts.Execute(w, "")
	} else {
		fmt.Println("rEDIcrect didnt work")
	}
}
