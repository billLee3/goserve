package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/goserve/internal/auth"
)

func (cfg *apiConfig) handlerPolkaWebhook(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserID string `json:"user_id"`
		} `json:"data"`
	}

	type response struct {
		Body string
	}

	api_key, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 401, fmt.Sprintf("%v", err))
		return
	}

	if api_key != cfg.polka_key {
		respondWithError(w, 401, "Unauthorized: api key doesn't match")
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	if params.Event != "user.upgraded" {
		respondWithError(w, 204, "nothing to do")
		return
	}

	userId, err := uuid.Parse(params.Data.UserID)
	if err != nil {
		respondWithError(w, 400, "invalid userid")
		return
	}

	_, err = cfg.db.UpgradeUserToChirpyRed(r.Context(), userId)
	if err != nil {
		respondWithError(w, 404, "unable to find user")
		return
	}

	respondWithJSON(w, 204, response{})

}
