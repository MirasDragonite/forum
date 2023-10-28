package handlers

import (
	"net/http"
)

func (h *Handler) logOut(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.errorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}

	cookie, err := r.Cookie("Token")
	err = h.Service.Authorization.DeleteToken(cookie)
	if err != nil {
		h.logError(w, r, err, http.StatusUnauthorized)
		return
	}
	// cookieMsG := fmt.Sprintf("Cookie:", cookie)

	h.infoLog("Sign out from session")
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/register", http.StatusSeeOther)
}
