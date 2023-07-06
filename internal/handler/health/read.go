package health

import "github.com/gofiber/fiber/v2"

func HandleHealth(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
