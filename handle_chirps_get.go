package main

import (
	"net/http"

	"github.com/google/uuid"
)

// Get all chirps
func (cfg *apiConfig) handleGetChirps(w http.ResponseWriter, r *http.Request) {
	chirps, err := cfg.db.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, 500, "unable to get chirps")
		return
	}

	var ChirpsStruct []Chirp

	for _, chirp := range chirps {
		var NewChirp Chirp
		NewChirp.Id = chirp.ID
		NewChirp.Created_At = chirp.CreatedAt
		NewChirp.Updated_At = chirp.UpdatedAt
		NewChirp.Body = chirp.Body
		NewChirp.User_id = chirp.UserID
		ChirpsStruct = append(ChirpsStruct, NewChirp)
	}

	respondWithJSON(w, 200, ChirpsStruct)
}

// Get single chirp
func (cfg *apiConfig) handleGetChirpById(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("chirpID"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid chirp ID")
		return
	}
	chirp, err := cfg.db.GetChirpById(r.Context(), id)
	if err != nil {
		respondWithError(w, 404, "no chirp with that id")
		return
	}

	respondWithJSON(w, 200, Chirp{
		Id:         chirp.ID,
		Created_At: chirp.CreatedAt,
		Updated_At: chirp.UpdatedAt,
		Body:       chirp.Body,
		User_id:    chirp.UserID,
	})
}
