package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	ID    string `json:"ID"`
	Title string `json:"FirstName"`
	Body  string `json:"LastName"`
	Email string `json:"Email"`
	Scope string `json:"Scope"`
}
