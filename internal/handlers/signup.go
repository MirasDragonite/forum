package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"forum/structs"
)

type Data struct {
	status int
}

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
		// if err != nil {
		// 	return
		// }
		err = json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("here2")
			h.logError(w, r, err, http.StatusBadRequest)
			return
		}
		// input.Username = r.Form.Get("username")
		// input.Email = r.Form.Get("email")
		// input.HashedPassword = r.Form.Get("password")
		fmt.Println(input)
		err = h.Service.Authorization.CreateUser(&input)

		if err != nil {
			fmt.Println("here1")
			h.errorLog(err.Error())
			h.errorHandler(w, r, 405)
			return
		}

		// ok := Data{
		// 	status: 200,
		// }
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		mapaa := make(map[string]int)
		mapaa["status"] = 200
		jsonResp, err := json.Marshal(mapaa)
		if err != nil {
			return
		}

		// DONT DELETE THIS CODE LINES:
		fmt.Println("Before redirect")
		// w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResp)
		// http.Redirect(w, r, "/signin", http.StatusSeeOther)
		return
		// w.WriteHeader(http.StatusOK)

	} else if r.Method == http.MethodGet {
		err := ts.Execute(w, "")
		if err != nil {
			h.logError(w, r, err, http.StatusInternalServerError)
		}
	} else {
		w.Write([]byte("Internal server error"))
	}
}
