package models

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	Image3 string `json:""`
}
