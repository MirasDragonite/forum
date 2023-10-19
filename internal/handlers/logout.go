package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handler) logOut(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.errorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}

	cookie, err := r.Cookie("Token")
	err = h.Service.Authorization.DeleteToken(cookie)
	h.logError(w, r, err, http.StatusUnauthorized)
	fmt.Println("Cookie:", cookie)
	fmt.Println("not error")
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/register", http.StatusSeeOther)
}
