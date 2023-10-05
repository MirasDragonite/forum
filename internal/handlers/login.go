package handlers

import "net/http"

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	err := r.ParseForm()
	if err != nil {
		return
	}
	// username := r.Form.Get("username")
	// password := r.Form.Get("password")


}
