package main

import (
	"fmt"
	"gorest/config"
	"gorest/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	if app != nil {
		fmt.Println("Fiber is running")
		config.ConnectDB()
		routes.Routes(app)
	} else {
		fmt.Println("Fiber is not running")
	}

	app.Get("/", func(c *fiber.Ctx) error {

		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":5000")

}
