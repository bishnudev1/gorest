package routes

import (
	"gorest/controllers"
	"gorest/middlewars"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/api/v1/users", controllers.GetUsers)
	app.Get("/api/v1/working", controllers.Working)
	app.Post("/api/v1/create-user", controllers.CreateUsers)
	app.Post("/api/v1/login-user", controllers.LoginUser)
	app.Get("/api/v1/me", middlewars.Auth, controllers.GetUser)
	app.Put("/api/v1/update-user", middlewars.Auth, controllers.UpdateUsers)
	app.Delete("/api/v1/delete-user", middlewars.Auth, controllers.DeleteUsers)
	app.Get("/api/v1/logout", controllers.Logout)
}
