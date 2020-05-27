package controller

import (
	"net/http"
)

// HomeHandler handle index page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "qrcode/generator", http.StatusTemporaryRedirect)
}
