package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"text/template"
	"time"
)

// logger
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

		id, err := h.Service.Authorization.GetUser(email, password)
		if err != nil {
			return
		}
		if id > 0 {
			time64 := time.Now().Unix()
			timeInt := string(time64)
			token := email + password + timeInt
			hashToken := md5.Sum([]byte(token))
			hashedToken := hex.EncodeToString(hashToken[:])
			h.Cache[hashedToken] = id
			livingTime := 60 * time.Minute
			expiration := time.Now().Add(livingTime)
			cookie := http.Cookie{Name: "Token", Value: url.QueryEscape(hashedToken), Expires: expiration}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/", http.StatusOK)
		}

	} else if r.Method == http.MethodGet {
		ts.Execute(w, "")
	} else {
		fmt.Println("rEDIcrect didnt work")
	}
}
