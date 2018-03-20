package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
)

// 客户端管理进程
type ClientManager struct {
	clients    map[*Client]string
	rooms      map[string]*Room
	messages   chan *Message
	unregister chan *Client
}

func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.unregister:
			if _, ok := manager.clients[conn]; ok {
				close(conn.send)
				delete(manager.clients, conn)
			}

		}
	}
}

var manager = &ClientManager{
	clients:  make(map[*Client]string),
	rooms:    make(map[string]*Room),
	messages: make(chan *Message),
}

// 客户端结构
type Client struct {
	id     string
	socket *websocket.Conn
	send   chan []byte
}

func (c *Client) read() {
	defer func() {
		manager.unregister <- c
		c.socket.Close()
	}()

	for {
		_, data, err := c.socket.ReadMessage()
		if err != nil {
			manager.unregister <- c
			c.socket.Close()
			break
		}
		p := &Point{}
		json.Unmarshal([]byte(data), p)
		manager.messages <- &Message{client: c, point: p}
	}
}

func main() {
	fmt.Println("starting server at port 38080")
	go manager.start()
	http.HandleFunc("/init", initHandler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/join", joinHandler)
	http.ListenAndServe(":38080", nil)
}

func initHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	res.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	res.Header().Set("content-type", "application/json")             //返回数据格式是json

	conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if error != nil {
		http.NotFound(res, req)
		return
	}
	id, _ := uuid.NewV4()
	client := &Client{id: id.String(), socket: conn, send: make(chan []byte)}

	go client.read()
}

func createHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	res.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	res.Header().Set("content-type", "application/json")             //返回数据格式是json
	id, _ := uuid.NewV4()
	room := &Room{id: id.String(), clients: make([]*Client, 2), code: randNum(4), is_full: false}
	manager.rooms[room.id] = room
	res.Write([]byte(room.code))
}

func joinHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	res.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	res.Header().Set("content-type", "application/json")             //返回数据格式是json
}

func randNum(n int) string {
	var ret = ""
	for i := 0; i < n; i++ {
		ret += strconv.Itoa(rand.Intn(9))
	}
	return ret
}

// 房间结构
type Room struct {
	id      string
	clients []*Client
	code    string
	is_full bool
}

// 消息结构
type Message struct {
	client *Client
	point  *Point
}

// 坐标结构
type Point struct {
	x int
	y int
}
