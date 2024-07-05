package handler

import (
	"backend/internal/app/schema"
	"encoding/json"
	"net/http"
)

// RegisterHandler handles user registration
func (h *Handler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req schema.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Validate(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// TODO Handle registration logic here
	w.WriteHeader(http.StatusCreated)
}

// LoginHandler handles user login
func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req schema.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Validate(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// TODO Handle login logic here, generate and return JWT
	w.WriteHeader(http.StatusOK)
}

// PasswordRecoveryHandler handles password recovery request
func (h *Handler) PasswordRecoveryHandler(w http.ResponseWriter, r *http.Request) {
	var req schema.PasswordRecoveryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Validate(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// TODO Handle password recovery logic here
	w.WriteHeader(http.StatusOK)
}

// PasswordResetHandler handles password reset
func (h *Handler) PasswordResetHandler(w http.ResponseWriter, r *http.Request) {
	var req schema.PasswordResetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Validate(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// TODO Handle password reset logic here
	w.WriteHeader(http.StatusOK)
}
