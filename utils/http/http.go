package utils

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ParseParam extract a param from request uri
func ParseParam(req *http.Request, name string) string {
	vars := mux.Vars(req)
	value := vars[name]
	return value
}
