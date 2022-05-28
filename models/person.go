package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	ID        int `gorm:"primaryKey"`
	Firstname string
	Lastname  string
}

type PersonRequest struct {
	Firstname string
	Lastname  string
}
