package handlers

import (
	"fmt"
	"goapp/database"
	"goapp/models"
	"goapp/utils"
	"goapp/validator"

	"github.com/gofiber/fiber/v2"
)

func GetNewTokenHandler(c *fiber.Ctx) error {
	var (
		reqBody    models.GetNewTokenRequest
		jtoken     = ""
		jwtCount   = 0
		statusCode = fiber.StatusOK
		errorMsg   = "New token generation successful"
		err        error
	)

	if err = c.BodyParser(&reqBody); err != nil {
		fmt.Print("Error parsing the request body : ", err)
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

	err = database.DB.QueryRow(models.Config.Queries.SelectQueries.CheckUserJwt, reqBody.Username, reqBody.JwtToken).Scan(&jwtCount)
	if err != nil {
		fmt.Println("Error getting the JWT count from the database : ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error getting the JWT count from the database",
		})
	}

	if jwtCount != 0 {
		// Generating new JWT token
		jtoken, err = utils.GenerateJWT(reqBody.Username)
		if err != nil {
			fmt.Println("Error generating the JWT token : ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error generating the JWT token",
			})
		}

		// Updating the new JWT token in the database
		_, err = database.DB.Exec(models.Config.Queries.InsertQueries.InsertJwt, reqBody.Username, jtoken)
		if err != nil {
			fmt.Println("Error updating the new JWT token in the database : ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error updating the new JWT token in the database",
			})
		}
	}
	return c.Status(statusCode).JSON(fiber.Map{
		"message":   errorMsg,
		"new_token": jtoken,
	})
}
