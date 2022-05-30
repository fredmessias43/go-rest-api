package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	ID        int    `gorm:"primaryKey"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`

	// hasMany
	Mails []Mail `gorm:"foreignKey:For"`
}

type PersonRequest struct{}
