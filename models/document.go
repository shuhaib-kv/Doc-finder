package models

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	Image string `json:"image"`
}
