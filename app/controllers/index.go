package controllers

import (
	"net/http"

	"github.com/google/uuid"
)

// HomeHandler handle index page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, genRedirectURI(), http.StatusTemporaryRedirect)
}

func genRedirectURI() string {
	uuid := uuid.New().String()
	uuid = uuid[len(uuid)-5:]
	return "qrcode/generator?string=" + uuid
}
