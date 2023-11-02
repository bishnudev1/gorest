package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Name      string    `json:"name" bson:"name" validate:"required"`
	Email     string    `json:"email" bson:"email" validate:"required"`
	Password  string    `json:"password" bson:"password" validate:"required"`
	Number    string    `json:"number" bson:"number"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

var validate = validator.New()
