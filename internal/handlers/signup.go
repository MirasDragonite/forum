package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"text/template"

	"forum/structs"
)

func (h *Handler) signup(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		h.errorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	ts, err := template.ParseFiles("./ui/templates/sign_up.html")
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

		input.Username = r.Form.Get("username")
		input.Email = r.Form.Get("email")
		input.HashedPassword = r.Form.Get("password")
		// err = json.NewDecoder(r.Body).Decode(&input)
		// if err != nil {
		// 	ok.Status = 400
		// 	w.Header().Set("Content-Type", "application/json")
		// 	w.WriteHeader(http.StatusOK)
		// 	json.NewEncoder(w).Encode(ok)
		// 	h.logError(w, r, err, http.StatusBadRequest)
		// 	return
		// }
		fmt.Println(input)
		err = h.Service.Authorization.CreateUser(&input)

		if err != nil {
			h.errorLog(err.Error())
			h.errorHandler(w, r, 400)
			return
		}
		http.Redirect(w, r, "/signin", http.StatusSeeOther)

		return

	} else if r.Method == http.MethodGet {
		err := ts.Execute(w, "")
		if err != nil {
			h.logError(w, r, err, http.StatusInternalServerError)
		}
	} else {
		h.logError(w, r, errors.New("Wrong Method"), http.StatusMethodNotAllowed)
		return
	}
}
