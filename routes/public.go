package routes

import (
	"getjv/backend/app/http/controller"
	"os"

	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString(os.Getenv("WELCOME_MSG"))
    })

	// Create routes group.
	route := app.Group("/api")

	route.Post("/signup", controller.AuthSignUp)
	route.Post("/login", controller.AuthLogin)
	
}