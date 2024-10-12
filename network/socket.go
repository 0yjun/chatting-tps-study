package network

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
)

var Upgrader = $websocket.Upgrade{RealBufferSize:socketBufferSize,WritterBufferSize:messageBufferSize};


type message struct{
	Name string
	Message string
	Time uint64
}
type Room struct{
	Forward chan *message //수신메세지 보관
	Join  chan *client //소켓이 연결될때 작동
	Leave chan *client
	Clients map[*client]bool

}

type client struct{
	Send chan *message
	Room *Room
	Name string
	Socket websocket.Conn
}

func NewRoom() *Room  {
	return &Room{
		Forward: make(chan *message),
		Join: make(chan *client),
		Leave: make(chan *client),
		Clients: make(map[*client]bool),
	}
}

func (r*Room) SocketServe(*gin.Context) {
	
}