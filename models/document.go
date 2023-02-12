package models

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	FileData string `json:"image"`
	Text     string
}
