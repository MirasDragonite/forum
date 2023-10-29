package handlers

import (
	"net/http"
	"text/template"
)

func (h *Handler) notify(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/notify" {
		h.errorHandler(w, r, http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		tmp, err := template.ParseFiles("./ui/templates/notification.html")
		if err != nil {
			h.logError(w, r, err, http.StatusInternalServerError)
			return
		}

		cookie, err := r.Cookie("Token")

		user, err := h.Service.Authorization.GetUserByToken(cookie.Value)

		h.Service.Notification.AllUserNotifications(user.Id)

		tmp.Execute(w, nil)
	}
}
