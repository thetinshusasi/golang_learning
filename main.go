package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thetinshusasi/golang_learning/database"
	"github.com/thetinshusasi/golang_learning/routes"
)

func setupRoutes(app *fiber.App) {

	app.Post("/api/user", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋! 👋!")
	})

}

func main() {
	app := fiber.New()
	database.ConnectDb()
	setupRoutes(app)

	app.Listen(":3000")
}
