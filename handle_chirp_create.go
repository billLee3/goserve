package main

import (
	"encoding/json"
	"io"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Chirp struct {
	Id         uuid.UUID `json:"id"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
	Body       string    `json:"body"`
	User_id    uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) handleChirpsCreate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// Type for responseBody
	type requestBody struct {
		Body string `json:"body"`
	}
	type responseBody struct {
		Cleaned_body string `json:"cleaned_body"`
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, 500, "couldn't read the request")
		return
	}
	MyReq := requestBody{}
	err = json.Unmarshal(data, &MyReq)
	if err != nil {
		respondWithError(w, 500, "Something went wrong")
		return
	}

	if len(MyReq.Body) > 140 {
		respondWithError(w, 400, "Chirp is too long")
		return
	}

	const stars = "****"
	words := strings.Split(MyReq.Body, " ")
	profane_words := []string{"kerfuffle", "sharbert", "fornax"}

	for index, word := range words {
		if slices.Contains(profane_words, strings.ToLower(word)) {
			words[index] = stars
		}
	}

	cleaned_sentence := strings.Join(words, " ")

	respondWithJSON(w, 200, responseBody{
		Cleaned_body: cleaned_sentence,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) error {
	resp, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(resp)
	return nil
}

func respondWithError(w http.ResponseWriter, code int, msg string) error {
	return respondWithJSON(w, code, map[string]string{"error": msg})
}
