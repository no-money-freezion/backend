package handler

import (
	"backend/internal/app/schema"
	"encoding/json"
	"net/http"
)

func (h *Handler) HealthHandler(w http.ResponseWriter, req *http.Request) {
	err := json.NewEncoder(w).Encode(schema.HealthcheckResponse{Status: "OK"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}
