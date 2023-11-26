package main

import (
	"fmt"
	"gorest/config"
	"gorest/routes"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	if app != nil {
		fmt.Println("Fiber is running")
		config.ConnectDB()
		routes.Routes(app)
		routes.ArticleRoutes(app)
	} else {
		fmt.Println("Fiber is not running")
	}

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("bishnudevkhutiasecretkey")},
	}))

	app.Get("/", func(c *fiber.Ctx) error {

		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":5000")

}
