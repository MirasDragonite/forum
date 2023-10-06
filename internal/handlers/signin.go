package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handler) signin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	err := r.ParseForm()
	if err != nil {
		return
	}
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	id, err := h.Service.Authorization.GetUser(email, password)
	cookie := http.Cookie{Name: "token"}
	fmt.Println(id)
}
