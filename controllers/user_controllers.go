package controllers

import (
	// "context"
	"context"
	"fmt"
	"gorest/config"
	"gorest/models"
	"gorest/responses"
	"time"

	// "gorest/models"
	// "gorest/responses"
	// "time"

	"github.com/gofiber/fiber/v2"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = config.GetCollection(config.DB, "igrosine")

func GetUsers(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	var users []models.User

	results, err := userCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "Internal Server Error",
			"data":    nil,
		})
	}

	defer results.Close(ctx)

	for results.Next(ctx) {
		var singleUser models.User
		if err = results.Decode(&singleUser); err != nil {
			return c.Status(500).JSON(responses.UserResponse{Status: 500, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		users = append(users, singleUser)
	}

	fmt.Println(users)

	return c.Status(200).JSON(
		responses.UserResponse{Status: 200, Message: "success", Data: &fiber.Map{"data": users}},
	)
}

func Working(c *fiber.Ctx) error {
	return c.SendString("External Routes are Working...")
}
