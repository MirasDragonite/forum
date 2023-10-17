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
		h.errorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	ts, err := template.ParseFiles("./ui/templates/sign_up.html")
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
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(ok)
			h.logError(w, r, err, http.StatusBadRequest)
			return
		}
		fmt.Println(input)
		err = h.Service.Authorization.CreateUser(&input)

		if err != nil {
			fmt.Println("here1")
			h.errorLog(err.Error())
			h.errorHandler(w, r, 405)
			return
		}
		ok.Status = 200
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ok)

		return

	} else if r.Method == http.MethodGet {
		err := ts.Execute(w, "")
		if err != nil {
			h.logError(w, r, err, http.StatusInternalServerError)
		}
	} else {
		w.Write([]byte("Internal server error"))
	}
}
