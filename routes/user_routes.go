package routes

import (
	"gorest/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/api/v1/users", controllers.GetUsers)
	app.Get("/api/v1/working", controllers.Working)
}
