package services

import (
	"errors"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"log"
	"time"
	"webapp/pkg/utils"
)

type WebsocketService struct {
	Server *socketio.Server
}

type WebSocketRequest struct {
	Room string `json:"room"`
	Msg  string `json:"msg"`
}

type ConnectRoomRequest struct {
	Room string `json:"room"`
}

func NewWebsocketService() *WebsocketService {
	server := socketio.NewServer(&engineio.Options{
		PingInterval: 20 * time.Second,
		PingTimeout:  60 * time.Second,
	})

	return &WebsocketService{
		Server: server,
	}
}

func (ws *WebsocketService) AttachToRouter(r *gin.Engine) {
	r.Any("/socket.io/*any", gin.WrapH(ws.Server))
}

func (ws *WebsocketService) Start() {

	ws.Server.OnConnect("/", func(s socketio.Conn) error {

		//Verify jwt token in header
		header := s.RemoteHeader()
		_, err := utils.ValidateToken(header.Get("JWT"))
		if err != nil {
			err := s.Close()
			if err != nil {
				return errors.New("JWT token invalid")
			}
		}

		s.SetContext("")
		log.Printf("connected: %v", s.ID())
		return nil
	})

	ws.Server.OnEvent("/", "join", func(s socketio.Conn, connectRoomRequest ConnectRoomRequest) {
		log.Printf("join event triggered with room: %v", connectRoomRequest.Room)
		if connectRoomRequest.Room == "" {
			log.Println("room name is empty")
			return
		}

		s.Join(connectRoomRequest.Room)
		log.Printf("joined: %v", connectRoomRequest.Room)
	})

	ws.Server.OnEvent("/", "leave", func(s socketio.Conn, room string) {
		s.Leave(room)
	})

	ws.Server.OnEvent("/", "message", func(s socketio.Conn, webSocketRequest WebSocketRequest) {

		//tested on postman response will always include 42... don't ask me why I didn't trim it
		ws.Server.BroadcastToRoom("", webSocketRequest.Room, "Data: ", webSocketRequest.Msg)
		log.Printf("message event triggered with room: %v and message %v", webSocketRequest.Room, webSocketRequest.Msg)

	})

	ws.Server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Printf("closed: %v", reason)
	})

	go func() {
		err := ws.Server.Serve()
		if err != nil {
			return
		}
	}()
}
