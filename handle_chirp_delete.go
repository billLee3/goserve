package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/goserve/internal/auth"
)

func (cfg *apiConfig) handlerChirpsDelete(w http.ResponseWriter, r *http.Request) {
	type response struct {
		User
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, 401, "Not Authorized")
		return
	}
	userID, err := auth.ValidateJWT(token, cfg.secret)
	if err != nil {
		respondWithError(w, 401, "Not Authorized")
		return
	}

	chirpIDString := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(chirpIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID")
		return
	}

	chirp, err := cfg.db.GetChirpById(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, 404, "chirp not found")
		return
	}

	if chirp.UserID != userID {
		respondWithError(w, 403, "Access denied")
		return
	}

	err = cfg.db.DeleteChripById(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, 404, "Chirp not found")
		return
	}

	respondWithJSON(w, 204, response{
		User{
			ID: userID,
		},
	})
}
