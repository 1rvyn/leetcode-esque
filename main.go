package main

import (
	"encoding/json"
	"fiberWebApi/routes"
	"fmt"
	"github.com/go-resty/resty/v2"
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
	{Title: "Problems", URL: "/problems"},
}

// array of pages to show in the header if they are logged in
var pages2 = []Page{
	{Title: "Home", URL: "/"},
	{Title: "Code", URL: "/code/1"},
	{Title: "Problems", URL: "/problems"},
	{Title: "Account", URL: "/account"},
	{Title: "Logout", URL: "/logout"},
}

func welcomeHome(c *fiber.Ctx) error {
	//  templating here

	// print the users cookies

	// fmt.Println(c.GetReqHeaders())
	cookie := c.Cookies("jwt")

	if cookie == "" {
		return c.Render("index", fiber.Map{
			"Title": "Home",
			"Pages": pages,
		})
	} else {
		return c.Render("index", fiber.Map{
			"Title": "Home",
			"Pages": pages2,
		})
	}
}

func login(c *fiber.Ctx) error {
	// do some kind of check to see if the user is already logged in
	//- similar to how we do it with /account

	fmt.Println("\n the ip here is:", c.IP())
	fmt.Println("\n The x-forwarded-for header is: ", c.Get("x-forwarded-for"))

	return c.Render("login", fiber.Map{
		"Title": "Login",
		"Pages": pages,
	})
}

func createQuestion(c *fiber.Ctx) error {
	return c.Render("createQuestion", fiber.Map{
		"Title": "Create Question",
	})
}

func accountHandle(c *fiber.Ctx) error {

	activeURL := c.Path()

	currCookie := c.Cookies("jwt")

	if currCookie != "" {
		// there is a cookie so lets show the header pages for logged-in users
		// & show the current account data :)
		return c.Render("account", fiber.Map{
			"Pages":     pages2,
			"ActiveURL": activeURL,
			"Item":      "this is the 'account' page ;) ",
		})
	} else {
		// there is no cookie lets redirect to the login page
		return c.Redirect("/login")
	}
}

func register(c *fiber.Ctx) error {

	fmt.Println("register was page was accessed :)")

	activeURL := c.Path()

	return c.Render("register", fiber.Map{
		"Pages":     pages,
		"ActiveURL": activeURL,
		"Register":  "this is the register page",
	})
}

func problems(c *fiber.Ctx) error {
	activeURL := c.Path()

	// Fetch the list of questions from the API
	client := resty.New()
	resp, err := client.R().Get("http://api.irvyn.xyz/questions")

	if err != nil {
		return c.Status(500).SendString("Error fetching data from API")
	}

	var questionList []map[string]interface{}
	err = json.Unmarshal(resp.Body(), &questionList)
	if err != nil {
		return c.Status(500).SendString("Error parsing API response")
	}

	currCookie := c.Cookies("jwt")

	if currCookie == "" {
		return c.Render("problems", fiber.Map{
			"Pages":        pages,
			"Title":        "Problems",
			"ActiveURL":    activeURL,
			"QuestionList": questionList,
		})
	} else {
		return c.Render("problems", fiber.Map{
			"Pages":        pages2,
			"Title":        "Problems",
			"ActiveURL":    activeURL,
			"QuestionList": questionList,
		})
	}
}

func setupRoutes(app *fiber.App) {
	// welcome endpoint
	app.Get("/api", welcome)
	app.Get("/", welcomeHome)
	app.Get("/login", login)
	app.Get("/account", accountHandle)
	app.Get("/register", register)
	app.Get("/problems", problems)

	app.Get("/new_question", createQuestion)

	// same thing for both

	// userRole auth routes/pages
	admin := app.Group("/admin", adminMiddleware)

	admin.Get("/", routes.Dashboard)

	// misc
	//app.Get("/ws", websocketF)

	app.Get("/code/:id", routes.CodePage) // code submission testing page
	app.Get("/codetemplate", routes.GetCodeTemplate)

	app.Get("/api/account", routes.GetAccount) // gets the current logged in user with the cookie
	// app.Get("/api/logout", routes.Logout)      // removes the cookie

}

func main() {
	// database.ConnectDb() // MICROSERVICES :_)

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// go sendPostReq(app)

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// prevent the app from caching
	app.Use(func(c *fiber.Ctx) error {

		c.Set("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Set("Pragma", "no-cache")
		c.Set("Expires", "0")
		c.Response().Header.Set("Access-Control-Allow-Origin", "https://irvyn.xyz")
		c.Response().Header.Set("Access-Control-Allow-Credentials", "true")
		c.Response().Header.Set("Access-Control-Allow-Headers", "Content-Type, Set-Cookie, Cookie")

		return c.Next()
	})

	app.Static("/", "./views/public")

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

func adminMiddleware(c *fiber.Ctx) error {
	fmt.Println("admin middleware called", c.Path())

	cookie := c.Cookies("jwt")
	if cookie == "" {
		return c.Redirect("/login")
	} else {
		// check if the user is an admin
		client := resty.New()
		resp, err := client.R().SetHeader("Cookie", "jwt="+cookie).Get("http://api.irvyn.xyz/admin")
		if err != nil {
			return c.Status(500).SendString("Error fetching data from API")
		}
		if resp.StatusCode() == 200 {
			return c.Next()
		} else {
			return c.Redirect("/login")
		}
	}
}
