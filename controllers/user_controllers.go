package controllers

import (
	"context"
	"fmt"
	"gorest/config"

	"gorest/models"
	"gorest/responses"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

//var usersCollection *mongo.Collection = config.MI.DB.Collection("users")

var validate = validator.New()

func CreateUsers(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	usersCollection := config.MI.DB.Collection("users")

	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).JSON(responses.UserResponse{Status: 500, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(422).JSON(responses.UserResponse{Status: 422, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newUsers := models.User{
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: time.Now(),
	}

	fmt.Println("New Users", newUsers)

	var existingUsers models.User

	err := usersCollection.FindOne(ctx, bson.M{"email": newUsers.Email}).Decode(&existingUsers)

	if err == nil {
		fmt.Println("Error Finding Users", err)
		return c.Status(500).JSON(responses.UserResponse{Status: 500, Message: "Users Already Exists", Data: &fiber.Map{"data": err.Error()}})
	}
	result, err := usersCollection.InsertOne(ctx, newUsers)

	if err != nil {
		return c.Status(500).JSON(responses.UserResponse{Status: 500, Message: "Failed to Create User", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(200).JSON(responses.UserResponse{Status: 200, Message: "Users Account has Created Successfully", Data: &fiber.Map{"data": result}})
}

func UpdateUsers(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	usersCollection := config.MI.DB.Collection("users")

	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).JSON(responses.UserResponse{Status: 500, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	fmt.Println("Current User", user)

	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(422).JSON(responses.UserResponse{Status: 422, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	fmt.Println("Current User", user)

	var oldUser models.User

	err := usersCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&oldUser)

	if err != nil {
		return c.Status(500).JSON(responses.UserResponse{Status: 500, Message: "User Not Found", Data: &fiber.Map{"data": err}})
	}

	fmt.Println("Current User", user)

	result, err := usersCollection.UpdateOne(ctx, bson.M{"email": user.Email}, bson.M{"$set": bson.M{"name": user.Name, "password": user.Password}})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "Internal Server Error",
			"data":    err,
		})
	}

	fmt.Println("Result", result.ModifiedCount)

	return c.Status(200).JSON(responses.UserResponse{Status: 200, Message: "success", Data: &fiber.Map{"data": "User Updated Successfully"}})
}

func DeleteUsers(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	usersCollection := config.MI.DB.Collection("users")

	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).JSON(responses.UserResponse{Status: 500, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(422).JSON(responses.UserResponse{Status: 422, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}
	err := usersCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&user)

	if err != nil {
		return c.Status(500).JSON(responses.UserResponse{Status: 500, Message: "User Not Found", Data: &fiber.Map{"data": err.Error()}})
	}

	results, err := usersCollection.DeleteOne(ctx, bson.M{"email": user.Email})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "Internal Server Error",
			"data":    nil,
		})
	}

	fmt.Println(results.DeletedCount)

	return c.Status(200).JSON(responses.UserResponse{Status: 200, Message: "success", Data: &fiber.Map{"data": "User Deleted Successfully"}})
}

func GetUsers(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	usersCollection := config.MI.DB.Collection("users")

	var Userss []models.User

	results, err := usersCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "Internal Server Error",
			"data":    nil,
		})
	}

	defer results.Close(ctx)

	for results.Next(ctx) {
		var singleUsers models.User
		if err = results.Decode(&singleUsers); err != nil {
			return c.Status(500).JSON(responses.UserResponse{Status: 500, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		Userss = append(Userss, singleUsers)
	}

	fmt.Println(Userss)

	return c.Status(200).JSON(
		responses.UserResponse{Status: 200, Message: "success", Data: &fiber.Map{"data": Userss}},
	)
}

func Working(c *fiber.Ctx) error {
	return c.SendString("External Routes are Working...")
}

// func DeleteCollection(c *fiber.Ctx) error {

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

// 	defer cancel()

// 	err := usersCollection.Drop(ctx)

// 	if err != nil {
// 		return c.Status(500).JSON(responses.UserResponse{Status: 500, Message: "error", Data: &fiber.Map{"data": err.Error()}})
// 	}

// 	return c.Status(200).JSON(responses.UserResponse{Status: 200, Message: "success", Data: &fiber.Map{"data": "Collection Deleted Successfully"}})
// }
