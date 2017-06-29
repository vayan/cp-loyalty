package main

import (
	"github.com/jinzhu/gorm"
)

// Ride model
type Ride struct {
	gorm.Model
	Price  int `json:"price"`
	User   User
	UserID uint
}

// Valid validate the model
func (u *Ride) Valid(db *gorm.DB) bool {
	return u.Price > 0 &&
		!db.NewRecord(u.User)
}

func fetchRide(id int, db *gorm.DB) Ride {
	var Ride Ride
	db.First(&Ride, id)
	return Ride
}

//Save to persit the Ride in the DB
func (u *Ride) Save(db *gorm.DB) {
	if db.NewRecord(u) {
		db.Create(&u)
	} else {
		db.Save(&u)
	}
}
