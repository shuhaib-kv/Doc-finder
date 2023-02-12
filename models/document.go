package models

import (
	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	FileData string `json:"image"`
	Text     string
}
type PDF struct {
	ID      int
	File    []byte
	Content string
	Text    string
}
