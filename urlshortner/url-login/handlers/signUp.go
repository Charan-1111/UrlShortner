package handlers

import (
	"fmt"
	"goapp/database"
	"goapp/models"
	"goapp/utils"
	"goapp/validator"

	"github.com/gofiber/fiber/v2"
)

func SignUpHandler(c *fiber.Ctx) error {
	var (
		requestBody models.SignUpRequest
		err         error
		emailCnt    = 0
		usernameCnt = 0
		statusCode = fiber.StatusOK
		errMsg = "User registered successfully..."
	)

	if err = c.BodyParser(&requestBody); err != nil {
		fmt.Print("Error parsing request body: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// validating the request body parameters, so that getting the expected values in each field
	msg := validator.ValidateRequestBody(requestBody)
	if msg != "" {
		fmt.Println("Validation error: ", msg)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": msg,
		})
	}

	// checking if the user already exists
	err = database.DB.QueryRow(models.Config.Queries.SelectQueries.CheckEmail, requestBody.Email).Scan(&emailCnt)
	if err != nil && err.Error() != "sql: no rows in result set" {
		fmt.Println("Error checking email existence: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	switch emailCnt {
	case 0:
		// checking for the username already taken or not by any other user...
		err = database.DB.QueryRow(models.Config.Queries.SelectQueries.CheckUsername, requestBody.Username).Scan(&usernameCnt)
		if err != nil && err.Error() != "sql: no rows in result set" {
			fmt.Println("Error checking username existence: ", err)
			statusCode = fiber.StatusInternalServerError
			errMsg = "Internal server error"
		}

		if usernameCnt == 0 {
			// hash the password before storing it in the database
			requestBody.Password, err = utils.HashPassword(requestBody.Password)
			if err != nil {
				fmt.Println("Error hashing password : ", err)
				statusCode = fiber.StatusInternalServerError
				errMsg = "Internal server errror while hashing the password : " + err.Error()
				break
			}


			// inserting the new user into the database
			_, err = database.DB.Exec(models.Config.Queries.InsertQueries.InsertUser, requestBody.Username, requestBody.Password, requestBody.Email)
			if err != nil {
				fmt.Println("Error inserting new user: ", err)
				statusCode = fiber.StatusInternalServerError
				errMsg = "Internal server error while inserting the new user"
				break
			}
		} else {
			statusCode = fiber.StatusConflict
			errMsg = "Username already taken, please choose another one"
		}
	default:
		statusCode = fiber.StatusConflict
		errMsg = "Email already registered, please use another email"
	}

	
	return c.Status(statusCode).JSON(fiber.Map{
		"message": errMsg,
	})
}
