package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) getUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.RequestURI)

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 32)
	user := FetchUser(uint(id), a.DB)
	respondWithJSON(w, http.StatusOK, user)
}

func (a *App) createUser(w http.ResponseWriter, r *http.Request) {
	var user User

	log.Printf("%s %s", r.Method, r.RequestURI)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Payload")
		return
	}
	user.SetBaseRank(a.DB)
	if user.Valid() {
		user.Save(a.DB)
	} else {
		respondWithError(w, http.StatusBadRequest, "Invalid Params")
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}

func (a *App) createRide(w http.ResponseWriter, r *http.Request) {
	var ride Ride
	var user User

	log.Printf("%s %s", r.Method, r.RequestURI)

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 32)
	user = FetchUser(uint(id), a.DB)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&ride)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Payload")
		return
	}
	ride.User = user
	if ride.Valid(a.DB) {
		ride.Save(a.DB)
		ride.User.UpdateLoyaltyRank(a.DB)
		ride.User.UpdateLoyaltyPoint(ride, a.DB)
	} else {
		respondWithError(w, http.StatusBadRequest, "Invalid Params")
		return
	}
	respondWithJSON(w, http.StatusOK, ride)
}
