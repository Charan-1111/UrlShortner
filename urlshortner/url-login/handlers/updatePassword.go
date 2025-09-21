package handlers

import (
	"fmt"
	"goapp/models"
	"goapp/validator"

	"github.com/gofiber/fiber/v2"
)

func UpdatePasswordHandler(c *fiber.Ctx) error {
	var (
		reqBody models.UpdatePasswordRequest
		err     error
	)

	if err = c.BodyParser(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate the request body
	msg := validator.ValidateRequestBody(reqBody)
	if msg != "" {
		fmt.Println("Validation error: ", msg)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": msg,
		})
	}


	

	return nil
}
