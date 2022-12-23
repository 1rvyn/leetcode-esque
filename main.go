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

type Page struct {
	Title string
	URL   string
}

// array of pages to show in the header if they are logged out
var pages = []Page{
	{Title: "Home", URL: "/"},
	{Title: "Login", URL: "/login"},
	{Title: "Register", URL: "/register"},
	{Title: "Account", URL: "/account"},
}

// array of pages to show in the header if they are logged in
var pages2 = []Page{
	{Title: "Home", URL: "/"},
	{Title: "Account", URL: "/account"},
	{Title: "Logout", URL: "/api/logout"},
}

func welcomeHome(c *fiber.Ctx) error {
	//  templating here

	//fmt.Println(c.GetReqHeaders())
	cookie := c.Cookies("jwt")

	activeURL := c.Path()

	if cookie == "" {
		return c.Render("index", fiber.Map{
			"Title":     "Home page WITHOUT cookie",
			"Pages":     pages,
			"ActiveURL": activeURL,
		})
	} else {
		return c.Render("index", fiber.Map{
			"Title":     "Home page WITH cookie",
			"Pages":     pages2,
			"ActiveURL": activeURL,
		})
	}
}

func login(c *fiber.Ctx) error {
	// do some kind of check to see if the user is already logged in
	//- similar to how we do it with /account

	activeURL := c.Path()

	return c.Render("nice", fiber.Map{
		"Pages":     pages,
		"ActiveURL": activeURL,
		"User":      "Irvyn Hall",
		"Email":     "irvynhall@gmail.com",
		"Status":    "logging-in",
	})
}

func accountHandle(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	// user isnt logged in
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.Redirect("/login")
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var account models.Account
	var session []models.Session

	fmt.Println("claims issuer: ", claims.Issuer)

	database.Database.Db.Where("id = ?", claims.Issuer).First(&account)
	database.Database.Db.Where("email = ?", account.Email).Find(&session)

	fmt.Println("sessions found are: ", session)

	//user := c.UserContext()

	activeURL := c.Path()

	currCookie := c.Cookies("jwt")

	if currCookie != "" {
		// there is a cookie so lets show the header pages for logged in users
		// & show the current account data :)
		return c.Render("account", fiber.Map{
			"Pages":     pages2,
			"ActiveURL": activeURL,
			"Item":      "this is the 'account' page ;) ",
			"ID":        account.ID,
			"Email":     account.Email,
			"CreatedAt": account.CreatedAt,
			"Name":      account.Name,
			"Session":   session,
		})
	} else {
		// there is no cookie lets redirect to the login page
		return c.Redirect("/login")
	}
}

func register(c *fiber.Ctx) error {

	activeURL := c.Path()

	return c.Render("register", fiber.Map{
		"Pages":     pages,
		"ActiveURL": activeURL,
		"Register":  "this is the register page",
	})
}
func setupRoutes(app *fiber.App) {
	// welcome endpoint
	app.Get("/api", welcome)
	app.Get("/", welcomeHome)
	app.Get("/login", login)
	app.Get("/account", accountHandle)
	app.Get("/register", register)
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
		Views: engine,
		//ViewsLayout: "layouts/layout",
	})

	// prevent the app from using cache / caching my main .css file

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Set("Pragma", "no-cache")
		c.Set("Expires", "0")
		return c.Next()
	})

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Static("/", "./views/public")

	setupRoutes(app)

	log.Fatal(app.Listen("127.0.0.1:3000"))
}
