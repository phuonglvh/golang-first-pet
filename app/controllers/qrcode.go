package controllers

import (
	"fmt"
	"image/png"
	"net/http"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/google/uuid"
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

	data := scheme + "://" + hostname + ":8080" + "/chat/rooms/" + dataString
	fmt.Println(data)

	qrCode, _ := qr.Encode(data, qr.L, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 256, 256)

	png.Encode(w, qrCode)
}
