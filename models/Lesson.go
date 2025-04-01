package models

import "gorm.io/gorm"

type Lesson struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	VideoUrl    string `gorm:"not null"`
	Module      Module `gorm:"not null"`
}
