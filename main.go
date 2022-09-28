package main

import (
	"fiberWebApi/database"
	"fiberWebApi/routes"
	"github.com/gofiber/fiber/v2"
	"log"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("start of the API welcome function return")
}

func setupRoutes(app *fiber.App) {
	// welcome endpoint
	app.Get("/api", welcome)

	// user endpoints
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)

	// product endpoints
	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/products/:id", routes.GetProduct)
	app.Put("/api/products/:id", routes.UpdateProduct)
	app.Delete("/api/products/:id", routes.DeleteProduct)

	// order endpoints
	app.Post("/api/orders", routes.CreateOrder)
	app.Get("/api/orders", routes.GetOrders)
	app.Get("/api/orders/:id", routes.GetOrder)

}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
