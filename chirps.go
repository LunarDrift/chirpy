package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type chirpRequest struct {
	Body string `json:"body"`
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

	cleanedBody := filterProfanity(req.Body)

	respondWithJSON(w, http.StatusOK, map[string]string{"cleaned_body": cleanedBody})
}

func filterProfanity(body string) string {
	profanity := map[string]bool{
		"kerfuffle": true,
		"sharbert":  true,
		"fornax":    true,
	}

	words := strings.Split(body, " ")
	for i, word := range words {
		if profanity[strings.ToLower(word)] {
			words[i] = "****"
		}
	}
	return strings.Join(words, " ")
}
