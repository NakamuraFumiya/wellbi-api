package model

import (
	"gorm.io/gorm"
	"time"
)

type Roadmap struct {
	//gorm.Model
	ID        uint `gorm:"primarykey"`
	Title     string
	Message   string
	ImageURL  *string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
