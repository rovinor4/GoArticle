package route

import (
	controllers "GoArticle/app/controller"

	"github.com/gofiber/fiber/v2"
)

func Api(app *fiber.App) {

	app.Get("/", controllers.GetAPI)

	auth := app.Group("/auth")
	auth.Post("/sign-up", controllers.Register)
	auth.Post("/login", controllers.Login)
}
