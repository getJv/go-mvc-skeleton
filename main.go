package main

import (
	"getjv/backend/app/http/middleware"
	"getjv/backend/kernel"
	"getjv/backend/routes"

	"github.com/gofiber/fiber/v2"

	_ "getjv/backend/docs"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)


func init(){
    kernel.DB =  kernel.DatabaseConnect()
    kernel.LoadMigrations();
}


// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
    // Define Fiber config.
	config := kernel.FiberConfig()
	// Define a new Fiber app with config.
	app := fiber.New(config)

    // Middlewares.
	middleware.Cors(app) // Register Cors middleware for app.
   
    // Routes.
	routes.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
	routes.PublicRoutes(app)  // Register a public routes for app.
	routes.PrivateRoutes(app) // Register a private routes for app.
	routes.NotFoundRoute(app) // Register route for 404 Error.
    
    

    kernel.StartApp(app);
}