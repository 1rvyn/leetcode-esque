package main

import (
	"bufio"
	"fiberWebApi/database"
	"fiberWebApi/models"
	"fiberWebApi/routes"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
	"github.com/gofiber/websocket/v2"
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
	{Title: "Code", URL: "/code"},
	{Title: "Problems", URL: "/problems"},
	{Title: "Account", URL: "/account"},
	{Title: "Logout", URL: "/logout"},
}

func welcomeHome(c *fiber.Ctx) error {
	//  templating here

	// print the users cookies

	// fmt.Println(c.GetReqHeaders())
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

	fmt.Println("\n the ip here is:", c.IP())
	fmt.Println("\n The x-forwarded-for header is: ", c.Get("x-forwarded-for"))

	return c.Render("nice", fiber.Map{
		"Pages":     pages,
		"ActiveURL": activeURL,
		"User":      "Irvyn Hall",
		"Email":     "irvynhall@gmail.com",
		"Status":    "logging--in",
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
	//fmt.Println("the problems page was accessed and the cookie is \n : ", c.Cookies("jwt"))

	activeURL := c.Path()

	// read the problems from the list.txt file inside the questions folder

	// Users/irvyn/go/src/fiberWebApi/questions/q-1/list.txt

	code, err := os.Open("./questions/q-1/list.txt")
	// fmt.Println("code is : ", code)
	if err != nil {
		fmt.Println("error reading question list file: ", err)
	}

	defer code.Close()

	var lines []string

	scanner := bufio.NewScanner(code)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error reading question list file: ", err)
	}

	currCookie := c.Cookies("jwt")

	if currCookie == "" {
		return c.Render("problems", fiber.Map{
			"Pages":       pages,
			"ActiveURL":   activeURL,
			"ProblemList": lines,
		})
	} else {
		return c.Render("problems", fiber.Map{
			"Pages":       pages2,
			"ActiveURL":   activeURL,
			"ProblemList": lines,
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
	// same thing for both

	// userRole auth routes/pages
	admin := app.Group("/admin", adminMiddleware)

	admin.Get("/", routes.Dashboard)

	// misc
	//app.Get("/ws", websocketF)

	app.Get("/code", routes.CodePage) // code submission testing page

	// account system
	// app.Post("/api/register", routes.CreateAccount) // store creds in the database
	// app.Post("/api/login", routes.GetLogin)         // checks the creds against the stored db creds

	app.Get("/api/account", routes.GetAccount) // gets the current logged in user with the cookie
	// app.Get("/api/logout", routes.Logout)      // removes the cookie
	app.Post("/api/code", routes.PythonCode) // get python code from textarea

}

func main() {
	// database.ConnectDb() // MICROSERVICES :_)

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,

		//ViewsLayout: "layouts/layout",
	})

	app.Use("/ws", func(c *fiber.Ctx) error {
		if c.Get("host") == "localhost:3000" {
			c.Locals("Host", "Localhost:3000")
			return c.Next()
		}
		return c.Status(403).SendString("Request origin not allowed")
	})

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		fmt.Println(c.Locals("Host")) // "Localhost:3000"
		for {
			mt, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recieved the message : %s", msg)
			err = c.WriteMessage(mt, []byte("hello from irvyn's backend go server"))
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}))

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

// TODO: revist this for a test case to test the api

// func sendPostReq(app *fiber.App) {
// 	type RequestBody struct {
// 		User     string `json:"user"`
// 		Password string `json:"password"`
// 		Email    string `json:"email"`
// 	}
// 	body := RequestBody{
// 		User:     "testingfrontend_box",
// 		Password: "password123",
// 		Email:    "testererere",
// 	}
// 	jsonBody, err := json.Marshal(body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Send POST request
// 	_, err = http.Post("https://api.irvyn.xyz/api/register", "application/json", bytes.NewBuffer(jsonBody))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("POST request sent successfully")
// }

func adminMiddleware(c *fiber.Ctx) error {
	// user := getUserFromSession(c)

	fmt.Println("the current context is : ", c)
	cookie := c.Cookies("jwt")

	if cookie == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// get the session from the cookie
	var session []models.Session
	database.Database.Db.Where("cookie = ?", cookie).Last(&session)

	// print the users email from the session we have found
	// fmt.Println("session email: ", session[0].Email)

	// get the user from the session
	var user []models.Account
	database.Database.Db.Where("email = ?", session[0].Email).Last(&user)

	// print the user we found
	fmt.Println("This users role is : ", user[0].UserRole)

	if user[0].UserRole == 2 {
		fmt.Println("the user is an admin")
		return c.Next()
	}

	// if the user is not an admin
	return c.Status(401).JSON(fiber.Map{
		"message": "Unauthorized",
	})
}
