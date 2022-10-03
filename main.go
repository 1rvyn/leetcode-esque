package main

import (
	"fiberWebApi/database"
	"fiberWebApi/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
	"log"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("start of the API welcome function return")
}

func welcomeHome(c *fiber.Ctx) error {
	//  templating here
	return c.Render("index", fiber.Map{
		"Title": "Hello World!",
	})
}

func login(c *fiber.Ctx) error {
	return c.Render("nice", fiber.Map{
		"User":   "Irvyn Hall",
		"Email":  "irvynhall@gmail.com",
		"Status": "logging-in",
	})
}

func setupRoutes(app *fiber.App) {
	// welcome endpoint
	app.Get("/api", welcome)
	app.Get("/", welcomeHome)
	app.Get("/login", login)
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
	app.Post("/api/register", routes.CreateAccount)
	app.Post("/api/login", routes.GetLogin)

	app.Get("/api/account", routes.GetAccount)
	app.Get("/api/logout", routes.Logout)

}

func main() {
	database.ConnectDb()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Static("/", "./views/public")

	setupRoutes(app)

	log.Fatal(app.Listen("127.0.0.1:3000"))
}
