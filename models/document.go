package models

import (
	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	FileData    string `json:"image"`
	Textinimage string
}
type PDF struct {
	ID      int
	File    []byte
	Content string
	Text    string
}
