package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"log"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// App will hold our application
type App struct {
	DB     *gorm.DB
	Router *mux.Router
}

// Initialize creates a database connexion
func (a *App) Initialize(dbName string) {
	var err error
	a.DB, err = gorm.Open("sqlite3", dbName)
	if err != nil {
		panic("Failed to connect database")
	}

	a.DB.AutoMigrate(&User{})

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

// Run start the app
func (a *App) Run(addr string) {
	log.Printf("App started on port %s", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// GetUser return a json serialized version of an User
func (a *App) getUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.RequestURI)
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user := fetchUser(id, a.DB)
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
	if user.Valid() {
		user.Save(a.DB)
	} else {
		respondWithError(w, http.StatusBadRequest, "Invalid Params")
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/users/{id:[0-9]+}", a.getUser).Methods("GET")
	a.Router.HandleFunc("/users", a.createUser).Methods("POST")
}
