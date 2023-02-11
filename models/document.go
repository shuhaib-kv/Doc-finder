package models

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	FatherName string `gorm:"type:varchar(100);not null"`
	FilePath   string `gorm:"type:varchar(100);not null"`
}
