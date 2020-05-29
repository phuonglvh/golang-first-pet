package controllers

import (
	"fmt"
	"image/png"
	"net/http"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/google/uuid"
	"github.com/phuonglvh/golang-first-pet/config"
	network "github.com/phuonglvh/golang-first-pet/utils/network"
)

// QRCodeGenerationHandler handle request code generation
func QRCodeGenerationHandler(w http.ResponseWriter, r *http.Request) {
	scheme := network.GetScheme(r)
	hostname := network.GetMyIP()
	dataString := r.URL.Query().Get("string")

	if dataString == "" {
		uuid := uuid.New().String()
		dataString = uuid[len(uuid)-5:]
	}

	data := GenerateRoomLink(scheme, hostname, dataString)
	fmt.Println(data)

	qrCode, _ := qr.Encode(data, qr.L, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 256, 256)

	png.Encode(w, qrCode)
}

// GenerateRoomLink handle create a new or provided room's link
func GenerateRoomLink(scheme string, hostname string, roomID string) string {
	if roomID == "" {
		uuid := uuid.New().String()
		roomID = uuid[len(uuid)-5:]
	}
	data := scheme + "://" + hostname + ":" + fmt.Sprint(config.Env.Server.Port) + "/chat/rooms/" + roomID
	return data
}

// GenCodeGeneratorLink returns the api path of qrcode generator endpoint
func GenCodeGeneratorLink(roomID string) string {
	return "/qrcode/generator?string=" + roomID
}
