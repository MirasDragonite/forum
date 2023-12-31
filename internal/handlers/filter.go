package handlers

import (
	"errors"
	"net/http"

	"forum/structs"
)

func (h *Handler) filter(w http.ResponseWriter, r *http.Request) []structs.Post {
	if r.Method != http.MethodPost {
		h.logError(w, r, errors.New("Wrong method"), http.StatusMethodNotAllowed)
		return nil
	}
	java := ""
	python := ""
	kotlin := ""
	topic := ""
	java = r.Form.Get("postTopicJava")
	python = r.Form.Get("postTopicPython")
	kotlin = r.Form.Get("postTopicKotlin")
	topic = r.Form.Get("postTopicInput")

	result, err := h.Service.Filter.Filter(java, kotlin, python, topic)
	if err != nil {
		h.logError(w, r, errors.New("Wrong method"), http.StatusBadRequest)
		return nil
	}

	return result
}
