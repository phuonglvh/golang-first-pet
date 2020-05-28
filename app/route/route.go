package route

import (
	"github.com/gorilla/mux"
	"github.com/phuonglvh/golang-first-pet/app/controllers"
	"github.com/phuonglvh/golang-first-pet/app/models"
)

// Routes handle routing request to appropriate controllers
func Routes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controllers.HomeHandler)
	router.HandleFunc("/qrcode/generator", controllers.QRCodeGenerationHandler)
	router.HandleFunc("/chat/rooms/{id}", controllers.ChatIndex)

	chatHandler := &controllers.ChatHandler{Rooms: make(map[string]*models.Room)}
	router.Handle("/chat/rooms/{id}/ws", chatHandler)
	return router
}
