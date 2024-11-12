package main

import (
	configDB "go-websocket-fiber/config/database"
	"go-websocket-fiber/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	
	// Initialize the database connection
	configDB.InitDB()
	configDB.InitPrismaDB()

	// Initialize the Fiber app
	app := fiber.New()

	// Set up routes
	routes.SetupRoutes(app)

	// Start the server
	log.Fatal(app.Listen(":8080"))
}
