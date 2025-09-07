package handlers

import (
	"goapp/models"

	"github.com/gofiber/fiber/v2"
)

func LoginHandler(c *fiber.Ctx) error {
	var (
		reqBody models.LoginRequest
		err error
	)

	if err = c.BodyParser(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}
	return nil
}