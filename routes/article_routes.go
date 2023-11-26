package routes

import (
	"gorest/controllers"
	"gorest/middlewars"

	"github.com/gofiber/fiber/v2"
)

func ArticleRoutes(app *fiber.App) {
	app.Post("/api/v1/post-articles", middlewars.Auth, controllers.CreateArticle)
	app.Get("/api/v1/get-articles", controllers.GetArticles)
	app.Get("/api/v1/get-article/:id", controllers.GetArticle)
	app.Delete("/api/v1/delete-article/:id", middlewars.Auth, controllers.DeleteArticle)
	app.Put("/api/v1/update-article", middlewars.Auth, controllers.UpdateArticle)
}
