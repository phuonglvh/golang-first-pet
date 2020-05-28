package route

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/phuonglvh/golang-first-pet/app/controllers"
)

// Routes handle routing request to appropriate controllers
func Routes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// index endpoint
	router.HandleFunc("/", controllers.HomeHandler)

	// qrcode endpoint
	router.HandleFunc("/qrcode/generator", controllers.QRCodeGenerationHandler)

	// chat endpoint
	router.HandleFunc("/chat/rooms/{id}", controllers.ChatPageHandler)
	router.Handle("/chat/rooms/{id}/ws", controllers.ChatWSHandler)

	// static endpoint
	router.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	return router
}
