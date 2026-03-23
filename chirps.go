package main

import (
	"encoding/json"
	"net/http"
)

type chirpRequest struct {
	Body string `json:"body"`
}

type chirpResponse struct {
	Valid bool   `json:"valid"`
	Error string `json:"error,omitempty"`
}

func (cfg *apiConfig) validateChirpHandler(w http.ResponseWriter, r *http.Request) {
	var req chirpRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	if len(req.Body) > 140 {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return
	}

	respondWithJSON(w, http.StatusOK, chirpResponse{Valid: true})
}
