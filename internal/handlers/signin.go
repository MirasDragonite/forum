package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"forum/structs"
)

// logger
func (h *Handler) signin(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/signin" {
		return
	}
	ts, err := template.ParseFiles("./ui/templates/sign_in.html")
	if err != nil {
		h.logError(w, r, err, http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodPost {
		var input *structs.User
		err = json.NewDecoder(r.Body).Decode(&input)

		cookie, err := h.Service.Authorization.GetUser(input.Email, input.HashedPassword)
		if err != nil {
			h.logError(w, r, err, http.StatusBadRequest)
			return
		}
		fmt.Println(cookie.Value)
		http.SetCookie(w, cookie)
		// http.Redirect(w, r, "/submit-post", http.StatusSeeOther)

	} else if r.Method == http.MethodGet {
		ts.Execute(w, "")
	} else {
	}
}
