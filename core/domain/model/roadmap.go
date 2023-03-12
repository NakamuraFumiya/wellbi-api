package model

import (
	"database/sql"
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
	DeletedAt sql.NullTime `gorm:"index"`
}
