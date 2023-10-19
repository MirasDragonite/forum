package handlers

import (
	"encoding/json"
	"forum/structs"
	"net/http"
	"text/template"
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
	ok := structs.Data{}

	if r.Method == http.MethodPost {
		var input structs.User
		err = json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			ok.Status = 400
			h.logError(w, r, err, http.StatusBadRequest)
			return
		}

		cookie, err := h.Service.Authorization.GetUser(input.Email, input.HashedPassword)
		if err != nil {
			h.logError(w, r, err, http.StatusBadRequest)
			return
		}

		http.SetCookie(w, cookie)
		ok.Status = 200
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ok)

		return

	} else if r.Method == http.MethodGet {
		ts.Execute(w, "")
	} else {
	}
}
