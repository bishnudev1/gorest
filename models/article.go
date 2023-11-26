package models

import (
	"time"
)

type Article struct {
	ID          int       `json:"id" bson:"id"`
	Title       string    `json:"title" bson:"title" validate:"required"`
	Description string    `json:"description" bson:"description" validate:"required"`
	Author      string    `json:"author" bson:"author"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
}
