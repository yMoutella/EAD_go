package models

import "gorm.io/gorm"

type Module struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	CourseID    uint
	Course      Course `gorm:"foreignKey:CourseID;constraint:OnDelete:CASCADE;"`
	Lessons     []Lesson
}
