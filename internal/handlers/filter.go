package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"forum/structs"
)

type filters struct {
	Topics []string `json:"topics"`
}

type Filtered struct {
	Result []structs.Post `json:"result"`
}

func (h *Handler) filter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.logError(w, r, errors.New("Wrong method"), http.StatusMethodNotAllowed)
		return
	}
	var filters filters
	err := json.NewDecoder(r.Body).Decode(&filters)
	if err != nil {
		h.logError(w, r, err, http.StatusBadRequest)
		return
	}
	fmt.Println(filters.Topics)
	filtered, err := h.Service.Filter.Filter(filters.Topics)

	err = json.NewEncoder(w).Encode(&filtered)
	if err != nil {
		return
	}
	
}
