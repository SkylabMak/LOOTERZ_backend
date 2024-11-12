package routes

import (
	httpTest "go-websocket-fiber/controllers/http/test"
	"go-websocket-fiber/controllers/socket"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/testGetAllUser", httpTest.TestGetAllUser)   // HTTP endpoint for user creation
	app.Post("/testGetAllUserGROM", httpTest.TestGetAllUser)   // HTTP endpoint for user creation

	//socket
	app.Get("/ws", websocket.New(socket.WebSocketHandler))
}
