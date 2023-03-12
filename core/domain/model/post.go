package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title    string
	Message  string
	ImageURL *string
	To       string
	From     string
}
