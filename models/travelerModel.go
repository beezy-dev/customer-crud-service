package models

import "gorm.io/gorm"

type Traveler struct {
	gorm.Model
	UniqueIdentifier string
	FirstName        string
	LastName         string
	MultiHop         bool
	Processed        bool
	WhenProcessed    string
	ReProcess        bool
}
