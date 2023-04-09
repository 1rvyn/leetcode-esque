package routes

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

type Page struct {
	Title string
	URL   string
}

var pages2 = []Page{
	{Title: "Home", URL: "/"},
	{Title: "Code", URL: "/code/1"},
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
		"Title":             questionData["title"],
		"QuestionID":        questionID,
		"Question":          questionData["problem"],
		"ExampleInput":      questionData["example_input"],
		"ExampleAnswer":     questionData["example_answer"],
		"Codetemplate":      questionData["template_code"].(map[string]interface{})["python"], // TODO: Update based on current language (get from session?)
		"ProblemType":       questionData["problem_type"],
		"ProblemDifficulty": questionData["problem_difficulty"],
	})
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
		codeTemplate = `def two_sum(nums, target):
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
