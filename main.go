package main

import (
	"fiberWebApi/database"
	"fiberWebApi/models"
	"fiberWebApi/routes"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
	"log"
)

const SecretKey = "secret"

func welcome(c *fiber.Ctx) error {
	return c.SendString("start of the API welcome function return")
}

func welcomeHome(c *fiber.Ctx) error {
	//  templating here

	fmt.Println(c.GetReqHeaders())

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

func account(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	// user isnt logged in
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var account models.Account

	database.Database.Db.Where("id = ?", claims.Issuer).First(&account)

	return c.Render("account", fiber.Map{
		"Item":      "this is the 'account' page ;) ",
		"ID":        account.ID,
		"Email":     account.Email,
		"CreatedAt": account.CreatedAt,
		"Name":      account.Name,
	})
}

func setupRoutes(app *fiber.App) {
	// welcome endpoint
	app.Get("/api", welcome)
	app.Get("/", welcomeHome)
	app.Get("/login", login)
	app.Get("/account", account)
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
	app.Post("/api/register", routes.CreateAccount) // store creds in the database
	app.Post("/api/login", routes.GetLogin)         // checks the creds against the stored db creds

	app.Get("/api/account", routes.GetAccount) // gets the current logged in user with the cookie
	app.Get("/api/logout", routes.Logout)      // removes the cookie

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
