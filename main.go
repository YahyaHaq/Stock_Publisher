package main

import (
	"log"

	api "github.com/BackendTest/api/server"
	"github.com/BackendTest/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("configurations could not be loaded :", err)
	}

	server, err := api.NewServer(config)
	if err != nil {
		log.Fatal("couldnt not create server :", err)
	}

	api.InitializeCache()

	server.StartServer()

	server.PublishTick()

}
