package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	// ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	To      string `json:"to" bson:"to"`
	From    string `json:"from" bson:"from"`
	Message string `json:"message" bson:"message"`
}
