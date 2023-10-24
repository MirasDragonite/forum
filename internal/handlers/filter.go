package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"forum/structs"
)

func (h *Handler) filter(w http.ResponseWriter, r *http.Request) []structs.Post {
	if r.Method != http.MethodPost {
		h.logError(w, r, errors.New("Wrong method"), http.StatusMethodNotAllowed)
		return nil
	}

	java := r.Form.Get("postTopicJava")
	python := r.Form.Get("postTopicPython")
	kotlin := r.Form.Get("postTopicKotlin")
	topic := r.Form.Get("postTopicInput")

	fmt.Println(java, python, kotlin, topic)
	topics := make([]string, 0)
	if java == "Java" {
		topics = append(topics, java)
	}
	if python == "Python" {
		topics = append(topics, python)
	}
	if kotlin == "Kotlin" {
		topics = append(topics, kotlin)
	}
	if strings.TrimSpace(topic) != "" {
		topics = append(topics, topic)
	}

	fmt.Println(topics)
	result, err := h.Service.Filter.Filter(topics)
	if err != nil {
		h.logError(w, r, errors.New("Wrong method"), http.StatusBadRequest)
		return nil
	}

	return result
	// http.Redirect(w, r, "/", http.StatusSeeOther)
}
