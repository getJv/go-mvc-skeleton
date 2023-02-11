package routes

import (
	"getjv/backend/app/http/controller"

	"github.com/gofiber/fiber/v2"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for POST method:
	//route.Get("/users", middleware.JWTProtected(), controllers.UserIndex) // retrieve all users
	route.Get("/users",  controller.UserIndex) // retrieve all users
	route.Post("/users", controller.UserStore) // save a users
	route.Get("/users/:id", controller.UserShow) // retrieve a users
	route.Put("/users/:id", controller.UserUpdate) // retrieve a users
	route.Delete("/users/:id", controller.UserDelete) // retrieve a users
	//a.Get("/users", controllers.UserIndex) // retrieve all users
}