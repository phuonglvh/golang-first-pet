package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/phuonglvh/golang-first-pet/app/route"
	"github.com/phuonglvh/golang-first-pet/config"
	logger "github.com/phuonglvh/golang-first-pet/utils/logger"
)

func main() {
	logger.Init(os.Stdout, os.Stdout, os.Stdout, os.Stderr)
	router := route.Routes()
	// host := config.Cfg.Server.Host
	port := strconv.FormatInt(int64(config.Env.Server.Port), 10)
	logger.Info.Println("Server is listening on port ", port)
	http.ListenAndServe(":"+port, router)
}
