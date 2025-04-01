package models

import "gorm.io/gorm"

type Module struct {
	gorm.Model
	Title       string   `gorm:"not null"`
	Description string   `gorm:"not null"`
	Course      Course   `gorm:"not null;foreignKey:ID"`
	Lessons     []Lesson `gorm:"not null;foreignKey:ID"`
}
