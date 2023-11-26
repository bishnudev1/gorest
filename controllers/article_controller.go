package controllers

import (
	"context"
	"gorest/config"
	"gorest/models"
	"gorest/responses"
	"gorest/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
)

func GetArticles(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	articleCollection := config.MI.DB.Collection("articles")

	var articles []models.Article

	cursor, err := articleCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(500).JSON(responses.ArticleResponse{Status: 500, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if err := cursor.All(ctx, &articles); err != nil {
		return c.Status(500).JSON(responses.ArticleResponse{Status: 500, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(200).JSON(responses.ArticleResponse{Status: 200, Message: "success", Data: &fiber.Map{"data": articles}})
}

func CreateArticle(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	print(ctx)

	articleCollection := config.MI.DB.Collection("articles")

	var article models.Article

	err := c.BodyParser(&article)

	if err != nil {
		return c.Status(500).JSON(responses.ArticleResponse{Status: 500, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	validationErr := validate.Struct(&article)

	if validationErr != nil {
		return c.Status(422).JSON(responses.ArticleResponse{Status: 422, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	currentUserToken := c.Locals("gorest").(*jwt.Token)

	claims := currentUserToken.Claims.(jwt.MapClaims)

	article.Author = claims["email"].(string)
	article.ID = utils.GenerateRandomID()

	newArticle := models.Article{
		ID:          article.ID,
		Title:       article.Title,
		Description: article.Description,
		Author:      article.Author,
		CreatedAt:   time.Now(),
	}

	result, err := articleCollection.InsertOne(ctx, newArticle)

	if err != nil {
		return c.Status(500).JSON(responses.ArticleResponse{Status: 500, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(201).JSON(responses.ArticleResponse{Status: 201, Message: "success", Data: &fiber.Map{"data": result}})

}

func GetArticle(c *fiber.Ctx) error {
	articleID := c.Params("id")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	id, err := strconv.Atoi(articleID)
	if err != nil {
		return c.Status(400).JSON(responses.ArticleResponse{Status: 400, Message: "error", Data: &fiber.Map{"data": "Invalid article ID"}})
	}

	articleCollection := config.MI.DB.Collection("articles")

	var article models.Article

	err = articleCollection.FindOne(ctx, bson.M{"id": id}).Decode(&article)

	if err != nil {
		return c.Status(500).JSON(responses.ArticleResponse{Status: 500, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(200).JSON(responses.ArticleResponse{Status: 201, Message: "success", Data: &fiber.Map{"data": article}})
}

func UpdateArticle(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	var bodyArticle models.Article

	err := c.BodyParser(&bodyArticle)

	if err != nil {
		return c.Status(500).JSON(responses.ArticleResponse{Status: 500, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	validationErr := validate.Struct(&bodyArticle)

	if validationErr != nil {
		return c.Status(422).JSON(responses.ArticleResponse{Status: 422, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	currentUserToken := c.Locals("gorest").(*jwt.Token)

	claims := currentUserToken.Claims.(jwt.MapClaims)

	authorEmail := claims["email"].(string)

	articleCollection := config.MI.DB.Collection("articles")

	var article models.Article

	err = articleCollection.FindOne(ctx, bson.M{"id": bodyArticle.ID}).Decode(&article)

	if err != nil {
		return c.Status(500).JSON(responses.ArticleResponse{Status: 500, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if authorEmail != article.Author {
		return c.Status(403).JSON(responses.ArticleResponse{Status: 403, Message: "false", Data: &fiber.Map{"data": &fiber.Map{
			"success": false,
			"data":    "You are not the author of this article",
		}}})
	}

	newArticleData := models.Article{
		ID:          article.ID,
		Title:       bodyArticle.Title,
		Description: bodyArticle.Description,
		Author:      authorEmail,
		CreatedAt:   article.CreatedAt,
	}

	result := articleCollection.FindOneAndUpdate(ctx, bson.M{"id": bodyArticle.ID}, bson.M{"$set": newArticleData})

	return c.Status(403).JSON(responses.ArticleResponse{Status: 403, Message: "false", Data: &fiber.Map{"data": result}})
}

func DeleteArticle(c *fiber.Ctx) error {
	articleID := c.Params("id")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	id, err := strconv.Atoi(articleID)
	if err != nil {
		return c.Status(400).JSON(responses.ArticleResponse{Status: 400, Message: "error", Data: &fiber.Map{"data": "Invalid article ID"}})
	}

	currentUserToken := c.Locals("gorest").(*jwt.Token)

	claims := currentUserToken.Claims.(jwt.MapClaims)

	authorEmail := claims["email"].(string)

	articleCollection := config.MI.DB.Collection("articles")

	var article models.Article

	err = articleCollection.FindOne(ctx, bson.M{"id": id}).Decode(&article)

	if err != nil {
		return c.Status(500).JSON(responses.ArticleResponse{Status: 500, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if authorEmail == article.Author {
		print("Same author! Delete request proceed")
		result := articleCollection.FindOneAndDelete(ctx, bson.M{"id": id})
		return c.Status(200).JSON(responses.ArticleResponse{Status: 201, Message: "success", Data: &fiber.Map{"data": result}})
	}

	return c.Status(403).JSON(responses.ArticleResponse{Status: 403, Message: "false", Data: &fiber.Map{"data": &fiber.Map{
		"success": false,
		"data":    "You are not the author of this article",
	}}})
}
