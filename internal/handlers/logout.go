package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handler) logOut(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorHandler(w, 400)
		return
	}
	cookie, err := r.Cookie("Token")

	err = h.Service.Authorization.DeleteToken(cookie)
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println("Cookie:", cookie)
	fmt.Println("not error")
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
