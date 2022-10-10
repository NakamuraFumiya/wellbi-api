package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	To      string
	From    string
	Message string
}
