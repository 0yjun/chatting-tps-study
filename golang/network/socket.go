package network

import (
	"chat-server-golang/types"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// WebSocket 업그레이드 설정
var Upgrader = &websocket.Upgrader{
	ReadBufferSize:  types.SocketBufferSize,  // types에서 상수를 가져와 사용
	WriteBufferSize: types.MessageBufferSize,
	CheckOrigin: func(r *http.Request) bool {return true},
}


type message struct{
	Name string
	Message string
	Time uint64
}
type Room struct{
	Forward chan *message //수신메세지 보관
	Join  chan *client //소켓이 연결될때 작동
	Leave chan *client
	clients map[*client]bool

}

type client struct{
	Send chan *message
	Room *Room
	Name string
	Socket *websocket.Conn
}

func (r *Room) RunInit() {
	//ROOM에 있는 모든 체널을 받음
	for{
		select{
		case client := <- r.Join:
			r.clients[client] = true
		case client := <- r.Leave:
			r.clients[client] = false
			close(client.Send)
			delete(r.clients,client)
		case msg := <- r.Forward:
			for client := range r.clients{
				client.Send <- msg
			}
		}
	}
}

func (c *client)Read()  {
	defer c.Socket.Close()
	//클라이언트가 메시지를 읽는 함수
	for msg := range c.Send{
		err:= c.Socket.ReadJSON(msg)
		if err != nil {
			if !websocket.IsUnexpectedCloseError(err,websocket.CloseGoingAway){
				break
			}else{
				panic(err)
			}
			
		}else{
			msg.Time= uint64(time.Now().Unix())
			msg.Name = c.Name
			c.Room.Forward <- msg
		}
	}
}

func (c *client)write()  {
	defer c.Socket.Close()
	//클라이언트가 메시지를 읽는 함수
	for msg := range c.Send{
		err:= c.Socket.WriteJSON(msg)
		if err != nil {
			panic(err)
		}
	}
}
func NewRoom() *Room  {
	return &Room{
		Forward: make(chan *message),
		Join: make(chan *client),
		Leave: make(chan *client),
		clients: make(map[*client]bool),
	}
}

func (r*Room) SocketServe(c*gin.Context) {
	socket,err := Upgrader.Upgrade(c.Writer, c.Request,nil)
	if err != nil {
		panic(err)
	}
	userCookie, err := c.Request.Cookie("auth")
	if err != nil {
		panic(err)
	}
	client:= &client{
		Socket: socket,
		Send: make(chan *message, types.MessageBufferSize),
		Room: r,
		Name: userCookie.Value,

	}

	r.Join <- client

	defer func ()  {r.Leave <-client}()

	go client.write()
	client.Read()
}