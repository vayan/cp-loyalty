package main

import (
	"github.com/jinzhu/gorm"
)

// LoyaltyRank model
type LoyaltyRank struct {
	gorm.Model
	Name               string `json:"name"`
	RequiredRidesCount int    `json:"required_rides_count"`
	Multiplier         int    `json:"multiplier"`
}

// Valid validate the model
func (u *LoyaltyRank) Valid() bool {
	return u.RequiredRidesCount >= 0 &&
		len(u.Name) > 0 &&
		u.Multiplier > 0
}

func fetchLoyaltyRank(id int, db *gorm.DB) LoyaltyRank {
	var LoyaltyRank LoyaltyRank

	db.First(&LoyaltyRank, id)
	return LoyaltyRank
}

//Save to persit the LoyaltyRank in the DB
func (u *LoyaltyRank) Save(db *gorm.DB) {
	if db.NewRecord(u) {
		db.Create(&u)
	} else {
		db.Save(&u)
	}
}
