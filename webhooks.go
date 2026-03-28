package main

import (
	"encoding/json"
	"net/http"

	"github.com/LunarDrift/chirpy/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handleUpgradeUserStatus(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil || apiKey != cfg.polkaKey {
		respondWithError(w, http.StatusUnauthorized, "Missing or invalid api key")
		return
	}

	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserID uuid.UUID `json:"user_id"`
		} `json:"data"`
	}
	var params parameters
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't decode parameters")
		return
	}

	if params.Event != "user.upgraded" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	result, err := cfg.db.UpgradeUserToChirpyRed(r.Context(), params.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't update user")
		return
	}

	// make sure we actually updated a user
	rows, err := result.RowsAffected()
	if err != nil || rows == 0 {
		respondWithError(w, http.StatusNotFound, "Couldn't find user")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
