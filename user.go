package main

import (
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	LoyaltyPoint  int `json:"loyalty_point"`
	LoyaltyRank   LoyaltyRank
	LoyaltyRankID uint
	Rides         []Ride
	RidesLeft     int
}

// Valid validate the model
func (u *User) Valid() bool {
	return u.LoyaltyPoint >= 0
}

// FetchUser find a user in the db using his ID
func FetchUser(id uint, db *gorm.DB) User {
	var user User

	db.Preload("LoyaltyRank").Preload("Rides").First(&user, id)

	user.RidesLeft = user.RidesLeftBeforeNextRank(db)

	return user
}

//Save to persit the user in the DB
func (u *User) Save(db *gorm.DB) {
	if db.NewRecord(u) {
		db.Create(&u)
	} else {
		db.Save(&u)
	}
}

//SetBaseRank sets the bronze loyalty rank to the user
func (u *User) SetBaseRank(db *gorm.DB) {
	var baseRank LoyaltyRank

	db.Where(&LoyaltyRank{Name: "bronze"}).First(&baseRank)

	u.LoyaltyRank = baseRank
}

//UpdateLoyaltyRank checks if we need to update the user loyalty rank
func (u *User) UpdateLoyaltyRank(db *gorm.DB) {
	var newRank LoyaltyRank

	*u = FetchUser(u.ID, db)

	db.Where("required_rides_count <= ?", len(u.Rides)).
		Order("required_rides_count desc").
		First(&newRank)

	if newRank != u.LoyaltyRank {
		u.LoyaltyRank = newRank
		u.Save(db)
	}
}

//UpdateLoyaltyPoint checks if we need to update the user loyalty points
func (u *User) UpdateLoyaltyPoint(ride Ride, db *gorm.DB) {
	u.LoyaltyPoint += (ride.Price * u.LoyaltyRank.Multiplier)
	u.Save(db)
}

//RidesLeftBeforeNextRank counts
func (u *User) RidesLeftBeforeNextRank(db *gorm.DB) int {
	var nextRank LoyaltyRank

	RidesCount := len(u.Rides)

	db.Where("required_rides_count > ?", RidesCount).
		Order("required_rides_count asc").
		First(&nextRank)

	return nextRank.RequiredRidesCount - RidesCount
}
