package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/phuonglvh/golang-first-pet/app/models"
	"github.com/phuonglvh/golang-first-pet/app/views"
	httpUtils "github.com/phuonglvh/golang-first-pet/utils/http"
	logger "github.com/phuonglvh/golang-first-pet/utils/logger"
)

// ChatHandler stores list of created rooms, serves http
type ChatHandler struct {
	Rooms map[string]*models.Room
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

// ChatWSHandler is the database of chat including rooms, hanlde ws requests
var ChatWSHandler = &ChatHandler{Rooms: make(map[string]*models.Room)}

// ChatPageHandler handle render the Chat Page
func ChatPageHandler(w http.ResponseWriter, r *http.Request) {
	roomID := httpUtils.ParseParam(r, "id")
	room := ChatWSHandler.Rooms[roomID]
	if room == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := views.ViewResolver.ExecuteTemplate(w, "chat", nil)
	if err != nil {
		fmt.Println(err)
	}
}

// ChatPageHandlerV2 handle render the sandbox chat page
func ChatPageHandlerV2(w http.ResponseWriter, r *http.Request) {
	roomID := httpUtils.ParseParam(r, "id")
	room := ChatWSHandler.Rooms[roomID]
	if room == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := views.ViewResolver.ExecuteTemplate(w, "chatbox", nil)
	if err != nil {
		fmt.Println(err)
	}
}

// GetRoomMessages handles get and return message of the room
func GetRoomMessages(w http.ResponseWriter, r *http.Request) {
	roomID := httpUtils.ParseParam(r, "id")
	room := ChatWSHandler.Rooms[roomID]
	if room == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(room.GetMessages())
}

func (chat *ChatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("X-Authorization")
	if err != nil {
		logger.Error.Println(err.Error())
	}
	clientID := cookie.Value
	var roomID string = httpUtils.ParseParam(r, "id")
	if len(roomID) < 1 {
		logger.Error.Println("Query param 'id' is missing")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var room *models.Room = chat.Rooms[roomID]
	if room == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	go room.Run()
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error.Fatalf("Chat websocket has encountered error while serving http: %s", err)
		return
	}
	client := &models.Client{
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
