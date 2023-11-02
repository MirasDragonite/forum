package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func (h *Handler) notify(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/notify" {
		h.errorHandler(w, r, http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		tmp, err := template.ParseFiles("./ui/templates/notifications_page.html")
		if err != nil {
			h.logError(w, r, err, http.StatusInternalServerError)
			return
		}

		cookie, err := r.Cookie("Token")

		user, err := h.Service.Authorization.GetUserByToken(cookie.Value)
		fmt.Println("UserID", user.Id)
		posts, err := h.Service.Notification.AllUserNotifications(user.Id)
		if err != nil {
			h.logError(w, r, err, http.StatusInternalServerError)
			return
		}

		result := map[string]interface{}{
			"Logged":        nil,
			"User":          user,
			"Notifications": posts,
		}
		tmp.Execute(w, result)
	}
}
