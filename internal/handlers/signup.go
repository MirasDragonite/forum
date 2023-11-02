package handlers

import (
	"errors"
	"forum/structs"
	"html/template"
	"net/http"
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
	errorsCatch := map[string]interface{}{
		"Unique":         nil,
		"PasswordLength": nil,
		"UserLength":     nil,
		"EmailLength":    nil,
	}
	if r.Method == http.MethodPost {

		var input structs.User

		input.Username = r.Form.Get("username")
		input.Email = r.Form.Get("email")
		input.HashedPassword = r.Form.Get("password")
		err = h.Service.Authorization.CreateUser(&input)

		if err != nil {

			if err.Error() == "UNIQUE constraint failed: users.email" {
				errorsCatch["Unique"] = true
			} else if err.Error() == "The length of the password is not up to standard " {
				errorsCatch["PasswordLength"] = true
			} else if err.Error() == "The length of the email is not up to standard " {
				errorsCatch["EmailLength"] = true
			} else if err.Error() == "The length of the user is not up to standard " {
				errorsCatch["UserLength"] = true
			}
			if errorsCatch["Unique"] == nil && errorsCatch["PasswordLength"] == nil && errorsCatch["UserLength"] == nil && errorsCatch["EmailLength"] == nil {
				h.errorLog(err.Error())
				h.errorHandler(w, r, 400)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			ts.Execute(w, errorsCatch)
			return

		}
		http.Redirect(w, r, "/signin", http.StatusSeeOther)

		return

	} else if r.Method == http.MethodGet {
		err := ts.Execute(w, errorsCatch)
		if err != nil {
			h.logError(w, r, err, http.StatusInternalServerError)
		}
	} else {
		h.logError(w, r, errors.New("Wrong Method"), http.StatusMethodNotAllowed)
		return
	}
}
