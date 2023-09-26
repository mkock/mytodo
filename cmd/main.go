package main

import (
	"log"

	"github.com/mkock/mytodo/server"
)

func main() {
	err := server.Serve()
	log.Fatal(err)
}
