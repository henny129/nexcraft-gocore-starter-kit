package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/henny129/nexcraft-gocore-starter-kit/controllers"
	"github.com/henny129/nexcraft-gocore-starter-kit/internal/middlewares"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Public routes
	api.Post("/login", controllers.Login)
	api.Post("/register", controllers.CreateUser)

	// Protected routes
	user := api.Group("/users", middlewares.JWTMiddleware())
	user.Get("/", controllers.GetUsers)
	user.Post("/", controllers.CreateUser)
	user.Get("/:id", controllers.GetUser)
	user.Put("/:id", controllers.UpdateUser)
	user.Delete("/:id", controllers.DeleteUser)
}
