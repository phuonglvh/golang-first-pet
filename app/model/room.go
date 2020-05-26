package model

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"

	"github.com/phuonglvh/golang-first-pet/util/logger"
)

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

// Room define model for a room
type Room struct {
	ID string
	// Forward is a channel that holds incoming messages
	// that should be forwarded to the other clients.
	Forward chan *RawClientMessage

	// join is a channel for clients wishing to join the room.
	Join chan *Client

	// Leave is a channel for clients wishing to Leave the room.
	Leave chan *Client

	// clients holds all current clients in this room.
	clients map[*Client]bool

	// Tracer will receive trace information of activity
	// in the room.
	// Tracer trace.Tracer

	messages []*Message
}

// NewRoom will create a new room
func NewRoom(ID string) *Room {
	return &Room{
		ID:       ID,
		Forward:  make(chan *RawClientMessage),
		Join:     make(chan *Client),
		Leave:    make(chan *Client),
		clients:  make(map[*Client]bool),
		messages: []*Message{},
	}
}

// Run will start listen to room event
func (room *Room) Run() {
	for {
		select {
		case client := <-room.Join:
			// joining
			room.clients[client] = true
			logger.Info.Printf("Client %s has joined the room %s", client.ID, room.ID)
			room.sendPastMessages(client)
		case client := <-room.Leave:
			// leaving
			delete(room.clients, client)
			close(client.Send)
			logger.Info.Printf("Client %s has left the room %s", client.ID, room.ID)
		case fwdMsg := <-room.Forward:
			message := &Message{
				ID:        uuid.New().String(),
				Content:   fwdMsg.Content,
				Sender:    fwdMsg.Sender,
				Timestamp: time.Now().Unix() * 1000,
			}
			room.messages = append(room.messages, message)
			logger.Trace.Printf("Client has sent message to others in room %s: %s", room.ID, fwdMsg.Content)
			// forward message to all clients
			room.sendMessageToAll(message)
		}
	}
}

func (room *Room) sendMessageToAll(message *Message) {
	for client := range room.clients {
		room.sendMessageToClient(client, message)
	}
}

func (room *Room) sendMessageToClient(client *Client, messagge *Message) {
	bytes, _ := json.Marshal(messagge)
	client.Send <- bytes
	logger.Trace.Printf("Sent a message to client %s: %s", client.ID, string(bytes))
}

func (room *Room) sendPastMessages(client *Client) {
	logger.Trace.Printf("Sending past %d messages to client %s", len(room.messages), client.ID)
	for _, msg := range room.messages {
		room.sendMessageToClient(client, msg)
	}
}

func (room *Room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		logger.Error.Fatalf("Room %s has encountered error while serving http: %s", room.ID, err)
		return
	}
	logger.Info.Printf("Room %s is waiting for clients", room.ID)
	client := &Client{
		Socket: socket,
		Send:   make(chan []byte, messageBufferSize),
		Room:   room,
	}
	room.Join <- client
	client.Read()
}
