package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Identifier string
	FirstName string
	LastName  string
}
