package model

import (
	"gorm.io/gorm"
)

type Roadmap struct {
	gorm.Model
	Title    string
	Message  string
	ImageURL *string
	To       string
	From     string
}
