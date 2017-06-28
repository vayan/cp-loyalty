package main

import (
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	LoyaltyPoint int `json:"loyalty_point"`
}

// Valid validate the model
func (u *User) Valid() bool {
	return u.LoyaltyPoint >= 0
}

func fetchUser(id int, db *gorm.DB) User {
	var user User
	db.First(&user, id)
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
