package models

import "gorm.io/gorm"

type Lesson struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	VideoUrl    string `gorm:"not null"`
	ModuleID    uint
	Module      Module `gorm:"foreignKey:ModuleID;constraint:OnDelete:CASCADE;"`
}
