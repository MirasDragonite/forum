package handlers

import (
	"html/template"
	"net/http"
)

const (
	pathToErrorPage = "./ui/templates/error_page.html"
)

func (h *Handler) errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	// var resp structs.Data
	w.WriteHeader(status)
	errs := "404"
	switch status {
	case 400:
		errs = "400"
	case 404:
		errs = "404"
	case 405:
		errs = "405"
	case 500:
		errs = "500"
	case 401:

		return

	}
	page, err := template.ParseFiles(pathToErrorPage)
	if err != nil {
		w.Write([]byte("Internal Server Error"))
		return
	}
	if status == 401 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return

	}
	err = page.Execute(w, errs)
	if err != nil {
		w.Write([]byte("Internal Server Error"))
		return
	}
	return
}

func (h *Handler) logError(w http.ResponseWriter, r *http.Request, err error, status int) {
	if err != nil {
		h.errorLog(err.Error())
		h.errorHandler(w, r, status)
		return
	}
}
