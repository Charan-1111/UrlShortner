package handlers

import (
	"fmt"
	"goapp/database"
	"goapp/models"
	"goapp/utils"

	"goapp/validator"

	"github.com/gofiber/fiber/v2"
)

func LoginHandler(c *fiber.Ctx) error {
	var (
		reqBody    models.LoginRequest
		username   string
		password   string
		jtoken     = ""
		err        error
		statusCode = fiber.StatusOK
		errorMsg   = "Login successful"
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
	
	err = database.DB.QueryRow(models.Config.Queries.SelectQueries.GetHashedPassword, reqBody.Username, reqBody.Email).Scan(&password)
	if err != nil {
		fmt.Println("Error getting the password from the database : ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error getting the password from the database",
		})
	}

	isPassowrdMatched := utils.CheckPasswordHash(reqBody.Password, password)

	if !isPassowrdMatched {
		statusCode = fiber.StatusUnauthorized
		errorMsg = "Invalid credentials"
	} else {
		// Generatign JWT token
		jtoken, err = utils.GenerateJWT(username)
		if err != nil {
			fmt.Println("Error generating JWT token : ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error generating JWT token",
			})
		}

		// Inserting the JWT token into the database
		_, err = database.DB.Exec(models.Config.Queries.InsertQueries.InsertJwt, username, jtoken)
		if err != nil {
			fmt.Println("Error inserting JWT token into the database : ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error inserting JWT token into the database",
			})
		}
	}
	return c.Status(statusCode).JSON(fiber.Map{
		"message": errorMsg,
		"token":   jtoken,
	})
}
