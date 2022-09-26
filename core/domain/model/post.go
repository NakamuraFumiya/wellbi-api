package model

import (
	"time"
)

type Post struct {
	ID        uint
	To        string
	From      string
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
