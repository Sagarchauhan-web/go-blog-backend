package main

import "net/http"

func (apiConfig *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `name`
	}
	respondWithJSON(w, 200, struct{}{})
}