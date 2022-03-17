package models

import "gorm.io/gorm"

type Phone struct {
	gorm.Model
	Phonetype   string `json:"type"`
	PhoneNumber string `json:"phone_number"`
	ContactID   uint   `json:"contact_id"`
}

type Address struct {
	gorm.Model
	Country   string `json:"country"`
	City      string `json:"city"`
	Street    string `json:"street"`
	ContactID string `json:"contact_id"`
}

type Contact struct {
	gorm.Model
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Phone     []Phone `json:"phone"`
	Address   Address `json:"address"`
	UserID    uint    `json:"user_id"`
}
