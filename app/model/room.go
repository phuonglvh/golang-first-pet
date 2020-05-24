package model

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/matryer/goblueprints/chapter1/trace"
)

// Room define model for a room
type Room struct {
	ID string
	// Forward is a channel that holds incoming messages
	// that should be forwarded to the other clients.
	Forward chan []byte

	// join is a channel for clients wishing to join the room.
	Join chan *Client

	// Leave is a channel for clients wishing to Leave the room.
	Leave chan *Client

	// clients holds all current clients in this room.
	clients map[*Client]bool

	// Tracer will receive trace information of activity
	// in the room.
	Tracer trace.Tracer

	messages [][]byte

	messageses []Message
}

// NewRoom will create a new room
func NewRoom(ID string) *Room {
	return &Room{
		ID:         ID,
		Forward:    make(chan []byte),
		Join:       make(chan *Client),
		Leave:      make(chan *Client),
		clients:    make(map[*Client]bool),
		Tracer:     trace.Off(),
		messages:   [][]byte{},
		messageses: []Message{},
	}
}

// Run will start listen to room event
func (room *Room) Run() {
	for {
		select {
		case client := <-room.Join:
			// joining
			room.clients[client] = true
			room.Tracer.Trace("New client joined")
			room.sendPastMessages(client)
		case client := <-room.Leave:
			// leaving
			delete(room.clients, client)
			close(client.Send)
			room.Tracer.Trace("Client left")
		case msg := <-room.Forward:
			msgString := string(msg)
			room.messageses = append(room.messageses, Message{ID: "dafdsa", Sender: "dafdsa", Content: msgString})
			room.Tracer.Trace("Message received: ", msgString)
			room.messages = append(room.messages, msg)
			// forward message to all clients
			room.sendMessageToAll(msg)
			// default:
			// 	room.Tracer.Trace("default room Run: ")
		}
	}
}

func (room *Room) sendMessageToAll(messageBytes []byte) {
	for client := range room.clients {
		room.sendMessageToClient(client, messageBytes)
	}
}

func (room *Room) sendMessageToClient(client *Client, messageBytes []byte) {
	client.Send <- messageBytes
	room.Tracer.Trace(" -- sent a message to client ", string(messageBytes))
}

func (room *Room) sendPastMessages(client *Client) {
	room.Tracer.Trace(" -- start sending past messages to client")
	for _, msg := range room.messages {
		room.sendMessageToClient(client, msg)
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

func (room *Room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	client := &Client{
		Socket: socket,
		Send:   make(chan []byte, messageBufferSize),
		Room:   room,
	}
	room.Join <- client
	// defer func() { room.Leave <- client }()
	// go client.Write()
	client.Read()
}
