package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/henny129/nexcraft-gocore-starter-kit/config"
	"github.com/henny129/nexcraft-gocore-starter-kit/database"
	"github.com/henny129/nexcraft-gocore-starter-kit/routes"
)

func main() {
	config.LoadEnv()
	database.ConnectDB()

	app := fiber.New()

	routes.RegisterRoutes(app)

	app.Listen(":3000")
}
