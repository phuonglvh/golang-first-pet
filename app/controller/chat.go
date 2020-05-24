package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/phuonglvh/pet/app/model"
)

// ChatHandler handle the chat feature
type ChatHandler struct {
	Rooms map[string]*model.Room
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

// ChatIndex handle render Chat Page
func ChatIndex(w http.ResponseWriter, r *http.Request) {
	template := model.TemplateHandler{Filename: "chat.html"}
	template.ServeHTTP(w, r)
}

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

func (chat *ChatHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}

	vars := mux.Vars(req)
	cookie, _ := req.Cookie("X-Authorization")
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(cookie.Value)
	}
	clientId := cookie.Value
	log.Println("number of rooms: ", len(chat.Rooms))

	var roomID string = vars["id"]
	if len(roomID) < 1 {
		log.Println("path param 'id' is missing")
		return
	}

	var room *model.Room = chat.Rooms[roomID]
	if room == nil {
		room = model.NewRoom(roomID)
		chat.Rooms[roomID] = room
		log.Printf("create a new room: %s -> total rooms: %d ", room.ID, len(chat.Rooms))
		go room.Run()
	} else {
		log.Printf("room is existing: %s", room.ID)
	}
	client := &model.Client{
		ID:     clientId,
		Socket: socket,
		Send:   make(chan []byte, messageBufferSize),
		Room:   room,
	}

	if client == nil {
		log.Fatal("error while attempting to create a new client")
	} else {
		log.Printf("Client joins room: %s", room.ID)
		room.Join <- client
		defer func() {
			log.Printf("Client left room: %s", room.ID)
			room.Leave <- client
		}()
		go client.Write()
		client.Read()
	}

}

// func main() {
// 	var addr = flag.String("addr", ":8080", "The addr of the application.")
// 	flag.Parse() // parse the flags

// 	r := model.NewRoom()
// 	r.Tracer = trace.New(os.Stdout)

// 	http.Handle("/", &model.TemplateHandler{Filename: "chat.html"})
// 	http.Handle("/room", r)

// 	// get the room going
// 	go r.Run()

// 	// start the web server
// 	log.Println("Starting web server on", *addr)
// 	if err := http.ListenAndServe(*addr, nil); err != nil {
// 		log.Fatal("ListenAndServe:", err)
// 	}

// }
