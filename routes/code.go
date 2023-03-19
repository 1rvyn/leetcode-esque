package routes

import (
	"bytes"
	"encoding/json"
	"fiberWebApi/database"
	"fiberWebApi/models"
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
	"os/exec"

	"github.com/gofiber/fiber/v2"
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
	{Title: "Problems", URL: "/problems"},
	{Title: "Account", URL: "/account"},
	{Title: "Logout", URL: "/api/logout"},
}

func CodePage(c *fiber.Ctx) error {
	questionID := c.Params("id")

	client := resty.New()
	resp, err := client.R().
		SetQueryParam("id", questionID).
		Get("https://api.irvyn.xyz/question/" + questionID)

	fmt.Println("response from the backend was \n", resp)

	if err != nil {
		return c.Status(500).SendString("Error fetching data from API")
	}

	var questionData map[string]interface{}
	err = json.Unmarshal(resp.Body(), &questionData)
	if err != nil {
		return c.Status(500).SendString("Error parsing API response")
	}

	return c.Render("code", fiber.Map{
		"Pages":             pages2,
		"Question":          questionData["problem"],
		"ExampleInput":      questionData["example_input"],
		"ExampleAnswer":     questionData["example_answer"],
		"Codetemplate":      questionData["template_code"].(map[string]interface{})["python"], // TODO: Update based on current language (get from session?)
		"ProblemType":       questionData["problem_type"],
		"ProblemDifficulty": questionData["problem_difficulty"],
	})
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

	file, err := os.Create("./remotecode/code.py")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(data["codeitem"])
	if err != nil {
		panic(err)
	}

	// run the python code

	cmd := exec.Command("python", "./remotecode/code.py")

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

func GetCodeTemplate(c *fiber.Ctx) error {
	language := c.Query("language")

	fmt.Println("sending new language is: ", language)
	// TODO: Make it so that it will take the question ID and get the code template for that question
	// question := c.Query("question")

	var codeTemplate string

	switch language {
	case "python":
		// get the python code template
		codeTemplate = `def twoSum(nums, target):
		# your code here
		answer = []
		return answer`

	case "javascript":
		// get the javascript code template
		codeTemplate = `var twoSum = function(nums, target) {
		// your code here
		answer = []
		return answer
	};`

	case "go":
		// get the go code template
		codeTemplate = `func twoSum(nums []int, target int) []int {
		// your code here
		answer := []int{}
		return answer
	}`

	default:
		codeTemplate = "Error: No code template found"
	}

	fmt.Println("the code template is: ", codeTemplate)
	return c.JSON(fiber.Map{
		"Codetemplate": codeTemplate,
	})
}
