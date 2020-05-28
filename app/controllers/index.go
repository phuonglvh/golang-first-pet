package controllers

import (
	"net/http"

	"github.com/phuonglvh/golang-first-pet/app/models"
	logger "github.com/phuonglvh/golang-first-pet/utils/logger"
	math "github.com/phuonglvh/golang-first-pet/utils/math"
)

// HomeHandler will create a new room, render a new qrcode to clients
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	newRoomID := math.ExtractLastNCharsOfUUID(5)
	newRoom := models.NewRoom(newRoomID)
	ChatWSHandler.Rooms[newRoomID] = newRoom
	logger.Trace.Printf("Add new room: %s to list of rooms", newRoomID)
	logger.Trace.Printf("Number of rooms: %d", len(ChatWSHandler.Rooms))
	http.Redirect(w, r, genRedirectURI(newRoomID), http.StatusTemporaryRedirect)
}

func genRedirectURI(roomID string) string {
	return "qrcode/generator?string=" + roomID
}
