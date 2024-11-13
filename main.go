package main

import (
	configDB "LOOTERZ_backend/config/database"
	"LOOTERZ_backend/routes"
	"LOOTERZ_backend/services/socket"

	// "go-websocket-fiber/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	
	// Initialize the database connection
	configDB.InitDB()
	configDB.InitPrismaDB()
	socket.InitRedis()

	// Initialize the Fiber app
	app := fiber.New()

	// Set up routes
	routes.SetupRoutes(app)

	// Start the server
	log.Fatal(app.Listen(":8080"))
}
