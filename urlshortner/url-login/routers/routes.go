package routers

import (
	"fmt"
	"goapp/handlers"

	"runtime/debug"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetUpFiberServer() *fiber.App {
	app := fiber.New(
		fiber.Config{
			Prefork: false,
		},
	)

	// Enabling panic recovery middleware to handle any unexpected errors gracefully
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true, // show stack trace
		StackTraceHandler: func(c *fiber.Ctx, e any) {
			// Print the panic error and stack trace to the console
			fmt.Printf("Panic: %v\n", e)
			fmt.Printf("Stack Trace:\n%s\n", debug.Stack())
		},
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
		MaxAge:           3600,
	}))

	// Setting up the routes
	apiGroup := app.Group("/urlshortner")
	apiGroup.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("URL Shortener is healthy!")
	})

	apiGroup.Post("/signup", handlers.SignUpHandler)
	apiGroup.Post("/login", handlers.LoginHandler)
	apiGroup.Post("/getnewtoken", handlers.GetNewTokenHandler)
	apiGroup.Post("/logout", handlers.LogoutHandler)
	return app
}
