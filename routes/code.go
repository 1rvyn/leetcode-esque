package routes

import (
	"bytes"
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

	return c.Render("code", fiber.Map{
		"Pages":        pages2,
		"ActiveURL":    activeURL,
		"Question":     questiontemplate,
		"Codetemplate": codetemplate, // this is the code which gets pre-loaded into the code editor
		"Code":         "This is the code page",
	})
}

func PythonCode(c *fiber.Ctx) error {
	// console log the body of the request

	// print the cookie that was sent with the request
	fmt.Println(c.Cookies("jwt"))

	fmt.Println("code was submitted")

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// save the code to the database

	//fmt.Println("the body of the data is: ", data)
	//
	//fmt.Println("payloadItem: ", data["codeitem"])

	// TODO: send error messages properly if code doesn't execute

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

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "code was submitted",
		"output":  string(output),
		"error":   string(errorOutput),
	})
	// since we have sent the response here we should save the code to the database

	//submission := models.Submission{
	//	Code:       data["codeitem"],
	//	Cookie:     c.Cookies("jwt"),
	//	Email:      c.Cookies("email"),
	//	IP:         c.IP(),
	//	Successout: string(output),
	//	Errorout:   string(errorOutput),
	//}
}
