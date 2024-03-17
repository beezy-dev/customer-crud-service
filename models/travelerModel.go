package models

import "gorm.io/gorm"

type Traveler struct {
	gorm.Model
	Uuid       string
	Firstname  string
	Lastname   string
	Multihop   bool
	Pprocessed bool
	Wprocessed string
	Reprocess  bool
	Offloaded  bool
}
