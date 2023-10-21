package handlers

import (
	"database/sql"
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
	errorsCatch := map[string]interface{}{
		"PasswordLength":        nil,
		"EmailLength":           nil,
		"ValidEmail":            nil,
		"PasswordNotCompatible": nil,
	}
	if r.Method == http.MethodPost {
		var input structs.User
		input.Email = r.Form.Get("email")
		input.HashedPassword = r.Form.Get("password")
		cookie, err := h.Service.Authorization.GetUser(input.Email, input.HashedPassword)
		if err != nil {
			if err.Error() == "The length of the password is not up to standard " {
				errorsCatch["PasswordLength"] = true
			} else if err.Error() == "The length of the email is not up to standard " {
				errorsCatch["UserLength"] = true
			} else if err.Error() == "Not valid email" {
				errorsCatch["ValidEmail"] = true
			} else if err.Error() == sql.ErrNoRows.Error() {
				errorsCatch["ValidEmail"] = true
			} else if err.Error() == "Passwords not compatible" {
				errorsCatch["PasswordNotCompatible"] = true
			} else {
				h.errorLog(err.Error())
				h.errorHandler(w, r, 400)
				return
			}

			ts.Execute(w, errorsCatch)
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
		ts.Execute(w, errorsCatch)
	} else {
		h.logError(w, r, errors.New("Wrong Method"), http.StatusMethodNotAllowed)
		return
	}
}
