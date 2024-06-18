package handler

import (
	"backend/internal/app/schema"
	"encoding/json"
	"net/http"
)

// RouletteInfoHandler returns roulette info
func (h *Handler) RouletteInfoHandler(w http.ResponseWriter, r *http.Request) {
	// Dummy response for example
	resp := schema.RouletteInfoResponse{
		Version: "1.0",
		Combinations: []schema.RouletteCombination{
			{Combination: "Triple 7", Payout: 100},
		},
		MaxBet: 1000,
		MinBet: 10,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// RouletteBetHandler handles a roulette bet
func (h *Handler) RouletteBetHandler(w http.ResponseWriter, r *http.Request) {
	var req schema.RouletteBetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Validate(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Check rules version
	if req.RulesVersion != "1.0" {
		http.Error(w, "Outdated rules version", http.StatusConflict)
		return
	}
	// Handle bet logic here
	resp := schema.RouletteBetResponse{
		Result:    "Lose",
		WinAmount: 0,
		Message:   "Better luck next time!",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
