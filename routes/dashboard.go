package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Dashboard(c *fiber.Ctx) error {

	return c.SendString("admin dashboard")
}
