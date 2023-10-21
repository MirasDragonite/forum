package handlers

import (
	"errors"
	"net/http"
	"text/template"

	"forum/structs"
)

// logger
func (h *Handler) signin(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/signin" {
		h.errorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	ts, err := template.ParseFiles("./ui/templates/sign_in.html")
	if err != nil {
		h.logError(w, r, err, http.StatusInternalServerError)
		return
	}

	err = r.ParseForm()
	if err != nil {
		h.logError(w, r, err, 500)
		return
	}
	if r.Method == http.MethodPost {
		var input structs.User
		input.Email = r.Form.Get("email")
		input.HashedPassword = r.Form.Get("password")
		cookie, err := h.Service.Authorization.GetUser(input.Email, input.HashedPassword)
		if err != nil {
			h.logError(w, r, err, http.StatusBadRequest)
			return
		}

		http.SetCookie(w, cookie)

		http.Redirect(w, r, "/", http.StatusSeeOther)
		// ok.Status = 200
		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		// json.NewEncoder(w).Encode(ok)

		return

	} else if r.Method == http.MethodGet {
		ts.Execute(w, "")
	} else {
		h.logError(w, r, errors.New("Wrong Method"), http.StatusMethodNotAllowed)
		return
	}
}
