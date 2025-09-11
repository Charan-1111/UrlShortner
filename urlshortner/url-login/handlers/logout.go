package handlers

import (
	"fmt"
	"strings"

	"goapp/database"
	"goapp/models"
	"goapp/utils"

	"github.com/gofiber/fiber/v2"
)

func LogoutHandler(c *fiber.Ctx) error {
	// reading the authentication token from the request header

	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing or invalid Authorization header",
		})
	}

	// TODO : Invalidate the token (This can be done by maintaining a token blacklist or changing a token version in the database)

	// For now, just return a success message

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid Authorization header format",
		})
	}

	token := parts[1] // the actual token part

	claims, isValidToken := utils.ValidateJWT(token)
	if isValidToken != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})
	}

	username := claims.Username

	fmt.Println("Logging out user: ", username)

	_, err := database.DB.Exec(models.Config.Queries.DeleteQueries.DeleteJwt, username, token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error deleting the JWT token from the database" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logout successful",
	})
}
