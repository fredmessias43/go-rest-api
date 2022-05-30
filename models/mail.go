package models

import "gorm.io/gorm"

type Mail struct {
	gorm.Model
	ID      int `gorm:"primaryKey"`
	Content string
	Subject string

	For int
	// belongsTo
	Person Person `gorm:"foreignKey:For"`
}
