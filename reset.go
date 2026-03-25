package main

import "net/http"

func (cfg *apiConfig) resetHandler(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		respondWithError(w, http.StatusForbidden, "Insufficient permissions")
		return
	}

	// delete users with sqlc query
	err := cfg.db.DeleteAllUsers(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete users")
		return
	}

	// reset hit count
	cfg.fileserverHits.Store(0)

	w.WriteHeader(http.StatusOK)
}
