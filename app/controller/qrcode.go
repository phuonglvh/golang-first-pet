package controller

import (
	"fmt"
	"image/png"
	"net/http"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	util "github.com/phuonglvh/golang-first-pet/util/ip"
)

// ViewCodeHandler handle request code generation
func ViewCodeHandler(w http.ResponseWriter, r *http.Request) {
	scheme := util.GetScheme(r)
	hostname := util.GetMyIP()
	dataString := r.URL.Query().Get("string")
	data := scheme + "://" + hostname + ":8080" + "/chat/rooms/" + dataString
	fmt.Println(data)

	qrCode, _ := qr.Encode(data, qr.L, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 256, 256)

	png.Encode(w, qrCode)
}
