package controllers

import (
	"net/http"

	"github.com/phuonglvh/golang-first-pet/app/models"
	"github.com/phuonglvh/golang-first-pet/app/views"
	logger "github.com/phuonglvh/golang-first-pet/utils/logger"
	math "github.com/phuonglvh/golang-first-pet/utils/math"
	network "github.com/phuonglvh/golang-first-pet/utils/network"
)

// HomeHandler will create a new room, render a new qrcode to clients
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	newRoomID := math.ExtractLastNCharsOfUUID(5)
	newRoom := models.NewRoom(newRoomID)
	ChatWSHandler.Rooms[newRoomID] = newRoom
	logger.Trace.Printf("Add new room: %s to list of rooms", newRoomID)
	data := struct {
		QRCodeURL string
		RoomURL   string
	}{
		QRCodeURL: genRedirectURI(newRoomID),
		RoomURL:   GenerateRoomLink(network.GetScheme(r), network.GetMyIP(), newRoomID),
	}
	views.ViewResolver.ExecuteTemplate(w, "qrcode", data)
}

func genRedirectURI(roomID string) string {
	return "qrcode/generator?string=" + roomID
}
