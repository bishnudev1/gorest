package routes

import (
	"gorest/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/api/v1/users", controllers.GetUsers)
	app.Get("/api/v1/working", controllers.Working)
	app.Post("/api/v1/create-user", controllers.CreateUsers)
	app.Put("/api/v1/update-user", controllers.UpdateUsers)
	app.Delete("/api/v1/delete-user", controllers.DeleteUsers)
}
