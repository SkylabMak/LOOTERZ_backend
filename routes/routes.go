package routes

import (
	httpTest "go-websocket-fiber/controllers/http/test"
	socketTest "go-websocket-fiber/controllers/socket/test"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/testGetAllUser", httpTest.TestGetAllUser)   // HTTP endpoint for user creation
	app.Post("/testGetAllUserGROM", httpTest.TestGetAllUser)   // HTTP endpoint for user creation

	//Each WebSocket handler runs as a goroutine by default because Fiber
	//socket
	app.Get("/ws/:roomID", websocket.New(socketTest.WebSocketHandler))
}
