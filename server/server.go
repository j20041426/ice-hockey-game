package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
)

// 客户端管理进程
type ClientManager struct {
	clients []*Client
	rooms   []*Room
}

func (manager *ClientManager) start() {

}

func main() {
	manager := ClientManager{
		clients: []*Client{},
		rooms:   []*Room{},
	}
	go manager.start()
	http.HandleFunc("/init", initHandler)
	fmt.Println("starting server at port 38080")
	http.ListenAndServe(":38080", nil)
}

func initHandler(res http.ResponseWriter, req *http.Request) {
	conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if error != nil {
		http.NotFound(res, req)
		return
	}
	id, error := uuid.NewV4()
	if error != nil {
		http.NotFound(res, req)
		return
	}
	client := &Client{id: id.String(), socket: conn, send: &Message{event: 0, data: ""}}
	fmt.Print(client)
}

// 房间结构
type Room struct {
	id      string
	clients []*Client
	code    string
}

// 客服端结构
type Client struct {
	id     string
	socket *websocket.Conn
	send   *Message
}

// 消息结构
type Message struct {
	event int
	data  string
}
