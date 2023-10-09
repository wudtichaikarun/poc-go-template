package main

import (
	"log"
	"os"

	"github.com/wudtichaikarun/poc-go-template/config"
	"github.com/wudtichaikarun/poc-go-template/server"
)

func main() {
	// Liveness Probe
	_, err := os.Create("/tmp/live")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove("/tmp/live")

	config := config.Load()
	server.NewServer(config).Start()
}
