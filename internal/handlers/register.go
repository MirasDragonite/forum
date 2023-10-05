package handlers

import (
	"fmt"
	"net/http"

	"forum/structs"
)

func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	err := r.ParseForm()
	if err != nil {
		return
	}

	input := structs.User{
		Username:       r.Form.Get("username"),
		Email:          r.Form.Get("email"),
		HashedPassword: r.Form.Get("password"),
	}
	fmt.Println(input)
	id, err := h.Service.Authorization.CreateUser(input)
	if err != nil {
		fmt.Println("Cannot create user")
		return
	}
	fmt.Println(id)
}
