package main

import (
	"fiberWebApi/database"
	"fiberWebApi/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("start of the API welcome function return")
}

func welcomeHome(c *fiber.Ctx) error {
	return c.SendString("this is the home page")
}

func setupRoutes(app *fiber.App) {
	// welcome endpoint
	app.Get("/api", welcome)
	app.Get("/", welcomeHome)
	// same thing for both

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

	// account system
	go app.Post("/api/register", routes.CreateAccount)
	app.Post("/api/login", routes.GetLogin)

	app.Get("/api/account", routes.GetAccount)
	app.Get("/api/logout", routes.Logout)

}

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	setupRoutes(app)

	log.Fatal(app.Listen("127.0.0.1:3000"))
}
