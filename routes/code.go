package routes

import (
	"bytes"
	"fiberWebApi/database"
	"fiberWebApi/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"os/exec"
)

type Submission struct {
	ID         uint   `json:"id"`
	Code       string `json:"code"`
	Cookie     string `json:"cookie"`
	Email      string `json:"email"`
	IP         string `json:"ip"`
	successout string `json:"successout"`
	errorout   string `json:"errorout"`
	//MetaData string `json:"meta_data"`
}

type Page struct {
	Title string
	URL   string
}

var pages2 = []Page{
	{Title: "Home", URL: "/"},
	{Title: "Code", URL: "/code"},
	{Title: "Account", URL: "/account"},
	{Title: "Logout", URL: "/api/logout"},
}

func CodePage(c *fiber.Ctx) error {
	activeURL := c.Path()

	code, err := os.ReadFile("/Users/irvyn/go/src/fiberWebApi/questions/q-1/TwoSum.py")

	question, err := os.ReadFile("/Users/irvyn/go/src/fiberWebApi/questions/q-1/TwoSum.txt")

	if err != nil {
		fmt.Print(err)
	}

	codetemplate := string(code)
	questiontemplate := string(question)

	fmt.Print("the parsed code as a string is: \n", codetemplate)

	currCookie := c.Cookies("jwt")

	if currCookie == "" {
		return c.Redirect("/login")
	} else {
		return c.Render("code", fiber.Map{
			"Pages":        pages2,
			"ActiveURL":    activeURL,
			"Question":     questiontemplate,
			"Codetemplate": codetemplate, // this is the code which gets pre-loaded into the code editor
			"Code":         "This is the code page",
		})
	}
}

func PythonCode(c *fiber.Ctx) error {
	// console log the body of the request

	// print the cookie that was sent with the request

	//TODO: make the code a byte array ?
	fmt.Println(c.Cookies("jwt"))

	fmt.Println("code was submitted")

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	file, err := os.Create("/Users/irvyn/go/src/fiberWebApi/remotecode/code.py")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(data["codeitem"])
	if err != nil {
		panic(err)
	}

	// run the python code

	cmd := exec.Command("python", "/Users/irvyn/go/src/fiberWebApi/remotecode/code.py")

	var outBuf, errBuf bytes.Buffer

	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	err = cmd.Run()

	if err != nil {
		fmt.Println(err.Error())
	}

	output := outBuf.String()
	errorOutput := errBuf.String()

	fmt.Print("the whole cookie is: ", c.Cookies("jwt"))

	// run the .py file

	// send the output back to the client along with setting the status code

	// save the output to the database
	submission := models.Submission{
		Code:       data["codeitem"],
		Cookie:     c.Cookies("jwt"),
		Email:      c.Cookies("email"),
		IP:         c.IP(),
		Successout: output,
		Errorout:   errorOutput,
	}

	// create a submission object and save it to the database
	database.Database.Db.Create(&submission)

	//TODO: right now this saves a submission if the code is unique -
	// but it should save a submission if tests it passes are unique

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "code was submitted",
		"output":  string(output),
		"error":   string(errorOutput),
	})
	// since we have sent the response here we should save the code to the database

}
