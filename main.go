package main

import (
	"log"
	"middle-developer-test/pkg/config"
	"middle-developer-test/pkg/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// setting up fiber
	config.LoadConfig(".env.local")

	app := fiber.New()

	app.Use(cors.New(), logger.New())

	log.Print(os.Getenv("DB_URL"))
	log.Print(os.Getenv("SERVER_URL"))

	routes.EmployeeRoutes(app)

	err := app.Listen(os.Getenv("SERVER_URL"))
	if err != nil {
		log.Printf("error during listening to port. reason: %v", err)
	}
}
