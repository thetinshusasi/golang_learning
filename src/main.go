package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thetinshusasi/golang_learning/src/database"
	"github.com/thetinshusasi/golang_learning/src/routes"
)

func setupRoutes(app *fiber.App) {

	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Delete("/api/users/:id", routes.DeleteUser)
	// Product endpoints
	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/products/:id", routes.GetProduct)
	app.Put("/api/products/:id", routes.UpdateProduct)
	// Order endpoints
	app.Post("/api/orders", routes.CreateOrder)
	app.Get("/api/orders", routes.GetOrders)
	app.Get("/api/orders/:id", routes.GetOrder)

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
