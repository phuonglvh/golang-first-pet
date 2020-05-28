package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/phuonglvh/golang-first-pet/app/controller"
	"github.com/phuonglvh/golang-first-pet/app/model"
	"github.com/phuonglvh/golang-first-pet/config"
	"github.com/phuonglvh/golang-first-pet/util/logger"
)

func main() {
	logger.Init(os.Stdout, os.Stdout, os.Stdout, os.Stderr)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controller.HomeHandler)
	router.HandleFunc("/qrcode/generator", controller.QRCodeGenerationHandler)
	router.HandleFunc("/chat/rooms/{id}", controller.ChatIndex)

	chatHandler := &controller.ChatHandler{Rooms: make(map[string]*model.Room)}
	router.Handle("/chat/rooms/{id}/ws", chatHandler)

	// host := config.Cfg.Server.Host
	port := strconv.FormatInt(int64(config.Env.Server.Port), 10)
	logger.Info.Println("Server is listening on port ", port)
	http.ListenAndServe(":"+port, router)
}
