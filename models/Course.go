package models

import (
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	Name           string `gorm:"not null"`
	CourseLevel    string `gorm:"not null"`
	CourseStatus   string `gorm:"not null"`
	Description    string `gorm:"not null"`
	Image_url      string
	UserInstructor uint
	Modules        []Module `gorm:"foreignKey:ID"`
}
