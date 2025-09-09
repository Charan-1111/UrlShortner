package handlers

import (
	"fmt"
	"goapp/models"

	"goapp/validator"

	"github.com/gofiber/fiber/v2"
)

func LoginHandler(c *fiber.Ctx) error {
	var (
		reqBody  models.LoginRequest
		username string
		err      error
	)

	if err = c.BodyParser(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	// Validating the request body parameters, so that getting the expected values in each fields
	msg := validator.ValidateRequestBody(reqBody)
	if msg != "" {
		fmt.Println("Invalid request body : ", msg)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": msg})
	}

	if reqBody.Username != "" {
		username = reqBody.Username
	} else {
		username = reqBody.Email
	}
	return nil
}
