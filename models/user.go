package models

import "time"

type User struct {
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	Password  string    `json:"password" bson:"password"`
	Number    string    `json:"number" bson:"number"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
