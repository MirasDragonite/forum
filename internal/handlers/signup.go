package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"forum/structs"
)

func (h *Handler) signup(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		return
	}
	ts, err := template.ParseFiles("./ui/templates/sign_up.html")
	if err != nil {
		h.logError(w, r, err, http.StatusInternalServerError)
		return
	}

	err = r.ParseForm()
	if err != nil {
		h.logError(w, r, err, http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		var input structs.User
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			h.logError(w, r, err, http.StatusBadRequest)
			return
		}

		fmt.Println(input)
		err = h.Service.Authorization.CreateUser(&input)

		if err != nil {
			h.errorLog(err.Error())
			h.errorHandler(w, r, 405)
			return
		}
		// DONT DELETE THIS CODE LINES:
		// http.Redirect(w, r, "/signin", http.StatusSeeOther)
	} else if r.Method == http.MethodGet {
		err := ts.Execute(w, "")
		if err != nil {
			h.logError(w, r, err, http.StatusInternalServerError)
		}
	} else {
		w.Write([]byte("Internal server error"))
	}
}
