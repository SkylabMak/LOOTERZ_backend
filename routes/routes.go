package routes

import (
	httpHome "LOOTERZ_backend/controllers/http/home"
	httpListRoom "LOOTERZ_backend/controllers/http/listRoom"
	httpLobby "LOOTERZ_backend/controllers/http/lobby"
	httpSetting "LOOTERZ_backend/controllers/http/setting"
	httpTest "LOOTERZ_backend/controllers/http/test"
	socketTest "LOOTERZ_backend/controllers/socket/test"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func SetupRoutes(app *fiber.App) {
	// HTTP endpoint test
	app.Post("/testGetAllUser", httpTest.TestGetAllUser)
	app.Post("/testGetAllUserGROM", httpTest.TestGetAllUser)
	app.Get("/testFunction", httpTest.TestFuntion)
	app.Post("/testSendReids", httpTest.SendMessageHandler)
	// HTTP endpoint
	//httpHome ball
	app.Post("/enter",httpHome.Enter)
	
	//httpListRoom game
	listRoomGroup := app.Group("/listRoom")
	listRoomGroup.Get("/getListRoom/:page",httpListRoom.GetListRoom)
	listRoomGroup.Get("/enterRoom",httpListRoom.EnterRoom)
	
	//httpLobby game
	app.Post("/playerReady",httpLobby.PlayerReady)
	
	//httpSetting Guy
	settingRoomGroup := app.Group("/setting")
	settingRoomGroup.Post("/createRoom",httpSetting.CreateRoom)
	settingRoomGroup.Post("/changeSetting",httpSetting.ChangeSetting)

	//Each WebSocket handler runs as a goroutine by default because Fiber
	//socket
	//test
	app.Get("/ws/:roomID", websocket.New(socketTest.WebSocketHandler))
	app.Get("/ws/reids/:roomID", websocket.New(socketTest.WebSocketHandler_redis))
}
