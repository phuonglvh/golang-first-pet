package controller

import (
	"image/png"
	"net/http"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

// ViewCodeHandler handle request code generation
func ViewCodeHandler(w http.ResponseWriter, r *http.Request) {
	dataString := r.FormValue("input")

	qrCode, _ := qr.Encode(dataString, qr.L, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 512, 512)

	png.Encode(w, qrCode)
}
