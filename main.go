package main

import (
	"log"
	"middle-developer-test/pkg/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// setting up fiber
	app := fiber.New()

	app.Use(cors.New(), logger.New())

	routes.AppRoutes(app)

	err := app.Listen(":8000")
	if err != nil {
		log.Printf("error during listening to port 8000. reason: %v", err)
	}
}
