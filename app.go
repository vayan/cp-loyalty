package main

import (
	"net/http"

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
	a.setupDB()
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

// Run start the app
func (a *App) Run(addr string) {
	log.Printf("App started on port %s", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/users/{id:[0-9]+}/rides", a.createRide).Methods("POST")
	a.Router.HandleFunc("/users/{id:[0-9]+}", a.getUser).Methods("GET")
	a.Router.HandleFunc("/users", a.createUser).Methods("POST")
}

func (a *App) setupDB() {
	a.DB.DropTable(&LoyaltyRank{})

	a.DB.AutoMigrate(&User{})
	a.DB.AutoMigrate(&LoyaltyRank{})
	a.DB.AutoMigrate(&Ride{})

	a.DB.Create(&LoyaltyRank{Name: "bronze", RequiredRidesCount: 0, Multiplier: 1})
	a.DB.Create(&LoyaltyRank{Name: "silver", RequiredRidesCount: 5, Multiplier: 3})
	a.DB.Create(&LoyaltyRank{Name: "gold", RequiredRidesCount: 15, Multiplier: 5})
	a.DB.Create(&LoyaltyRank{Name: "platinum", RequiredRidesCount: 30, Multiplier: 10})
}
