package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/phuonglvh/golang-first-pet/app/model"
	"github.com/phuonglvh/golang-first-pet/util/logger"
)

// ChatHandler handle the chat feature
type ChatHandler struct {
	Rooms map[string]*model.Room
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

// ChatIndex handle render Chat Page
func ChatIndex(w http.ResponseWriter, r *http.Request) {
	template := model.TemplateHandler{Filename: "chat.html"}
	template.ServeHTTP(w, r)
}

func (chat *ChatHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		logger.Error.Fatalf("Chat websocket has encountered error while serving http: %s", err)
		return
	}

	vars := mux.Vars(req)
	cookie, _ := req.Cookie("X-Authorization")
	if err != nil {
		logger.Error.Println(err.Error())
	}
	clientID := cookie.Value
	var roomID string = vars["id"]
	if len(roomID) < 1 {
		logger.Error.Println("Query param 'id' is missing")
		return
	}

	var room *model.Room = chat.Rooms[roomID]
	if room == nil {
		room = model.NewRoom(roomID)
		chat.Rooms[roomID] = room
		log.Printf("Add new room: %s to list of rooms", room.ID)
		logger.Trace.Printf("Number of rooms: %d", len(chat.Rooms))
		go room.Run()
	}
	client := &model.Client{
		ID:     clientID,
		Socket: socket,
		Send:   make(chan []byte, messageBufferSize),
		Room:   room,
	}

	if client == nil {
		logger.Error.Fatal("Error while attempting to create a new client")
	} else {
		logger.Info.Printf("Client %s has joined the room %s", client.ID, room.ID)
		room.Join <- client
		defer func() {
			logger.Info.Printf("Client %s has left the room %s", client.ID, room.ID)
			room.Leave <- client
		}()
		go client.Write()
		client.Read()
	}

}
