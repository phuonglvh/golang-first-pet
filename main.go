package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/phuonglvh/golang-first-pet/app/route"
	"github.com/phuonglvh/golang-first-pet/config"
	logger "github.com/phuonglvh/golang-first-pet/utils/logger"
	network "github.com/phuonglvh/golang-first-pet/utils/network"
)

func main() {
	logger.Init(os.Stdout, os.Stdout, os.Stdout, os.Stderr)
	router := route.Routes()

	host := config.Env.Server.Host
	port := fmt.Sprint(config.Env.Server.Port)
	addr := host + ":" + port
	logger.Info.Println("Server is serving at: ", addr)
	logger.Info.Printf("Homepage: http://%s:%s/", network.GetMyIP(), port)
	http.ListenAndServe(addr, router)
}
