package handler

import (
	"encoding/json"
	"net/http"
)

type HealthcheckResponse struct {
	Status string `json:"status"`
}

func (h *Handler) HealthHandler(w http.ResponseWriter, req *http.Request) {
	err := json.NewEncoder(w).Encode(HealthcheckResponse{Status: "OK"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}
